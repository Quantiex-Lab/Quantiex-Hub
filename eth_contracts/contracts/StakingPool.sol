pragma solidity ^0.5.0;

import "../../node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol";
import "../../node_modules/openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";


contract StakingPool {

    uint256 constant SECONDS_YEAR = 365 * 24 * 3600;

    /*
     * @dev: Public variable declarations
     */
    address public operator;
    address public tokenAddress;
    ERC20 public token;

    struct StakeDetail {
        uint256 StakeTime;
        uint256 Balance;
    }

    struct StakeRecord {
        uint256 idx;
        StakeDetail[] Records;
    }

    mapping(address => uint256) totalBalances;
    mapping(address => StakeRecord) stakeRecords;
    uint256 public totalBalance;

    uint256 public interestFund;
    uint256 public interestRate; // Multiplied by 100

    mapping(address => address) nextStaker;
    uint256 public stakersCount;
    address constant GUARD = address(1);
    uint256 minStakeTime;

    /*
     * @dev: Event declarations
     */
    event LogSetTokenAddress(
        address token
    );

    event LogStake(
        address staker,
        uint256 amount,
        uint256 balance
    );

    event LogRevoke(
        address staker,
        uint256 amount,
        uint256 interest,
        uint256 balance
    );

    event LogDepositInterestFund(
        address sender,
        uint256 amount,
        uint256 interestFund
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
    constructor(address _operator, uint256 _interestRate, uint256 _minStakeTime)
        public
    {
        operator = _operator;
        interestRate = _interestRate;
        minStakeTime = _minStakeTime;
        nextStaker[GUARD] = GUARD;
    }

    function setTokenAddress(address _token)
        public onlyOperator
    {
        tokenAddress = _token;
        token = ERC20(_token);

        emit LogSetTokenAddress(tokenAddress);
    }

    function depositInterestFund(uint256 amount)
        public
    {
        require(amount != 0, "Amount must not zero.");
        token.transferFrom(msg.sender, address(this), amount);
        interestFund += amount;

        emit LogDepositInterestFund(msg.sender, amount, interestFund);
    }

    function stake(uint256 amount)
        public
    {
        require(amount != 0, "Amount must not zero.");

        token.transferFrom(msg.sender, address(this), amount);

        if(nextStaker[msg.sender] == address(0)) {
            addStaker(msg.sender, amount);
        } else {
            updateBalance(msg.sender, totalBalances[msg.sender] + amount);
        }

        StakeDetail memory detail = StakeDetail(now, amount);
        stakeRecords[msg.sender].Records.push(detail);
        totalBalance += amount;

        emit LogStake(msg.sender, amount, totalBalances[msg.sender]);
    }

    function revoke()
        public
    {
        uint256 totalAmount = 0;
        uint256 totalInterest = 0;
        uint idx = stakeRecords[msg.sender].idx;

        while (stakeRecords[msg.sender].Records.length > idx) {
            uint256 stakedTime = now - stakeRecords[msg.sender].Records[idx].StakeTime;
            if (stakedTime < minStakeTime)
                break;

            uint256 interest = stakeRecords[msg.sender].Records[idx].Balance * stakedTime / SECONDS_YEAR * interestRate / 100;
            if (interestFund < totalInterest + interest)
                break;

            totalInterest += interest;
            totalAmount += stakeRecords[msg.sender].Records[idx].Balance;
            delete stakeRecords[msg.sender].Records[idx];
            stakeRecords[msg.sender].idx = ++idx;
        }

        if (totalAmount > 0) {
            require(interestFund >= totalInterest, "Interest fund is insufficient to revoke.");

            updateBalance(msg.sender, totalBalances[msg.sender] - totalAmount);
            totalBalance -= totalAmount;

            token.transfer(msg.sender, totalAmount + totalInterest);
            interestFund -= totalInterest;
        }

        emit LogRevoke(msg.sender, totalAmount, totalInterest, totalBalances[msg.sender]);
    }

    function getStakeTimes(address staker)
        public
        view
        returns(uint)
    {
        return stakeRecords[staker].Records.length - stakeRecords[staker].idx;
    }

    function getStakeRecord(address staker, uint idx)
        public
        view
        returns(uint256, uint256)
    {
        uint startIdx = stakeRecords[staker].idx;
        require(startIdx + idx < stakeRecords[staker].Records.length, "Out of range of the stakeRecords.");

        return (stakeRecords[staker].Records[startIdx + idx].StakeTime, stakeRecords[staker].Records[startIdx + idx].Balance);
    }

    function weightOf(address staker)
        public
        view
        returns (uint256)
    {
        return totalBalances[staker];
    }

    function threshold()
        public
        view
        returns (uint256)
    {
        return getTopBalance(10);
    }

    function addStaker(address staker, uint256 amount)
        internal
    {
        require(nextStaker[staker] == address(0), "Staker already exists.");

        address index = _findIndex(amount);
        totalBalances[staker] = amount;
        nextStaker[staker] = nextStaker[index];
        nextStaker[index] = staker;
        stakersCount++;
    }

    function removeStaker(address staker)
        internal
    {
        require(nextStaker[staker] != address(0), "Staker does not exist.");

        address prevStaker = _findPrevStaker(staker);
        nextStaker[prevStaker] = nextStaker[staker];
        nextStaker[staker] = address(0);
        totalBalances[staker] = 0;
        stakersCount--;
    }

    function updateBalance(address staker, uint256 newBalance)
        internal
    {
        require(nextStaker[staker] != address(0), "Staker does not exist.");

        address prevStaker = _findPrevStaker(staker);
        address nextStaker = nextStaker[staker];

        if(_verifyIndex(prevStaker, newBalance, nextStaker)) {
            totalBalances[staker] = newBalance;
        } else {
            removeStaker(staker);
            addStaker(staker, newBalance);
        }
    }

    function getTopBalance(uint256 k)
        internal
        view
        returns(uint256)
    {
        if(k > stakersCount) {
            k = stakersCount;
        }

        uint256 total = 0;
        address currentAddress = nextStaker[GUARD];
        for(uint256 i = 0; i < k; ++i) {
            total += totalBalances[currentAddress];
            currentAddress = nextStaker[currentAddress];
        }

        return total;
    }

    function getTop(uint256 k)
        public
        view
        returns(address[] memory)
    {
        if(k > stakersCount) {
            k = stakersCount;
        }

        address[] memory stakerList = new address[](k);
        address currentAddress = nextStaker[GUARD];
        for(uint256 i = 0; i < k; ++i) {
            stakerList[i] = currentAddress;
            currentAddress = nextStaker[currentAddress];
        }

        return stakerList;
    }

    function _verifyIndex(address prevStaker, uint256 newValue, address nextStaker)
        internal
        view
        returns(bool)
    {
        return (prevStaker == GUARD || totalBalances[prevStaker] >= newValue) &&
        (nextStaker == GUARD || newValue > totalBalances[nextStaker]);
    }

    function _findIndex(uint256 newValue)
        internal
        view
        returns(address)
    {
        address candidateAddress = GUARD;
        while(true) {
            if(_verifyIndex(candidateAddress, newValue, nextStaker[candidateAddress]))
                return candidateAddress;
            candidateAddress = nextStaker[candidateAddress];
        }

        return candidateAddress;
    }

    function _isPrevStaker(address staker, address prevStaker)
        internal
        view
        returns(bool)
    {
        return nextStaker[prevStaker] == staker;
    }

    function _findPrevStaker(address staker)
        internal
        view
        returns(address)
    {
        address currentAddress = GUARD;

        while(nextStaker[currentAddress] != GUARD) {
            if(_isPrevStaker(staker, currentAddress))
                return currentAddress;
            currentAddress = nextStaker[currentAddress];
        }

        return address(0);
    }
}
