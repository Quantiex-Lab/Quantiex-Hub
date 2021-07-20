pragma solidity ^0.5.0;

import "../../node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./Valset.sol";
import "./QuantiexERC20Bridge.sol";
import "./QuantiexERC721Bridge.sol";
import "./StakingPool.sol";


contract Oracle {
    using SafeMath for uint256;

    /*
     * @dev: Public variable declarations
     */
    QuantiexERC20Bridge public quantiexERC20Bridge;
    QuantiexERC721Bridge public quantiexERC721Bridge;
    StakingPool public stakingPool;
    Valset public valset;
    address public operator;
    uint256 public consensusThreshold; // e.g. 75 = 75%

    // Tracks the number of OracleClaims made on an individual BridgeClaim
    mapping(uint256 => address[]) public oracleClaimValidators;
    mapping(uint256 => mapping(address => bool)) public hasMadeClaim;

    /*
     * @dev: Event declarations
     */
    event LogNewOracleClaim(
        uint256 _prophecyID,
        bytes32 _message,
        address _validatorAddress,
        bytes _signature
    );

    event LogProphecyProcessed(
        uint256 _prophecyID,
        uint256 _prophecyPowerCurrent,
        uint256 _prophecyPowerThreshold,
        address _submitter
    );

    /*
     * @dev: Modifier to restrict access to the operator.
     */
    modifier onlyOperator() {
        require(msg.sender == operator, "Must be the operator.");
        _;
    }

    /*
     * @dev: Modifier to restrict access to current ValSet validators
     */
    modifier onlyValidator() {
        require(
            valset.isActiveValidator(msg.sender),
            "Must be an active validator"
        );
        _;
    }

    /*
     * @dev: Modifier to restrict access to current ValSet validators
     */
    modifier isPending(uint256 _prophecyID) {
        if (_prophecyID % 2 == 0)
        {
            require(quantiexERC20Bridge.isProphecyClaimActive(_prophecyID) == true,
                "The prophecy must be pending for this operation");
        } else {
            require(quantiexERC721Bridge.isProphecyClaimActive(_prophecyID) == true,
                "The prophecy must be pending for this operation");
        }
        _;
    }

    /*
     * @dev: Constructor
     */
    constructor(
        address _operator,
        address _valset,
        address _quantiexERC20Bridge,
        address _quantiexERC721Bridge,
        address _stakingPool,
        uint256 _consensusThreshold
    ) public {
        require(
            _consensusThreshold > 0,
            "Consensus threshold must be positive."
        );
        operator = _operator;
        quantiexERC20Bridge = QuantiexERC20Bridge(_quantiexERC20Bridge);
        quantiexERC721Bridge = QuantiexERC721Bridge(_quantiexERC721Bridge);
        valset = Valset(_valset);
        stakingPool = StakingPool(_stakingPool);
        consensusThreshold = _consensusThreshold;
    }

    /*
     * @dev: newOracleClaim
     *       Allows validators to make new OracleClaims on an existing Prophecy
     */
    function newOracleClaim(
        uint256 _prophecyID,
        bytes32 _message,
        bytes memory _signature
    ) public onlyValidator isPending(_prophecyID) {
        address validatorAddress = msg.sender;

        // Validate the msg.sender's signature
        require(
            validatorAddress == valset.recover(_message, _signature),
            "Invalid message signature."
        );

        // Confirm that this address has not already made an oracle claim on this prophecy
        require(
            !hasMadeClaim[_prophecyID][validatorAddress],
            "Cannot make duplicate oracle claims from the same address."
        );

        hasMadeClaim[_prophecyID][validatorAddress] = true;
        oracleClaimValidators[_prophecyID].push(validatorAddress);

        emit LogNewOracleClaim(
            _prophecyID,
            _message,
            validatorAddress,
            _signature
        );

        // Process the prophecy
        (
            bool valid,
            uint256 prophecyPowerCurrent,
            uint256 prophecyPowerThreshold
        ) = getProphecyThreshold(_prophecyID);

        if (valid) {
            completeProphecy(_prophecyID);

            emit LogProphecyProcessed(
                _prophecyID,
                prophecyPowerCurrent,
                prophecyPowerThreshold,
                msg.sender
            );
        }
    }

    /*
     * @dev: processBridgeProphecy
     *       Pubically available method which attempts to process a bridge prophecy
     */
    function processBridgeProphecy(uint256 _prophecyID)
        public
        isPending(_prophecyID)
    {
        // Process the prophecy
        (
            bool valid,
            uint256 prophecyPowerCurrent,
            uint256 prophecyPowerThreshold
        ) = getProphecyThreshold(_prophecyID);

        require(
            valid,
            "The cumulative power of signatory validators does not meet the threshold"
        );

        // Update the BridgeClaim's status
        completeProphecy(_prophecyID);

        emit LogProphecyProcessed(
            _prophecyID,
            prophecyPowerCurrent,
            prophecyPowerThreshold,
            msg.sender
        );
    }

    /*
     * @dev: checkBridgeProphecy
     *       Operator accessor method which checks if a prophecy has passed
     *       the validity threshold, without actually completing the prophecy.
     */
    function checkBridgeProphecy(uint256 _prophecyID)
        public
        view
        onlyOperator
        isPending(_prophecyID)
        returns (bool, uint256, uint256)
    {
        return getProphecyThreshold(_prophecyID);
    }

    /*
     * @dev: processProphecy
     *       Calculates the status of a prophecy. The claim is considered valid if the
     *       combined active signatory validator powers pass the consensus threshold.
     *       The threshold is x% of Total power, where x is the consensusThreshold param.
     */
    function getProphecyThreshold(uint256 _prophecyID)
        internal
        view
        returns (bool, uint256, uint256)
    {
        uint256 signedPower = 0;
        uint256 totalPower = stakingPool.threshold();

        // Iterate over the signatory addresses
        for (
            uint256 i = 0;
            i < oracleClaimValidators[_prophecyID].length;
            i = i.add(1)
        ) {
            address signer = oracleClaimValidators[_prophecyID][i];

            // Only add the power of active validators
            if (valset.isActiveValidator(signer)) {
                uint256 signerWeight = stakingPool.weightOf(signer);
                signedPower = signedPower.add(signerWeight);
            }
        }

        // Prophecy must reach total signed power % threshold in order to pass consensus
        uint256 prophecyPowerThreshold = totalPower.mul(consensusThreshold);
        // consensusThreshold is a decimal multiplied by 100, so signedPower must also be multiplied by 100
        uint256 prophecyPowerCurrent = signedPower.mul(100);
        bool hasReachedThreshold = prophecyPowerCurrent >= prophecyPowerThreshold;

        return (
            hasReachedThreshold,
            prophecyPowerCurrent,
            prophecyPowerThreshold
        );
    }

    /*
     * @dev: completeProphecy
     *       Completes a prophecy by completing the corresponding BridgeClaim
     *       on the QuantiexERC20Bridge.
     */
    function completeProphecy(uint256 _prophecyID) internal {
        if (_prophecyID % 2 == 0)
        {
            quantiexERC20Bridge.completeProphecyClaim(_prophecyID);
        } else {
            quantiexERC721Bridge.completeProphecyClaim(_prophecyID);
        }
    }
}
