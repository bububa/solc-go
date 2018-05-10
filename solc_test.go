package solc

import (
	"encoding/json"
	"log"
	"testing"
)

func TestLicense(t *testing.T) {
	license := License()
	log.Printf("License: %s\n", license)
}

func TestVersion(t *testing.T) {
	version := Version()
	log.Printf("Version: %s\n", version)
}

func TestCompileJSON(t *testing.T) {
	input := `pragma solidity ^0.4.22;

contract owned {
    address owner;

    modifier onlyowner() {
        if (msg.sender == owner) {
            _;
        }
    }

    constructor() public {
        owner = msg.sender;
    }
}`
	ret := CompileJSON(input, true)
	log.Println(ret)
}

func TestCompileJSONMulti(t *testing.T) {
	input := map[string]string{
		"owned.sol": `pragma solidity ^0.4.22;

contract owned {
    address owner;

    modifier onlyowner() {
        if (msg.sender == owner) {
            _;
        }
    }

    constructor() public {
        owner = msg.sender;
    }
}`,
		"motal.sol": `pragma solidity ^0.4.0;

import "./owned.sol";

contract mortal is owned {
    function kill() public {
        if (msg.sender == owner)
            selfdestruct(owner);
    }
}`,
		"Token.sol": `pragma solidity ^0.4.0;

contract Token {
    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);

    function totalSupply() constant public returns (uint256 supply);
    function balanceOf(address _owner) constant public returns (uint256 balance);
    function transfer(address _to, uint256 _value) public returns (bool success);
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success);
    function approve(address _spender, uint256 _value) public returns (bool success);
    function allowance(address _owner, address _spender) constant public returns (uint256 remaining);
}`,
		"StandardToken.sol": `pragma solidity ^0.4.22;

import "./Token.sol";

contract StandardToken is Token {
    uint256 supply;
    mapping (address => uint256) balance;
    mapping (address =>
        mapping (address => uint256)) m_allowance;

    constructor(address _initialOwner, uint256 _supply) public {
        supply = _supply;
        balance[_initialOwner] = _supply;
    }

    function balanceOf(address _account) constant public returns (uint) {
        return balance[_account];
    }

    function totalSupply() constant public returns (uint) {
        return supply;
    }

    function transfer(address _to, uint256 _value) public returns (bool success) {
        return doTransfer(msg.sender, _to, _value);
    }

    function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
        if (m_allowance[_from][msg.sender] >= _value) {
            if (doTransfer(_from, _to, _value)) {
                m_allowance[_from][msg.sender] -= _value;
            }
            return true;
        } else {
            return false;
        }
    }

    function doTransfer(address _from, address _to, uint _value) internal returns (bool success) {
        if (balance[_from] >= _value && balance[_to] + _value >= balance[_to]) {
            balance[_from] -= _value;
            balance[_to] += _value;
            emit Transfer(_from, _to, _value);
            return true;
        } else {
            return false;
        }
    }

    function approve(address _spender, uint256 _value) public returns (bool success) {
        m_allowance[msg.sender][_spender] = _value;
        emit Approval(msg.sender, _spender, _value);
        return true;
    }

    function allowance(address _owner, address _spender) constant public returns (uint256) {
        return m_allowance[_owner][_spender];
    }
}`,
	}
	js, _ := json.Marshal(input)
	ret := CompileJSON(string(js), true)
	log.Println(ret)
}

func TestCompileStandard(t *testing.T) {
	input := `{ "language": "Solidity",
  "sources":
  {
    "owned.sol":
    {
        "content": "pragma solidity ^0.4.22;  contract owned { address owner; modifier onlyowner() { if (msg.sender == owner) { _;  } } constructor() public {owner = msg.sender; } }"
    },
    "motal.sol":
    {
        "content": "pragma solidity ^0.4.0;

import \"owned.sol\";

contract mortal is owned {
    function kill() public {
        if (msg.sender == owner)
            selfdestruct(owner);
    }
}"
    },
    "Token.sol": {
        "content": "pragma solidity ^0.4.0;

contract Token {
    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);

    function totalSupply() constant public returns (uint256 supply);
    function balanceOf(address _owner) constant public returns (uint256 balance);
    function transfer(address _to, uint256 _value) public returns (bool success);
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success);
    function approve(address _spender, uint256 _value) public returns (bool success);
    function allowance(address _owner, address _spender) constant public returns (uint256 remaining);
}"
    },
    "StandardToken.sol": {
        "content": "pragma solidity ^0.4.22;

import \"Token.sol\";

contract StandardToken is Token {
    uint256 supply;
    mapping (address => uint256) balance;
    mapping (address =>
        mapping (address => uint256)) m_allowance;

    constructor(address _initialOwner, uint256 _supply) public {
        supply = _supply;
        balance[_initialOwner] = _supply;
    }

    function balanceOf(address _account) constant public returns (uint) {
        return balance[_account];
    }

    function totalSupply() constant public returns (uint) {
        return supply;
    }

    function transfer(address _to, uint256 _value) public returns (bool success) {
        return doTransfer(msg.sender, _to, _value);
    }

    function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
        if (m_allowance[_from][msg.sender] >= _value) {
            if (doTransfer(_from, _to, _value)) {
                m_allowance[_from][msg.sender] -= _value;
            }
            return true;
        } else {
            return false;
        }
    }

    function doTransfer(address _from, address _to, uint _value) internal returns (bool success) {
        if (balance[_from] >= _value && balance[_to] + _value >= balance[_to]) {
            balance[_from] -= _value;
            balance[_to] += _value;
            emit Transfer(_from, _to, _value);
            return true;
        } else {
            return false;
        }
    }

    function approve(address _spender, uint256 _value) public returns (bool success) {
        m_allowance[msg.sender][_spender] = _value;
        emit Approval(msg.sender, _spender, _value);
        return true;
    }

    function allowance(address _owner, address _spender) constant public returns (uint256) {
        return m_allowance[_owner][_spender];
    }
}"
    }
  },
  "settings":
  {
    "optimizer": {
      "enabled": true,
      "runs": 200
    },
    "evmVersion": "homestead",
    "metadata": {
      "useLiteralContent": true
    },
    "outputSelection": {
      "*": {
        "*": [ "abi", "evm.bytecode.object"]
      }
    }
  }
}`
	ret := CompileStandard(input)
	log.Println(ret)
}
