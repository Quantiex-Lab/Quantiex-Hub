pragma solidity ^0.5.0;

import "../../../node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./BridgeToken.sol";


/**
 * @title EthereumBank
 * @dev Manages the deployment and minting of ERC20 compatible BridgeTokens
 *      which represent assets based on the binance blockchain.
 **/

contract EthereumBank {
    using SafeMath for uint256;

    uint256 public bridgeTokenCount;
    mapping(string => address) controlledBridgeTokens;
    mapping(bytes32 => EthereumDeposit) ethereumDeposit;

    struct EthereumDeposit {
        address ethereumSender;
        address payable binanceRecipient;
        address bridgeTokenAddress;
        uint256 amount;
        bool locked;
    }

    /*
     * @dev: Event declarations
     */
    event LogNewBridgeToken(address _token, string _symbol);

    event LogBridgeTokenMint(
        address _token,
        string _symbol,
        uint256 _amount,
        address _beneficiary
    );

    /*
     * @dev: Constructor, sets bridgeTokenCount
     */
    constructor() public {
        bridgeTokenCount = 0;
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
     * @param _amount: The amount in the deposit.
     * @return: The newly created EthereumDeposit's unique id.
     */
    function newEthereumDeposit(
        address _ethereumSender,
        address payable _binanceRecipient,
        address _token,
        uint256 _amount
    ) internal returns (bytes32) {
        bytes32 depositID = keccak256(
            abi.encodePacked(_ethereumSender, _binanceRecipient, _token, _amount)
        );

        ethereumDeposit[depositID] = EthereumDeposit(
            _ethereumSender,
            _binanceRecipient,
            _token,
            _amount,
            true
        );

        return depositID;
    }

    /*
     * @dev: Deploys a new BridgeToken contract
     *
     * @param _symbol: The BridgeToken's symbol
     */
    function deployNewBridgeToken(string memory _symbol)
        internal
        returns (address)
    {
        bridgeTokenCount = bridgeTokenCount.add(1);

        // Deploy new bridge token contract
        BridgeToken newBridgeToken = (new BridgeToken)(_symbol);

        // Set address in tokens mapping
        address newBridgeTokenAddress = address(newBridgeToken);
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
     * @param _amount: number of bridge tokens to be minted
     */
    function mintNewBridgeTokens(
        address _ethereumSender,
        address payable _intendedRecipient,
        address _bridgeTokenAddress,
        string memory _symbol,
        uint256 _amount
    ) internal {
        require(
            controlledBridgeTokens[_symbol] == _bridgeTokenAddress,
            "Token must be a controlled bridge token"
        );

        // Mint bridge tokens
        require(
            BridgeToken(_bridgeTokenAddress).mint(_intendedRecipient, _amount),
            "Attempted mint of bridge tokens failed"
        );

        newEthereumDeposit(
            _ethereumSender,
            _intendedRecipient,
            _bridgeTokenAddress,
            _amount
        );

        emit LogBridgeTokenMint(
            _bridgeTokenAddress,
            _symbol,
            _amount,
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
        returns (address, address payable, address, uint256)
    {
        EthereumDeposit memory deposit = ethereumDeposit[_id];

        return (
            deposit.ethereumSender,
            deposit.binanceRecipient,
            deposit.bridgeTokenAddress,
            deposit.amount
        );
    }
}
