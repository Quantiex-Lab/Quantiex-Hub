pragma solidity ^0.5.0;


contract BridgeRegistry {
    address public quantiexBridge;
    address public bridgeBank;
    address public oracle;
    address public valset;
    address public stakingPool;

    event LogContractsRegistered(
        address _quantiexBridge,
        address _bridgeBank,
        address _oracle,
        address _valset,
        address _stakingPool
    );

    constructor(
        address _quantiexBridge,
        address _bridgeBank,
        address _oracle,
        address _valset,
        address _stakingPool
    ) public {
        quantiexBridge = _quantiexBridge;
        bridgeBank = _bridgeBank;
        oracle = _oracle;
        valset = _valset;
        stakingPool = _stakingPool;

        emit LogContractsRegistered(quantiexBridge, bridgeBank, oracle, valset, stakingPool);
    }
}
