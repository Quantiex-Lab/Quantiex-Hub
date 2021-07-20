pragma solidity ^0.5.0;

import "../../../node_modules/openzeppelin-solidity/contracts/token/ERC721/ERC721Full.sol";

/**
 * @title BridgeERC721Token
 * @dev ERC721Full BankToken for use by BridgeERC721Bank
 **/

contract BridgeERC721Token is ERC721Full {

    address private minter;

    constructor(address _minter, string memory _symbol, string memory _baseURI)
        public
        ERC721Full(_symbol, _symbol)
    {
        minter = _minter;
        _setBaseURI(_baseURI);
    }

    /**
     * @dev Internal function to mint a new token.
     * @param to The address that will own the minted token
     * @param tokenId uint256 ID of the token to be minted
     * @param tokenURI string URI of the token to be minted
     */
    function mintTo(address to, uint256 tokenId, string memory tokenURI) public returns (bool) {
        require(msg.sender == minter, "BridgeERC721Token: caller does not have the Minter role");
        _mint(to, tokenId);
        _setTokenURI(tokenId, tokenURI);
        return true;
    }

    /**
     * @dev Burns a specific ERC721 token.
     * @param tokenId uint256 id of the ERC721 token to be burned.
     */
    function burn(uint256 tokenId) public {
        require(_isApprovedOrOwner(msg.sender, tokenId), "BridgeERC721Token: caller is not owner nor approved");
        _burn(tokenId);
    }

    /**
     * @dev Gets the list of token IDs of the requested owner.
     * @param owner address owning the tokens
     * @return uint256[] List of token IDs owned by the requested address
     */
    function listTokens(address owner) public view returns (uint256[] memory) {
        require(owner != address(0), "BridgeERC721Token: list tokens query for the zero address");
        return _tokensOfOwner(owner);
    }
}
