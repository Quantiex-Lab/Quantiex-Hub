pragma solidity ^0.5.0;

import "../../../node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./BridgeERC721Token.sol";


/*
 *  @title: BinanceERC721Bank
 *  @dev: Ethereum bank which locks Ethereum/ERC20 token deposits, and unlocks
 *        Ethereum/ERC20 tokens once the prophecy has been successfully processed.
 */
contract BinanceERC721Bank {
    using SafeMath for uint256;

    uint256 public lockBurnNonce;
    mapping(address => uint256) public lockedFunds;
    mapping(string => address)  lockedTokenAddresses;
    mapping(string => bool)  lockedTokenList;

    /*
     * @dev: Event declarations
     */
    event LogBurn(
        string _chainName,
        address _from,
        address _to,
        address _token,
        string _symbol,
        uint256 _tokenId,
        string _baseURI,
        string _tokenURI,
        uint256 _nonce
    );

    event LogLock(
        string _chainName,
        address _from,
        address _to,
        address _token,
        string _symbol,
        uint256 _tokenId,
        string _baseURI,
        string _tokenURI,
        uint256 _nonce
    );

    event LogUnlock(
        address _to,
        address _token,
        string _symbol,
        uint256 _tokenId,
        string _baseURI,
        string _tokenURI
    );

    /*
     * @dev: Modifier declarations
     */

    modifier availableNonce() {
        require(lockBurnNonce + 1 > lockBurnNonce, "No available nonces.");
        _;
    }

    /*
     * @dev: Constructor which sets the lock nonce
     */
    constructor() public {
        lockBurnNonce = 0;
    }

    /*
     * @dev: Gets the contract address of locked tokens by symbol.
     *
     * @param _symbol: The asset's symbol.
     */
    function getLockedTokenAddress(string memory _symbol)
        public
        view
        returns (address)
    {
        return lockedTokenAddresses[_symbol];
    }

    /*
     * @dev: Gets the amount of locked tokens by symbol.
     *
     * @param _symbol: The asset's symbol.
     */
    function getLockedFunds(string memory _symbol)
        public
        view
        returns (uint256)
    {
        if (lockedTokenList[_symbol]) {
            return lockedFunds[lockedTokenAddresses[_symbol]];
        } else {
            return 0;
        }
    }

    /*
     * @dev: Creates a new Ethereum deposit with a unique id.
     *
     * @param _sender: The sender's ethereum address.
     * @param _recipient: The intended recipient's binance address.
     * @param _token: The currency type, either erc20 or ethereum.
     * @param _tokenId: token ID.
     * @param _tokenURI: token URI
     */
    function burnFunds(
        address payable _sender,
        address _recipient,
        address _token,
        uint256 _tokenId,
        string memory _tokenURI
    ) internal {
        lockBurnNonce = lockBurnNonce.add(1);

        string memory symbol = BridgeERC721Token(_token).symbol();
        string memory baseURI = BridgeERC721Token(_token).baseURI();
        emit LogBurn("Binance", _sender, _recipient, _token, symbol, _tokenId, baseURI, _tokenURI, lockBurnNonce);
    }

    /*
     * @dev: Creates a new Ethereum deposit with a unique id.
     *
     * @param _sender: The sender's ethereum address.
     * @param _recipient: The intended recipient's binance address.
     * @param _token: The currency type, either erc20 or ethereum.
     * @param _tokenId: token ID.
     */
    function lockFunds(
        address payable _sender,
        address _recipient,
        address _token,
        uint256 _tokenId
    ) internal {
        lockBurnNonce = lockBurnNonce.add(1);

        string memory symbol = BridgeERC721Token(_token).symbol();

        // Increment locked funds by the amount of tokens to be locked
        lockedTokenList[symbol] = true;
        lockedTokenAddresses[symbol] = _token;
        lockedFunds[_token] = lockedFunds[_token].add(1);

        string memory baseURI = BridgeERC721Token(_token).baseURI();
        string memory tokenURI = BridgeERC721Token(_token).tokenURI(_tokenId);
        emit LogLock("Binance", _sender, _recipient, _token, symbol, _tokenId, baseURI, tokenURI, lockBurnNonce);
    }

    /*
     * @dev: Unlocks funds held on contract and sends them to the
     *       intended recipient
     *
     * @param _recipient: recipient's Ethereum address
     * @param _token: token contract address
     * @param _symbol: token symbol
     * @param _tokenId: token ID
     */
    function unlockFunds(
        address payable _recipient,
        address _token,
        string memory _symbol,
        uint256 _tokenId
    ) internal {
        // Decrement locked funds mapping by the amount of tokens to be unlocked
        lockedFunds[_token] = lockedFunds[_token].sub(1);

        // Transfer funds to intended recipient
        BridgeERC721Token(_token).safeTransferFrom(address(this), _recipient, _tokenId);
        require(
            BridgeERC721Token(_token).ownerOf(_tokenId) == _recipient,
            "ERC721 Token transfer failed"
        );

        string memory baseURI = BridgeERC721Token(_token).baseURI();
        string memory tokenURI = BridgeERC721Token(_token).tokenURI(_tokenId);
        emit LogUnlock(_recipient, _token, _symbol, _tokenId, baseURI, tokenURI);
    }
}
