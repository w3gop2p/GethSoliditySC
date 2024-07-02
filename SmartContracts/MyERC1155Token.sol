// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../node_modules/@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";

contract MyERC1155Token is ERC1155, Ownable {
    // Token ID constants
    uint256 public constant GOLD = 0;
    uint256 public constant SILVER = 1;
    uint256 public constant SWORD = 2;
    uint256 public constant SHIELD = 3;

    constructor(address initialOwner)
    ERC1155("http://127.0.0.1:8080/ipfs/{id}?filename=")
    Ownable(initialOwner)
    {
        // Mint initial supply of tokens to the contract creator
        _mint(msg.sender, GOLD, 1000, "");
        _mint(msg.sender, SILVER, 1000, "");
        _mint(msg.sender, SWORD, 10, "");
        _mint(msg.sender, SHIELD, 10, "");
    }

    function uri(uint256 tokenId) public view override returns (string memory) {
        // Generate the token-specific URI based on the tokenId
        if (tokenId == GOLD) {
            return "http://127.0.0.1:8080/ipfs/Qmd7gCM2HAm1BZWEKLnRkSbbk8fTiLZ2LxTTYn4ZC9ikzx?filename=gold.png";
        } else if (tokenId == SILVER) {
            return "http://127.0.0.1:8080/ipfs/QmTgDXuVtnL49XakCeRzeNmEbYwbPne3ZFNgsDsrAzuM1k?filename=silver.png";
        } else if (tokenId == SWORD) {
            return "http://127.0.0.1:8080/ipfs/Qmd5iFti9miWtsGuijxAhnX8Yr1srKnJXZrrruHm2NuD46?filename=sword.png";
        } else if (tokenId == SHIELD) {
            return "http://127.0.0.1:8080/ipfs/QmVF8aEhBndSzmkNC63EVUsScZhcgREN38CVY6UF2xqF3J?filename=shield.png";
        } else {
            return super.uri(tokenId); // default URI from ERC1155 base contract
        }
    }

    function mint(
        address to,
        uint256 id,
        uint256 amount,
        bytes memory data
    ) public onlyOwner {
        _mint(to, id, amount, data);
    }

    function mintBatch(
        address to,
        uint256[] memory ids,
        uint256[] memory amounts,
        bytes memory data
    ) public onlyOwner {
        _mintBatch(to, ids, amounts, data);
    }

    function burn(
        address from,
        uint256 id,
        uint256 amount
    ) public {
        require(msg.sender == from, "Only token owner can burn tokens");
        _burn(from, id, amount);
    }

    function burnBatch(
        address from,
        uint256[] memory ids,
        uint256[] memory amounts
    ) public {
        require(msg.sender == from, "Only token owner can burn tokens");
        _burnBatch(from, ids, amounts);
    }
}
