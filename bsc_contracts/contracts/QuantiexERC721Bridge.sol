pragma solidity ^0.5.0;

import "../../node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./Valset.sol";
import "./BridgeERC721Bank/BridgeERC721Bank.sol";


contract QuantiexERC721Bridge {
    using SafeMath for uint256;

    string NATIVE_ASSET_PREFIX = "PEGGY";

    /*
     * @dev: Public variable declarations
     */
    address public operator;
    Valset public valset;
    address public oracle;
    bool public hasOracle;
    BridgeERC721Bank public bridgeBank;
    bool public hasBridgeBank;

    uint256 public prophecyClaimCount;
    mapping(uint256 => ProphecyClaim) public prophecyClaims;
    mapping(string => bool) public lockburnTxHash;

    enum Status {Null, Pending, Success, Failed}

    enum ClaimType {Unsupported, Burn, Lock}

    struct ProphecyClaim {
        ClaimType claimType;
        address ethereumSender;
        address payable binanceReceiver;
        address originalValidator;
        address tokenAddress;
        string symbol;
        uint256 tokenId;
        string tokenURI;
        string txHash;
        Status status;
    }

    /*
     * @dev: Event declarations
     */

    event LogOracleSet(address _oracle);

    event LogBridgeBankSet(address _bridgeBank);

    event LogNewProphecyClaim(
        uint256 _prophecyID,
        ClaimType _claimType,
        address _ethereumSender,
        address payable _binanceReceiver,
        address _validatorAddress,
        address _tokenAddress,
        string _symbol,
        uint256 _tokenId,
        string _tokenURI,
        string _txHash
    );

    event LogProphecyCompleted(uint256 _prophecyID, ClaimType _claimType, string txHash);

    /*
     * @dev: Modifier which only allows access to currently pending prophecies
     */
    modifier isPending(uint256 _prophecyID) {
        require(
            isProphecyClaimActive(_prophecyID),
            "Prophecy claim is not active"
        );
        _;
    }

    /*
     * @dev: Modifier to restrict access to the operator.
     */
    modifier onlyOperator() {
        require(msg.sender == operator, "Must be the operator.");
        _;
    }

    /*
     * @dev: The bridge is not active until oracle and bridge bank are set
     */
    modifier isActive() {
        require(
            hasOracle == true && hasBridgeBank == true,
            "The Operator must set the oracle and bridge bank for bridge activation"
        );
        _;
    }

    /*
     * @dev: Constructor
     */
    constructor(address _operator, address _valset) public {
        prophecyClaimCount = 1;
        operator = _operator;
        valset = Valset(_valset);
        hasOracle = false;
        hasBridgeBank = false;
    }

    /*
     * @dev: setOracle
     */
    function setOracle(address _oracle) public onlyOperator {
        require(
            !hasOracle,
            "The Oracle cannot be updated once it has been set"
        );

        hasOracle = true;
        oracle = _oracle;

        emit LogOracleSet(oracle);
    }

    /*
     * @dev: setBridgeBank
     */
    function setBridgeBank(address payable _bridgeERC721Bank) public onlyOperator {
        require(
            !hasBridgeBank,
            "The Bridge Bank cannot be updated once it has been set"
        );

        hasBridgeBank = true;
        bridgeBank = BridgeERC721Bank(_bridgeERC721Bank);

        emit LogBridgeBankSet(address(bridgeBank));
    }

    /*
     * @dev: newProphecyClaim
     *       Creates a new burn or lock prophecy claim, adding it to the prophecyClaims mapping.
     *       Burn claims require that there are enough locked Ethereum assets to complete the prophecy.
     *       Lock claims have a new token contract deployed or use an existing contract based on symbol.
     */
    function newProphecyClaim(
        ClaimType _claimType,
        address _ethereumSender,
        address payable _binanceReceiver,
        string memory _symbol,
        uint256 _tokenId,
        string memory _baseURI,
        string memory _tokenURI,
        string memory _txHash
    ) public isActive {
        require(
            lockburnTxHash[_txHash] == false,
            "Filter redundant lock/burn transaction"
        );
        require(
            valset.isActiveValidator(msg.sender),
            "Must be an active validator"
        );

        address tokenAddress;
        string memory symbol;
        if (_claimType == ClaimType.Burn) {
            require(
                bridgeBank.getLockedFunds(_symbol) >= 1,
                "Not enough locked assets to complete the proposed prophecy"
            );
            symbol = _symbol;
            tokenAddress = bridgeBank.getLockedTokenAddress(_symbol);
        } else if (_claimType == ClaimType.Lock) {
            symbol = concat(NATIVE_ASSET_PREFIX, _symbol); // Add 'PEGGY' symbol prefix
            tokenAddress = bridgeBank.getBridgeToken(symbol);
            if (tokenAddress == address(0)) {
                // First lock of this asset, deploy new contract and get new symbol/token address
                tokenAddress = bridgeBank.createNewBridgeToken(symbol, _baseURI);
            }
        } else {
            revert("Invalid claim type, only burn and lock are supported.");
        }

        // Increment count and add the new ProphecyClaim to the mapping
        lockburnTxHash[_txHash] = true;
        prophecyClaimCount = prophecyClaimCount.add(2);
        prophecyClaims[prophecyClaimCount] = ProphecyClaim(
            _claimType,
            _ethereumSender,
            _binanceReceiver,
            msg.sender,
            tokenAddress,
            symbol,
            _tokenId,
            _tokenURI,
            _txHash,
            Status.Pending
        );

        emit LogNewProphecyClaim(
            prophecyClaimCount,
            _claimType,
            _ethereumSender,
            _binanceReceiver,
            msg.sender,
            tokenAddress,
            symbol,
            _tokenId,
            _tokenURI,
            _txHash
        );
    }

    /*
     * @dev: completeProphecyClaim
     *       Allows for the completion of ProphecyClaims once processed by the Oracle.
     *       Burn claims unlock tokens stored by BridgeERC721Bank.
     *       Lock claims mint BridgeTokens on BridgeERC721Bank's token whitelist.
     */
    function completeProphecyClaim(uint256 _prophecyID)
        public
        isPending(_prophecyID)
    {
        require(
            msg.sender == oracle,
            "Only the Oracle may complete prophecies"
        );

        prophecyClaims[_prophecyID].status = Status.Success;

        ClaimType claimType = prophecyClaims[_prophecyID].claimType;
        if (claimType == ClaimType.Burn) {
            unlockTokens(_prophecyID);
        } else {
            issueBridgeTokens(_prophecyID);
        }

        emit LogProphecyCompleted(_prophecyID, claimType, prophecyClaims[_prophecyID].txHash);
    }

    /*
     * @dev: issueBridgeTokens
     *       Issues a request for the BridgeERC721Bank to mint new BridgeTokens
     */
    function issueBridgeTokens(uint256 _prophecyID) internal {
        ProphecyClaim memory prophecyClaim = prophecyClaims[_prophecyID];

        bridgeBank.mintBridgeTokens(
            prophecyClaim.ethereumSender,
            prophecyClaim.binanceReceiver,
            prophecyClaim.tokenAddress,
            prophecyClaim.symbol,
            prophecyClaim.tokenId,
            prophecyClaim.tokenURI
        );
    }

    /*
     * @dev: unlockTokens
     *       Issues a request for the BridgeERC721Bank to unlock funds held on contract
     */
    function unlockTokens(uint256 _prophecyID) internal {
        ProphecyClaim memory prophecyClaim = prophecyClaims[_prophecyID];

        bridgeBank.unlock(
            prophecyClaim.binanceReceiver,
            prophecyClaim.symbol,
            prophecyClaim.tokenId
        );
    }

    /*
     * @dev: isProphecyClaimActive
     *       Returns boolean indicating if the ProphecyClaim is active
     */
    function isProphecyClaimActive(uint256 _prophecyID)
        public
        view
        returns (bool)
    {
        return prophecyClaims[_prophecyID].status == Status.Pending;
    }

    /*
     * @dev: isProphecyValidatorActive
     *       Returns boolean indicating if the validator that originally
     *       submitted the ProphecyClaim is still an active validator
     */
    function isProphecyClaimValidatorActive(uint256 _prophecyID)
        public
        view
        returns (bool)
    {
        return
            valset.isActiveValidator(
                prophecyClaims[_prophecyID].originalValidator
            );
    }

    /*
     * @dev: Performs low gas-comsuption string concatenation
     *
     * @param _prefix: start of the string
     * @param _suffix: end of the string
     */
    function concat(string memory _prefix, string memory _suffix)
        internal
        pure
        returns (string memory)
    {
        return string(abi.encodePacked(_prefix, _suffix));
    }
}
