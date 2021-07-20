pragma solidity ^0.5.0;

import "../../../node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./BridgeERC20Token.sol";


/*
 *  @title: EthereumERC20Bank
 *  @dev: Ethereum bank which locks Ethereum/ERC20 token deposits, and unlocks
 *        Ethereum/ERC20 tokens once the prophecy has been successfully processed.
 */
contract EthereumERC20Bank {
    using SafeMath for uint256;

    uint256 public lockBurnNonce;
    mapping(string => bool) lockedChainList;
    mapping(string => mapping(string => bool)) lockedTokenList;
    mapping(string => mapping(address => uint256)) public lockedFunds;
    mapping(string => mapping(string => address)) lockedTokenAddresses;

    /*
     * @dev: Event declarations
     */
    event LogBurn(
        string _chainName,
        address _from,
        address _to,
        address _token,
        string _symbol,
        uint256 _value,
        uint256 _nonce
    );

    event LogLock(
        string _chainName,
        address _from,
        address _to,
        address _token,
        string _symbol,
        uint256 _value,
        uint256 _nonce
    );

    event LogUnlock(
        string _chainName,
        address _to,
        address _token,
        string _symbol,
        uint256 _value
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
    function getLockedTokenAddress(string memory _chainName, string memory _symbol)
        public
        view
        returns (address)
    {
        return lockedTokenAddresses[_chainName][_symbol];
    }

    /*
     * @dev: Gets the amount of locked tokens by symbol.
     *
     * @param _symbol: The asset's symbol.
     */
    function getLockedFunds(string memory _chainName, string memory _symbol)
        public
        view
        returns (uint256)
    {
        if (lockedChainList[_chainName] && lockedTokenList[_chainName][_symbol]) {
            return lockedFunds[_chainName][lockedTokenAddresses[_chainName][_symbol]];
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
     * @param _amount: The amount of erc20 tokens/ ethereum (in wei) to be itemized.
     */
    function burnFunds(
        string memory _chainName,
        address payable _sender,
        address _recipient,
        address _token,
        string memory _symbol,
        uint256 _amount
    ) internal {
        lockBurnNonce = lockBurnNonce.add(1);
        emit LogBurn(_chainName, _sender, _recipient, _token, _symbol, _amount, lockBurnNonce);
    }

    /*
     * @dev: Creates a new Ethereum deposit with a unique id.
     *
     * @param _sender: The sender's ethereum address.
     * @param _recipient: The intended recipient's binance address.
     * @param _token: The currency type, either erc20 or ethereum.
     * @param _amount: The amount of erc20 tokens/ ethereum (in wei) to be itemized.
     */
    function lockFunds(
        string memory _chainName,
        address payable _sender,
        address _recipient,
        address _token,
        string memory _symbol,
        uint256 _amount
    ) internal {
        lockBurnNonce = lockBurnNonce.add(1);

        // Increment locked funds by the amount of tokens to be locked
        lockedChainList[_chainName] = true;
        lockedTokenList[_chainName][_symbol] = true;
        lockedTokenAddresses[_chainName][_symbol] = _token;
        lockedFunds[_chainName][_token] = lockedFunds[_chainName][_token].add(_amount);

        emit LogLock(_chainName, _sender, _recipient, _token, _symbol, _amount, lockBurnNonce);
    }

    /*
     * @dev: Unlocks funds held on contract and sends them to the
     *       intended recipient
     *
     * @param _recipient: recipient's Ethereum address
     * @param _token: token contract address
     * @param _symbol: token symbol
     * @param _amount: wei amount or ERC20 token count
     */
    function unlockFunds(
        string memory _chainName,
        address payable _recipient,
        address _token,
        string memory _symbol,
        uint256 _amount
    ) internal {
        // Decrement locked funds mapping by the amount of tokens to be unlocked
        lockedFunds[_chainName][_token] = lockedFunds[_chainName][_token].sub(_amount);

        // Transfer funds to intended recipient
        if (_token == address(0)) {
            _recipient.transfer(_amount);
        } else {
            require(
                BridgeERC20Token(_token).transfer(_recipient, _amount),
                "ERC20 Token transfer failed"
            );
        }

        emit LogUnlock(_chainName, _recipient, _token, _symbol, _amount);
    }
}
