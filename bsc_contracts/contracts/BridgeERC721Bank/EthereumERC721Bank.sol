pragma solidity ^0.5.0;

import "../../../node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./BridgeERC721Token.sol";
import "./ERC721TokenFactory.sol";


/**
 * @title EthereumERC721Bank
 * @dev Manages the deployment and minting of ERC20 compatible BridgeTokens
 *      which represent assets based on the binance blockchain.
 **/

contract EthereumERC721Bank {
    using SafeMath for uint256;

    ERC721TokenFactory tokenFactory;
    uint256 public bridgeTokenCount;
    mapping(string => address) controlledBridgeTokens;
    mapping(bytes32 => EthereumDeposit) ethereumDeposit;

    struct EthereumDeposit {
        address ethereumSender;
        address payable binanceRecipient;
        address bridgeTokenAddress;
        uint256 tokenId;
        string tokenURI;
        bool locked;
    }

    /*
     * @dev: Event declarations
     */
    event LogNewBridgeToken(address _token, string _symbol);

    event LogBridgeTokenMint(
        address _token,
        string _symbol,
        uint256 _tokenId,
        string _tokenURI,
        address _beneficiary
    );

    /*
     * @dev: Constructor, sets bridgeTokenCount
     */
    constructor() public {
        bridgeTokenCount = 0;
    }

    function setTokenFactory(address _tokenFactory)
        internal
    {
        tokenFactory = ERC721TokenFactory(_tokenFactory);
    }

    /*
     * @dev: Get a token symbol's corresponding bridge token address.
     *
     * @param _symbol: The token's symbol/denom without 'PEGGY' prefix.
     * @return: Address associated with the given symbol. Returns address(0) if none is found.
     */
    function getBridgeToken(string memory _symbol)
        public
        view
        returns (address)
    {
        return (controlledBridgeTokens[_symbol]);
    }

    /*
     * @dev: Creates a new EthereumDeposit with a unique ID
     *
     * @param _ethereumSender: The sender's binance address.
     * @param _binanceRecipient: The intended recipient's Binance address.
     * @param _token: The currency type
     * @return: The newly created EthereumDeposit's unique id.
     */
    function newEthereumDeposit(
        address _ethereumSender,
        address payable _binanceRecipient,
        address _token,
        string memory _symbol,
        uint256 _tokenId,
        string memory _tokenURI
    ) internal returns (bytes32) {
        bytes32 depositID = keccak256(
            abi.encodePacked(_ethereumSender, _binanceRecipient, _token, _symbol, _tokenId, _tokenURI)
        );

        ethereumDeposit[depositID] = EthereumDeposit(
            _ethereumSender,
            _binanceRecipient,
            _token,
            _tokenId,
            _tokenURI,
            true
        );

        return depositID;
    }

    /*
     * @dev: Deploys a new BridgeERC721Token contract
     *
     * @param _symbol: The BridgeERC721Token's symbol
     */
    function deployNewBridgeToken(string memory _symbol, string memory _baseURI)
        internal
        returns (address)
    {
        bridgeTokenCount = bridgeTokenCount.add(1);

        // Deploy new bridge token contract
        address newBridgeTokenAddress = tokenFactory.createNewToken(_symbol, _baseURI);

        // Set address in tokens mapping
        controlledBridgeTokens[_symbol] = newBridgeTokenAddress;

        emit LogNewBridgeToken(newBridgeTokenAddress, _symbol);
        return newBridgeTokenAddress;
    }

    /*
     * @dev: Mints new binance tokens
     *
     * @param _ethereumSender: The sender's ethereum address.
     * @param _binanceRecipient: The intended recipient's Ethereum address.
     * @param _binanceTokenAddress: The currency type
     * @param _symbol: eth or erc20 token symbol
     */
    function mintNewBridgeTokens(
        address _ethereumSender,
        address payable _intendedRecipient,
        address _bridgeTokenAddress,
        string memory _symbol,
        uint256 _tokenId,
        string memory _tokenURI
    ) internal {
        require(
            controlledBridgeTokens[_symbol] == _bridgeTokenAddress,
            "Token must be a controlled bridge token"
        );

        // Mint bridge tokens
        require(
            BridgeERC721Token(_bridgeTokenAddress).mintTo(_intendedRecipient, _tokenId, _tokenURI),
            "Attempted mint of bridge tokens failed"
        );

        newEthereumDeposit(
            _ethereumSender,
            _intendedRecipient,
            _bridgeTokenAddress,
            _symbol,
            _tokenId,
            _tokenURI
        );

        emit LogBridgeTokenMint(
            _bridgeTokenAddress,
            _symbol,
            _tokenId,
            _tokenURI,
            _intendedRecipient
        );
    }

    /*
     * @dev: Checks if an individual EthereumDeposit exists.
     *
     * @param _id: The unique EthereumDeposit's id.
     * @return: Boolean indicating if the EthereumDeposit exists in memory.
     */
    function isLockedEthereumDeposit(bytes32 _id) internal view returns (bool) {
        return (ethereumDeposit[_id].locked);
    }

    /*
     * @dev: Gets an item's information
     *
     * @param _Id: The item containing the desired information.
     * @return: Sender's address.
     * @return: Recipient's address.
     * @return: Token address.
     * @return: Amount of ethereum/erc20 in the item.
     * @return: Unique nonce of the item.
     */
    function getEthereumDeposit(bytes32 _id)
        internal
        view
        returns (address, address payable, address, uint256, string memory)
    {
        EthereumDeposit memory deposit = ethereumDeposit[_id];

        return (
            deposit.ethereumSender,
            deposit.binanceRecipient,
            deposit.bridgeTokenAddress,
            deposit.tokenId,
            deposit.tokenURI
        );
    }
}
