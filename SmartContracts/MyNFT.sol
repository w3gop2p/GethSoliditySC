// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MyNFT {
    // Events
    event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
    event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId);

    // Token name and symbol
    string public name = "MyNFT";
    string public symbol = "MNFT";

    // Mapping from token ID to owner
    mapping(uint256 => address) private _owners;

    // Mapping owner address to token count
    mapping(address => uint256) private _balances;

    // Mapping from token ID to approved address
    mapping(uint256 => address) private _tokenApprovals;

    // Total supply of tokens
    uint256 private _totalSupply;

    // Base URI for token metadata
    string private _baseURI;

    // Constructor to set the base URI
    constructor(string memory baseURI_) {
        _baseURI = baseURI_;
    }

    // Returns the number of tokens in owner's account
    function balanceOf(address owner) public view returns (uint256) {
        require(owner != address(0), "Address zero is not a valid owner");
        return _balances[owner];
    }

    // Returns the owner of the tokenId token
    function ownerOf(uint256 tokenId) public view returns (address) {
        address owner = _owners[tokenId];
        require(owner != address(0), "Token ID does not exist");
        return owner;
    }

    // Mints a new token
    function mint(address to) public returns (uint256) {
        require(to != address(0), "Mint to the zero address");

        _totalSupply++;
        uint256 tokenId = _totalSupply;

        _balances[to]++;
        _owners[tokenId] = to;

        emit Transfer(address(0), to, tokenId);

        return tokenId;
    }

    // Transfers tokenId token from from to to
    function transferFrom(address from, address to, uint256 tokenId) public {
        require(from == ownerOf(tokenId), "Transfer from incorrect owner");
        require(to != address(0), "Transfer to the zero address");

        _balances[from]--;
        _balances[to]++;
        _owners[tokenId] = to;

        emit Transfer(from, to, tokenId);
    }

    // Approves to to operate on tokenId
    function approve(address to, uint256 tokenId) public {
        address owner = ownerOf(tokenId);
        require(to != owner, "Approval to current owner");
        require(msg.sender == owner, "Approve caller is not owner");

        _tokenApprovals[tokenId] = to;
        emit Approval(owner, to, tokenId);
    }

    // Returns the approved address for tokenId token
    function getApproved(uint256 tokenId) public view returns (address) {
        require(_owners[tokenId] != address(0), "Token ID does not exist");
        return _tokenApprovals[tokenId];
    }

    // Returns the base URI
    function baseURI() public view returns (string memory) {
        return _baseURI;
    }

    // Returns the URI for tokenId token
    function tokenURI(uint256 tokenId) public view returns (string memory) {
        require(_owners[tokenId] != address(0), "Token ID does not exist");
        return string(abi.encodePacked(_baseURI, uint2str(tokenId)));
    }

    // Utility function to convert uint to string
    function uint2str(uint256 _i) internal pure returns (string memory _uintAsString) {
        if (_i == 0) {
            return "0";
        }
        uint256 j = _i;
        uint256 len;
        while (j != 0) {
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint256 k = len;
        while (_i != 0) {
            k = k - 1;
            uint8 temp = (48 + uint8(_i - _i / 10 * 10));
            bytes1 b1 = bytes1(temp);
            bstr[k] = b1;
            _i /= 10;
        }
        return string(bstr);
    }
}
