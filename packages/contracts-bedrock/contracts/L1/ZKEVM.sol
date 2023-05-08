// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { CircuitConfig } from "./CircuitConfig.sol";
import { OptimismPortal } from "./OptimismPortal.sol";

contract ZKEVM is CircuitConfig {
    struct BatchData {
        bytes blockWitness;
        bytes32 stateRoot;
        uint256 timestamp;
        bytes transactions; // RLP encode
        bytes32 globalExitRoot; // withdraw merkle tree
    }

    OptimismPortal public immutable PORTAL;

    // Last batch sent by the sequencers
    uint64 public lastBatchSequenced;

    uint256 public constant _WINDOW = 100;

    uint256 public constant MIN_DEPOSIT = 1000000000000000000; // 1 eth

    mapping(uint64 => bytes32) public commitments;
    mapping(uint64 => bytes32) public stateRoots;
    mapping(uint64 => uint256) public originTimestamps;
    mapping(uint64 => uint256) public challenges; // batch => amount

    event SubmitBatches(uint64 indexed numBatch);

    constructor(OptimismPortal _portal) {
        lastBatchSequenced = 0;
        PORTAL = _portal;
    }

    function submitBatches(BatchData[] calldata batches) external payable{
        uint256 batchesNum = batches.length;
        uint64 currentBatchSequenced = lastBatchSequenced;

        for (uint256 i = 0; i < batchesNum; i++) {
            uint256 chainId = 99;
            uint256 blockGas = 63000;
            (uint256 MAX_TXS, uint256 MAX_CALLDATA) = _getCircuitConfig(blockGas);
            uint256[] memory publicInput = _buildCommitment(
                MAX_TXS,
                MAX_CALLDATA,
                chainId,
                batches[i].blockWitness,
                true
            );

            bytes32 hash;
            assembly {
                hash := keccak256(add(publicInput, 32), mul(mload(publicInput), 32))
            }

            currentBatchSequenced++;
            commitments[currentBatchSequenced] = hash;
            stateRoots[currentBatchSequenced] = batches[i].stateRoot;
            originTimestamps[currentBatchSequenced] = batches[i].timestamp;
        }
        lastBatchSequenced = currentBatchSequenced;

        emit SubmitBatches(lastBatchSequenced);
    }

    // challengeState challenges a batch by submitting a deposit.
    function challengeState(uint64 batch) external payable {
        // check challenge window
        bool insideChallengeWindow = originTimestamps[batch] + _WINDOW > block.timestamp;
        require(insideChallengeWindow);

        uint256 deposit = msg.value;

        // check challenge amount
        require(deposit >= MIN_DEPOSIT);

        challenges[batch] = deposit;
    }

    // proveState proves a batch by submitting a proof.
    function proveState(uint64 batch, bytes calldata proof) external {
        // check challenge window
        bool insideChallengeWindow = originTimestamps[batch] + _WINDOW > block.timestamp;

        // check challenge exists
        require(challenges[batch] != 0);

        if (!insideChallengeWindow) {
            // pause PORTAL contracts
            // PORTAL.pause();
        } else {
            // check proof
            require(proof.length > 0);
            require(_verifyProof(proof, commitments[batch]));
        }

        // delete challenge
        delete challenges[batch];
    }

    function _verifyProof(bytes calldata proof, bytes32 commitment) internal returns (bool) {
        // TODO
        return true;
    }

    function _buildCommitment(
        uint256 MAX_TXS,
        uint256 MAX_CALLDATA,
        uint256 chainId,
        bytes calldata witness,
        bool clearMemory
    ) internal pure returns (uint256[] memory table) {
        // TODO
        return table;
    }
}
