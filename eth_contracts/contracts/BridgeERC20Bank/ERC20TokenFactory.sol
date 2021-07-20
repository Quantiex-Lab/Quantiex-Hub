pragma solidity ^0.5.0;

import "./BridgeERC20Token.sol";

/**
 * @title ERC20TokenFactory
 * @dev Create ERC20 token
 **/

contract ERC20TokenFactory {

    constructor() public
    {
    }

    /**
     * @dev external function to create a new token.
     * @param _symbol string Symbol
     */
    function createNewToken(string memory _symbol) public returns (address) {
        BridgeERC20Token newBridgeToken = (new BridgeERC20Token)(_symbol);
        return address(newBridgeToken);
    }
}
