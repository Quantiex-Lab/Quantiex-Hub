pragma solidity ^0.5.0;

import "../../../node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./BridgeERC721Token.sol";
import "./ERC721TokenFactory.sol";


/**
 * @title BinanceERC721Bank
 * @dev Manages the deployment and minting of ERC721 compatible BridgeTokens
 *      which represent assets based on the Binance blockchain.
 **/

contract BinanceERC721Bank {
    using SafeMath for uint256;

    ERC721TokenFactory tokenFactory;
    uint256 public bridgeTokenCount;
    mapping(string => mapping(string => address)) controlledBridgeTokens;
    mapping(bytes32 => BinanceDeposit) binanceDeposits;

    struct BinanceDeposit {
        string _chainName;
        address binanceSender;
        address payable ethereumRecipient;
        address bridgeTokenAddress;
        uint256 tokenId;
        string tokenURI;
        bool locked;
    }

    /*
     * @dev: Event declarations
     */
    event LogNewBridgeToken(string _chainName, address _token, string _symbol);

    event LogBridgeTokenMint(
        string _chainName,
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
    function getBridgeToken(string memory _chainName, string memory _symbol)
        public
        view
        returns (address)
    {
        return (controlledBridgeTokens[_chainName][_symbol]);
    }

    /*
     * @dev: Creates a new BinanceDeposit with a unique ID
     *
     * @param _binanceSender: The sender's Binance address in bytes.
     * @param _ethereumRecipient: The intended recipient's Ethereum address.
     * @param _token: The currency type
     * @return: The newly created BinanceDeposit's unique id.
     */
    function newBinanceDeposit(
        string memory _chainName,
        address _binanceSender,
        address payable _ethereumRecipient,
        address _token,
        string memory _symbol,
        uint256 _tokenId,
        string memory _tokenURI
    ) internal returns (bytes32) {
        bytes32 depositID = keccak256(
            abi.encodePacked(_binanceSender, _ethereumRecipient, _token, _symbol, _tokenId, _tokenURI)
        );

        binanceDeposits[depositID] = BinanceDeposit(
            _chainName,
            _binanceSender,
            _ethereumRecipient,
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
    function deployNewBridgeToken(string memory _chainName, string memory _symbol, string memory _baseURI)
        internal
        returns (address)
    {
        bridgeTokenCount = bridgeTokenCount.add(1);

        // Deploy new bridge token contract
        address newBridgeTokenAddress = tokenFactory.createNewToken(_symbol, _baseURI);

        // Set address in tokens mapping
        controlledBridgeTokens[_chainName][_symbol] = newBridgeTokenAddress;

        emit LogNewBridgeToken(_chainName, newBridgeTokenAddress, _symbol);
        return newBridgeTokenAddress;
    }

    /*
     * @dev: Mints new binance tokens
     *
     * @param _binanceSender: The sender's Binance address in bytes.
     * @param _ethereumRecipient: The intended recipient's Ethereum address.
     * @param _binanceTokenAddress: The currency type
     * @param _symbol: comsos token symbol
     */
    function mintNewBridgeTokens(
        string memory _chainName,
        address _binanceSender,
        address payable _intendedRecipient,
        address _bridgeTokenAddress,
        string memory _symbol,
        uint256 _tokenId,
        string memory _tokenURI
    ) internal {
        require(
            controlledBridgeTokens[_chainName][_symbol] == _bridgeTokenAddress,
            "Token must be a controlled bridge token"
        );

        // Mint bridge tokens
        require(
            BridgeERC721Token(_bridgeTokenAddress).mintTo(_intendedRecipient, _tokenId, _tokenURI),
            "Attempted mint of bridge tokens failed"
        );

        newBinanceDeposit(
            _chainName,
            _binanceSender,
            _intendedRecipient,
            _bridgeTokenAddress,
            _symbol,
            _tokenId,
            _tokenURI
        );

        emit LogBridgeTokenMint(
            _chainName,
            _bridgeTokenAddress,
            _symbol,
            _tokenId,
            _tokenURI,
            _intendedRecipient
        );
    }

    /*
     * @dev: Checks if an individual BinanceDeposit exists.
     *
     * @param _id: The unique BinanceDeposit's id.
     * @return: Boolean indicating if the BinanceDeposit exists in memory.
     */
    function isLockedBinanceDeposit(bytes32 _id) internal view returns (bool) {
        return (binanceDeposits[_id].locked);
    }

    /*
     * @dev: Gets an item's information
     *
     * @param _Id: The item containing the desired information.
     * @return: Sender's address.
     * @return: Recipient's address in bytes.
     * @return: Token address.
     * @return: Amount of ethereum/erc20 in the item.
     * @return: Unique nonce of the item.
     */
    function getBinanceDeposit(bytes32 _id)
        internal
        view
        returns (address, address, address, uint256, string memory)
    {
        return (
            binanceDeposits[_id].binanceSender,
            binanceDeposits[_id].ethereumRecipient,
            binanceDeposits[_id].bridgeTokenAddress,
            binanceDeposits[_id].tokenId,
            binanceDeposits[_id].tokenURI
        );
    }
}
