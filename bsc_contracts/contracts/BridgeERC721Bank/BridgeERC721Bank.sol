pragma solidity ^0.5.0;

import "./BinanceERC721Bank.sol";
import "./EthereumERC721Bank.sol";
import "../Oracle.sol";
import "../QuantiexERC721Bridge.sol";


/**
 * @title BridgeERC721Bank
 * @dev Bank contract which coordinates asset-related functionality.
 *      EthereumERC721Bank manages the minting and burning of tokens which
 *      represent binance based assets, while BinanceERC721Bank manages
 *      the locking and unlocking of Ethereum and ERC721 token assets
 *      based on Ethereum.
 **/

contract BridgeERC721Bank is EthereumERC721Bank,BinanceERC721Bank,IERC721Receiver {
    using SafeMath for uint256;

    address public operator;
    QuantiexERC721Bridge public quantiexBridge;

    /*
     * @dev: Constructor, sets operator
     */
    constructor(
        address _operatorAddress,
        address _quantiexBridgeAddress,
        address _tokenFactory
    ) public {
        operator = _operatorAddress;
        quantiexBridge = QuantiexERC721Bridge(_quantiexBridgeAddress);
        setTokenFactory(_tokenFactory);
    }

    /*
     * @dev: Modifier to restrict access to operator
     */
    modifier onlyOperator() {
        require(msg.sender == operator, "Must be BridgeERC721Bank operator.");
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
     * @dev: Creates a new BridgeERC721Token
     *
     * @param _symbol: The new BridgeERC721Token's symbol
     * @return: The new BridgeERC721Token contract's address
     */
    function createNewBridgeToken(string memory _symbol, string memory _baseURI)
        public
        onlyQuantiexBridge
        returns (address)
    {
        return deployNewBridgeToken(_symbol, _baseURI);
    }

    /*
     * @dev: Mints new BankTokens
     *
     * @param _ethereumSender: The sender's ethereum address.
     * @param _intendedRecipient: The intended recipient's address.
     * @param _bridgeTokenAddress: The currency type
     * @param _symbol: comsos token symbol
     * @param _tokenId: token tokenId
     * @param _tokenURI: token tokenURI
     */
    function mintBridgeTokens(
        address _ethereumSender,
        address payable _intendedRecipient,
        address _bridgeTokenAddress,
        string memory _symbol,
        uint256 _tokenId,
        string memory _tokenURI
    ) public onlyQuantiexBridge {
        mintNewBridgeTokens(
            _ethereumSender,
            _intendedRecipient,
            _bridgeTokenAddress,
            _symbol,
            _tokenId,
            _tokenURI
        );
    }

    /*
     * @dev: Burns BridgeTokens representing native binance assets.
     *
     * @param _recipient: representation of destination address.
     * @param _token: token address in origin chain (0x0 if ethereum)
     * @param _tokenId: uint256 ID of the token
     */
    function burn(address _recipient, address _token, uint256 _tokenId)
        public
        availableNonce()
    {
        string memory tokenURI = BridgeERC721Token(_token).tokenURI(_tokenId);
        BridgeERC721Token(_token).burn(_tokenId);
        burnFunds(msg.sender, _recipient, _token, _tokenId, tokenURI);
    }

    /*
     * @dev: Locks received Ethereum/ERC721 funds.
     *
     * @param _recipient: representation of destination address.
     * @param _token: token address in origin chain (0x0 if ethereum)
     * @param _tokenId: uint256 ID of the token
     */
    function lock(address _recipient, address _token, uint256 _tokenId)
        public
        availableNonce()
    {
        BridgeERC721Token(_token).safeTransferFrom(msg.sender, address(this), _tokenId);
        lockFunds(msg.sender, _recipient, _token, _tokenId);
    }

    /*
     * @dev: Unlocks Ethereum and ERC721 tokens held on the contract.
     *
     * @param _recipient: recipient's Ethereum address
     * @param _token: token contract address
     * @param _symbol: token symbol
     * @param _tokenId: uint256 ID of the token
     */
    function unlock(
        address payable _recipient,
        string memory _symbol,
        uint256 _tokenId
    ) public onlyQuantiexBridge {
        // Confirm that the bank has sufficient locked balances of this token type
        require(
            getLockedFunds(_symbol) >= 1,
            "The Bank does not hold enough locked tokens to fulfill this request."
        );

        address tokenAddress = lockedTokenAddresses[_symbol];
        unlockFunds(_recipient, tokenAddress, _symbol, _tokenId);
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
        returns (address, address payable, address, uint256, string memory)
    {
        return getEthereumDeposit(_id);
    }

    /**
     * @notice Handle the receipt of an NFT
     * Note: the ERC721 contract address is always the message sender.
     * @param operator The address which called `safeTransferFrom` function
     * @param from The address which previously owned the token
     * @param tokenId The NFT identifier which is being transferred
     * @param data Additional data with no specified format
     * @return bytes4 `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`
     */
    function onERC721Received(address operator, address from, uint256 tokenId, bytes memory data)
        public returns (bytes4)
	{
		return bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"));
	}
}
