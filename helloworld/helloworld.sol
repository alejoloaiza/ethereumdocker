pragma solidity ^0.4.24;
contract HelloWorld {
    string hello;
    constructor(string memory _hello) public {
       hello = _hello;
    }
    function Set(string memory _hello) public {
       hello = _hello;
    }
    function Get() public view returns (string memory) {
       return hello;
    }
}