pragma solidity ^0.5.0;


contract BridgeRegistry {
    address public quantiexERC20Bridge;
    address public quantiexERC721Bridge;
    address public bridgeERC20Bank;
    address public bridgeERC721Bank;
    address public oracle;
    address public valset;
    address public stakingPool;

    event LogContractsRegistered(
        address _quantiexERC20Bridge,
        address _quantiexERC721Bridge,
        address _bridgeERC20Bank,
        address _bridgeERC721Bank,
        address _oracle,
        address _valset,
        address _stakingPool
    );

    constructor(
        address _quantiexERC20Bridge,
        address _quantiexERC721Bridge,
        address _bridgeERC20Bank,
        address _bridgeERC721Bank,
        address _oracle,
        address _valset,
        address _stakingPool
    ) public {
        quantiexERC20Bridge = _quantiexERC20Bridge;
        quantiexERC721Bridge = _quantiexERC721Bridge;
        bridgeERC20Bank = _bridgeERC20Bank;
        bridgeERC721Bank = _bridgeERC721Bank;
        oracle = _oracle;
        valset = _valset;
        stakingPool = _stakingPool;

        emit LogContractsRegistered(quantiexERC20Bridge, quantiexERC721Bridge, bridgeERC20Bank, bridgeERC721Bank, oracle, valset, stakingPool);
    }
}
