pragma solidity ^0.4.4;
contract Transfer {

    uint public threshold;
    address public owner;

    function Transfer(uint _threshold) {
        owner = msg.sender;
        threshold = _threshold;
    }

    function() payable {
        if (this.balance >= threshold) {
           owner.send(this.balance);
        }
    }
}

contract Spawn {

    address public subAddress;

    function createTransfer() returns (address) {
        subAddress = new Transfer(5000);
        return subAddress;
    }

    function() payable {
        
    }
}