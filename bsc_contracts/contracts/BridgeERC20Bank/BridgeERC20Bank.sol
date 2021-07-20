pragma solidity ^0.5.0;

import "./BinanceERC20Bank.sol";
import "./EthereumERC20Bank.sol";
import "../Oracle.sol";
import "../QuantiexERC20Bridge.sol";


/**
 * @title BridgeERC20Bank
 * @dev Bank contract which coordinates asset-related functionality.
 *      EthereumERC20Bank manages the minting and burning of tokens which
 *      represent binance based assets, while BinanceERC20Bank manages
 *      the locking and unlocking of Ethereum and ERC20 token assets
 *      based on Ethereum.
 **/

contract BridgeERC20Bank is EthereumERC20Bank,BinanceERC20Bank {
    using SafeMath for uint256;

    address public operator;
    QuantiexERC20Bridge public quantiexBridge;

    /*
     * @dev: Constructor, sets operator
     */
    constructor(
        address _operatorAddress,
        address _quantiexBridgeAddress
    ) public {
        operator = _operatorAddress;
        quantiexBridge = QuantiexERC20Bridge(_quantiexBridgeAddress);
    }

    /*
     * @dev: Modifier to restrict access to operator
     */
    modifier onlyOperator() {
        require(msg.sender == operator, "Must be BridgeERC20Bank operator.");
        _;
    }

    /*
     * @dev: Modifier to restrict access to the binance bridge
     */
    modifier onlyQuantiexBridge() {
        require(
            msg.sender == address(quantiexBridge),
            "Access restricted to the binance bridge"
        );
        _;
    }

    /*
     * @dev: Fallback function allows operator to send funds to the bank directly
     *       This feature is used for testing and is available at the operator's own risk.
     */
    function() external payable onlyOperator { }

    /*
     * @dev: Creates a new BridgeERC20Token
     *
     * @param _symbol: The new BridgeERC20Token's symbol
     * @return: The new BridgeERC20Token contract's address
     */
    function createNewBridgeToken(string memory _symbol)
        public
        onlyQuantiexBridge
        returns (address)
    {
        return deployNewBridgeToken(_symbol);
    }

    /*
     * @dev: Mints new BankTokens
     *
     * @param _ethereumSender: The sender's ethereum address.
     * @param _intendedRecipient: The intended recipient's address.
     * @param _bridgeTokenAddress: The currency type
     * @param _symbol: comsos token symbol
     * @param _amount: number of comsos tokens to be minted
     */
    function mintBridgeTokens(
        address _ethereumSender,
        address payable _intendedRecipient,
        address _bridgeTokenAddress,
        string memory _symbol,
        uint256 _amount
    ) public onlyQuantiexBridge {
        mintNewBridgeTokens(
            _ethereumSender,
            _intendedRecipient,
            _bridgeTokenAddress,
            _symbol,
            _amount
        );
    }

    /*
     * @dev: Burns BridgeTokens representing native binance assets.
     *
     * @param _recipient: representation of destination address.
     * @param _token: token address in origin chain (0x0 if ethereum)
     * @param _amount: value of deposit
     */
    function burn(address _recipient, address _token, uint256 _amount)
        public
        availableNonce()
    {
        BridgeERC20Token(_token).burnFrom(msg.sender, _amount);
        string memory symbol = BridgeERC20Token(_token).symbol();
        burnFunds(msg.sender, _recipient, _token, symbol, _amount);
    }

    /*
     * @dev: Locks received Ethereum/ERC20 funds.
     *
     * @param _recipient: representation of destination address.
     * @param _token: token address in origin chain (0x0 if ethereum)
     * @param _amount: value of deposit
     */
    function lock(address _recipient, address _token, uint256 _amount)
        public
        payable
        availableNonce()
    {
        string memory symbol;

        // Ethereum deposit
        if (msg.value > 0) {
            require(
                _token == address(0),
                "Ethereum deposits require the 'token' address to be the null address"
            );
            require(
                msg.value == _amount,
                "The transactions value must be equal the specified amount (in wei)"
            );
            symbol = "BNB";
        // ERC20 deposit
        } else {
            require(
                BridgeERC20Token(_token).transferFrom(
                    msg.sender,
                    address(this),
                    _amount
                ),
                "Contract token allowances insufficient to complete this lock request"
            );
            // Set symbol to the ERC20 token's symbol
            symbol = BridgeERC20Token(_token).symbol();
        }

        lockFunds(msg.sender, _recipient, _token, symbol, _amount);
    }

    /*
     * @dev: Unlocks Ethereum and ERC20 tokens held on the contract.
     *
     * @param _recipient: recipient's Ethereum address
     * @param _token: token contract address
     * @param _symbol: token symbol
     * @param _amount: wei amount or ERC20 token count
     */
    function unlock(
        address payable _recipient,
        string memory _symbol,
        uint256 _amount
    ) public onlyQuantiexBridge {
        // Confirm that the bank has sufficient locked balances of this token type
        require(
            getLockedFunds(_symbol) >= _amount,
            "The Bank does not hold enough locked tokens to fulfill this request."
        );

        // Confirm that the bank holds sufficient balances to complete the unlock
        address tokenAddress = lockedTokenAddresses[_symbol];
        if (tokenAddress == address(0)) {
            address thisadd = address(this);
            require(
                thisadd.balance >= _amount,
                "Insufficient ethereum balance for delivery."
            );
        } else {
            require(
                BridgeERC20Token(tokenAddress).balanceOf(address(this)) >= _amount,
                "Insufficient ERC20 token balance for delivery."
            );
        }
        unlockFunds(_recipient, tokenAddress, _symbol, _amount);
    }

    /*
     * @dev: Exposes an item's current status.
     *
     * @param _id: The item in question.
     * @return: Boolean indicating the lock status.
     */
    function getDepositStatus(bytes32 _id) public view returns (bool) {
        return isLockedEthereumDeposit(_id);
    }

    /*
     * @dev: Allows access to deposit's information via its unique identifier.
     *
     * @param _id: The deposit to be viewed.
     * @return: Original sender's Ethereum address.
     * @return: Intended binance recipient's address.
     * @return: The lock deposit's currency, denoted by a token address.
     * @return: The amount locked in the deposit.
     * @return: The deposit's unique nonce.
     */
    function viewDeposit(bytes32 _id)
        public
        view
        returns (address, address payable, address, uint256)
    {
        return getEthereumDeposit(_id);
    }
}
