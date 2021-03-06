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

    struct Issue {
        uint threshold;
        string name;
    }

    mapping(address => Issue) public issues;
    address[] public addressLUT; 

    function createIssue(string name, uint threshold) returns (address) {
        address subAddress = new Transfer(threshold);

        issues[subAddress].name = name;
        issues[subAddress].threshold = threshold;
        addressLUT.push(subAddress);
        return subAddress;
    }

    function size() constant returns (uint) {
        return addressLUT.length;
    }

    function() payable {
        
    }
}