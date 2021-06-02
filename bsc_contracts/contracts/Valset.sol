pragma solidity ^0.5.0;

import "../../node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol";


contract Valset {
    using SafeMath for uint256;

    /*
     * @dev: Variable declarations
     */
    address public operator;
    uint256 public currentValsetVersion;
    uint256 public validatorCount;
    mapping(bytes32 => bool) public validators;

    /*
     * @dev: Event declarations
     */
    event LogValidatorAdded(
        address _validator,
        uint256 _currentValsetVersion,
        uint256 _validatorCount,
        uint256 _totalPower
    );

    event LogValidatorPowerUpdated(
        address _validator,
        uint256 _power,
        uint256 _currentValsetVersion,
        uint256 _validatorCount,
        uint256 _totalPower
    );

    event LogValidatorRemoved(
        address _validator,
        uint256 _currentValsetVersion,
        uint256 _validatorCount,
        uint256 _totalPower
    );

    event LogValsetReset(
        uint256 _newValsetVersion,
        uint256 _validatorCount,
        uint256 _totalPower
    );

    event LogValsetUpdated(
        uint256 _newValsetVersion,
        uint256 _validatorCount,
        uint256 _totalPower
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
    constructor(
        address _operator,
        address[] memory _initValidators
    ) public {
        operator = _operator;
        currentValsetVersion = 0;

        updateValset(_initValidators);
    }

    function recover(bytes32 _message, bytes memory _signature)
        public
        pure
        returns (address)
    {
        return verify(bscMessageHash(_message), _signature);
    }

    /*
     * @dev: addValidator
     */
    function addValidator(address _validatorAddress)
        public
        onlyOperator
    {
        addValidatorInternal(_validatorAddress);
    }

    /*
     * @dev: removeValidator
     */
    function removeValidator(address _validatorAddress) public onlyOperator {
        // Create a unique key which for this validator's position in the current version of the mapping
        bytes32 key = keccak256(
            abi.encodePacked(currentValsetVersion, _validatorAddress)
        );

        require(validators[key], "Can only remove active valdiators");

        // Update validator count and total power
        validatorCount = validatorCount.sub(1);

        // Delete validator and power
        delete validators[key];

        emit LogValidatorRemoved(
            _validatorAddress,
            currentValsetVersion,
            validatorCount,
            0
        );
    }

    /*
     * @dev: updateValset
     */
    function updateValset(
        address[] memory _validators
    ) public onlyOperator {
        resetValset();

        for (uint256 i = 0; i < _validators.length; i = i.add(1)) {
            addValidatorInternal(_validators[i]);
        }

        emit LogValsetUpdated(currentValsetVersion, validatorCount, 0);
    }

    /*
     * @dev: isActiveValidator
     */
    function isActiveValidator(address _validatorAddress)
        public
        view
        returns (bool)
    {
        // Recreate the unique key for this address given the current mapping version
        bytes32 key = keccak256(
            abi.encodePacked(currentValsetVersion, _validatorAddress)
        );

        // Return bool indicating if this address is an active validator
        return validators[key];
    }

    /*
     * @dev: recoverGas
     */
    function recoverGas(uint256 _valsetVersion, address _validatorAddress)
        external
        onlyOperator
    {
        require(
            _valsetVersion < currentValsetVersion,
            "Gas recovery only allowed for previous validator sets"
        );

        // Recreate the unique key used to identify this validator in the given version
        bytes32 key = keccak256(
            abi.encodePacked(_valsetVersion, _validatorAddress)
        );

        // Delete from mappings and recover gas
        delete (validators[key]);
    }

    /*
     * @dev: addValidatorInternal
     */
    function addValidatorInternal(
        address _validatorAddress
    ) internal {
        // Create a unique key which for this validator's position in the current version of the mapping
        bytes32 key = keccak256(
            abi.encodePacked(currentValsetVersion, _validatorAddress)
        );

        validatorCount = validatorCount.add(1);

        // Set validator as active and set their power
        validators[key] = true;

        emit LogValidatorAdded(
            _validatorAddress,
            currentValsetVersion,
            validatorCount,
            0
        );
    }

    /*
     * @dev: resetValset
     */
    function resetValset() internal {
        currentValsetVersion = currentValsetVersion.add(1);
        validatorCount = 0;

        emit LogValsetReset(currentValsetVersion, validatorCount, 0);
    }

    /**
     * @dev Recover signer address from a message by using their signature
     * @param h bytes32 message, the hash is the signed message. What is recovered is the signer address.
     * @param signature bytes signature, the signature is generated using web3.eth.sign()
     */
    function verify(bytes32 h, bytes memory signature)
        internal
        pure
        returns (address)
    {
        bytes32 r;
        bytes32 s;
        uint8 v;

        // Check the signature length
        if (signature.length != 65) {
            return (address(0));
        }

        // Divide the signature in r, s and v variables
        // ecrecover takes the signature parameters, and the only way to get them
        // currently is to use assembly.
        // solium-disable-next-line security/no-inline-assembly
        assembly {
            r := mload(add(signature, 32))
            s := mload(add(signature, 64))
            v := byte(0, mload(add(signature, 96)))
        }

        // Version of signature should be 27 or 28, but 0 and 1 are also possible versions
        if (v < 27) {
            v += 27;
        }

        // If the version is correct return the signer address
        if (v != 27 && v != 28) {
            return (address(0));
        } else {
            // solium-disable-next-line arg-overflow
            return ecrecover(h, v, r, s);
        }
    }

    /**
     * toBscSignedMessageHash
     * @dev prefix a bytes32 value with "\x19Binance Signed Message:"
     * and hash the result
     */
    function bscMessageHash(bytes32 message) internal pure returns (bytes32) {
        return
        keccak256(
            abi.encodePacked("\x19Binance Signed Message:\n32", message)
        );
    }
}
