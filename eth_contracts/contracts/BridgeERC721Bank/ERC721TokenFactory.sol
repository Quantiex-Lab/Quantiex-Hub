pragma solidity ^0.5.0;

import "./BridgeERC721Token.sol";

/**
 * @title ERC721TokenFactory
 * @dev Create ERC721 token
 **/

contract ERC721TokenFactory {

    constructor() public
    {
    }

    /**
     * @dev public function to create a new token.
     * @param symbol string Symbol
     * @param baseURI uint256 ID of the token to be minted
     */
    function createNewToken(string memory symbol, string memory baseURI) public returns (address) {
        BridgeERC721Token newToken = new BridgeERC721Token(msg.sender, symbol, baseURI);
        return address(newToken);
    }
}
