// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract EthSender {
    address public owner;
    uint256 constant amountToSend = 42 ether;

    // Mapping to store the total amount of ETH sent to each EOA
    mapping(address => uint256) public totalSent;

    // Event to log the sending of ETH
    event EthSent(address indexed recipient, uint256 amount);

    // Modifier to check if the caller is the owner
    modifier onlyOwner() {
        require(msg.sender == owner, "Caller is not the owner");
        _;
    }

    // Constructor to set the owner of the contract
    constructor() {
        owner = msg.sender;
    }

    // Function to send 42 ETH to the caller
    function sendEth() external {
        require(address(this).balance >= amountToSend, "Not enough balance in contract");
        payable(msg.sender).transfer(amountToSend);
        totalSent[msg.sender] += amountToSend;
        emit EthSent(msg.sender, amountToSend);
    }

    // Function to view the total amount of ETH sent to an address
    function getTotalSent(address _address) external view returns (uint256) {
        return totalSent[_address];
    }

    // Function to receive ETH into the contract
    receive() external payable {}

    // Function to allow the owner to withdraw funds from the contract
    function withdraw(uint256 _amount) external onlyOwner {
        require(address(this).balance >= _amount, "Not enough balance in contract");
        payable(owner).transfer(_amount);
    }
}
