pragma solidity >=0.8.0;

contract InferCallContract {
    function inferCall(
        string calldata modelName,
        string calldata inputData
    ) public returns (bytes32) {
        bytes32[2] memory output;
        bytes memory args = abi.encodePacked(modelName, "-", inputData);
        assembly {
            if iszero(
                staticcall(
                    not(0),
                    0x100,
                    add(args, 32),
                    mload(args),
                    output,
                    12
                )
            ) {
                revert(0, 0)
            }
        }
        return output[0];
    }
}

/**
    This smart contract demo the ability to use ML/AI inference directly on-chain using NATIVE SMART CONTRACT CAll
 */
contract VannaInferCallDemo is InferCallContract {
    bytes32 volatility;

    function setVolatility(
        string calldata modelName,
        string calldata inputData
    ) public {
        volatility = inferCall(modelName, inputData);
    }

    function readVolatility() public view returns (string memory) {
        return string(abi.encodePacked(volatility));
    }

    function readVolatilityRaw() public view returns (bytes32) {
        return volatility;
    }
}
