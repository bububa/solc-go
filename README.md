# solc-go
solidity compiler golang wrapper using libsolc.a c API.

## Installation
- download boost_1_67_0.tar.gz into dependencies folder;
- make deps;

## Usage

```

import "github.com/bububa/solc-go"
import "log"

func main() {
    input := "pragma solidity ^0.4.22; contract owned { address owner; modifier onlyowner() { if (msg.sender == owner) { _; } } constructor() public { owner = msg.sender; } }`
    ret := CompileJSON(input, true)
    log.Println(ret)
}
```