pragma solidity ^0.5.0;

import "../../node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol";
import "../../node_modules/openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";


contract StakingPool {

    /*
     * @dev: Public variable declarations
     */
    address public operator;
    address public tokenAddress;
    ERC20 public token;

    mapping(address => uint256) public balances;
    mapping(address => address) _nextStaker;
    uint256 public stakersCount;
    address constant GUARD = address(1);

    uint256 public totalBalance;

    /*
     * @dev: Event declarations
     */
    event LogSetTokenAddress(
        address _token
    );

    event LogStake(
        address _staker,
        uint256 _value,
        uint256 _totalAmount
    );

    event LogRevoke(
        address _staker,
        uint256 _value,
        uint256 _totalAmount
    );

    /*
     * @dev: Modifier which restricts access to the operator.
     */
    modifier onlyOperator() {
        require(msg.sender == operator, "Must be the operator.");
        _;
    }

    /*
     * @dev: Constructor
     */
    constructor(address _operator)
    public {
        operator = _operator;
        _nextStaker[GUARD] = GUARD;
    }

    function setTokenAddress(address _token)
    public onlyOperator {
        tokenAddress = _token;
        token = ERC20(_token);

        emit LogSetTokenAddress(tokenAddress);
    }

    function stake(uint256 amount)
    public {
        require(amount != 0, "Amount must not zero.");

        token.transferFrom(msg.sender, address(this), amount);

        if(_nextStaker[msg.sender] == address(0)) {
            addStaker(msg.sender, amount);
        } else {
            updateBalance(msg.sender, balances[msg.sender] + amount);
        }
        totalBalance += amount;

        emit LogStake(msg.sender, amount, balances[msg.sender]);
    }

    function revoke(address recipient, uint256 amount)
    public onlyOperator {
        require(amount <= balances[recipient], "Staking insufficient to revoke.");

        token.transfer(recipient, amount);
        updateBalance(recipient, balances[recipient] - amount);
        totalBalance -= amount;

        emit LogRevoke(recipient, amount, balances[recipient]);
    }

    function weightOf(address staker)
    public
    view
    returns (uint256) {
        return balances[staker];
    }

    function threshold()
    public
    view
    returns (uint256) {
        return getTopBalance(10);
    }

    function addStaker(address staker, uint256 amount)
    internal {
        require(_nextStaker[staker] == address(0), "Staker already exists.");

        address index = _findIndex(amount);
        balances[staker] = amount;
        _nextStaker[staker] = _nextStaker[index];
        _nextStaker[index] = staker;
        stakersCount++;
    }

    function removeStaker(address staker)
    internal {
        require(_nextStaker[staker] != address(0), "Staker does not exist.");

        address prevStaker = _findPrevStaker(staker);
        _nextStaker[prevStaker] = _nextStaker[staker];
        _nextStaker[staker] = address(0);
        balances[staker] = 0;
        stakersCount--;
    }

    function updateBalance(address staker, uint256 newBalance)
    internal {
        require(_nextStaker[staker] != address(0), "Staker does not exist.");

        address prevStaker = _findPrevStaker(staker);
        address nextStaker = _nextStaker[staker];

        if(_verifyIndex(prevStaker, newBalance, nextStaker)) {
            balances[staker] = newBalance;
        } else {
            removeStaker(staker);
            addStaker(staker, newBalance);
        }
    }

    function getTopBalance(uint256 k)
    internal
    view
    returns(uint256) {
        if(k > stakersCount) {
            k = stakersCount;
        }

        uint256 total = 0;
        address currentAddress = _nextStaker[GUARD];
        for(uint256 i = 0; i < k; ++i) {
            total += balances[currentAddress];
            currentAddress = _nextStaker[currentAddress];
        }

        return total;
    }

    function getTop(uint256 k)
    internal
    view
    returns(address[] memory) {
        if(k > stakersCount) {
            k = stakersCount;
        }

        address[] memory stakerList = new address[](k);
        address currentAddress = _nextStaker[GUARD];
        for(uint256 i = 0; i < k; ++i) {
            stakerList[i] = currentAddress;
            currentAddress = _nextStaker[currentAddress];
        }

        return stakerList;
    }

    function _verifyIndex(address prevStaker, uint256 newValue, address nextStaker)
    internal
    view
    returns(bool)
    {
        return (prevStaker == GUARD || balances[prevStaker] >= newValue) &&
        (nextStaker == GUARD || newValue > balances[nextStaker]);
    }

    function _findIndex(uint256 newValue)
    internal
    view
    returns(address) {
        address candidateAddress = GUARD;
        while(true) {
            if(_verifyIndex(candidateAddress, newValue, _nextStaker[candidateAddress]))
                return candidateAddress;
            candidateAddress = _nextStaker[candidateAddress];
        }

        return candidateAddress;
    }

    function _isPrevStaker(address staker, address prevStaker)
    internal
    view
    returns(bool) {
        return _nextStaker[prevStaker] == staker;
    }

    function _findPrevStaker(address staker)
    internal
    view
    returns(address) {
        address currentAddress = GUARD;

        while(_nextStaker[currentAddress] != GUARD) {
            if(_isPrevStaker(staker, currentAddress))
                return currentAddress;
            currentAddress = _nextStaker[currentAddress];
        }

        return address(0);
    }
}
