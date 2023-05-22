// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { CircuitConfig } from "./CircuitConfig.sol";
import { OptimismPortal } from "./OptimismPortal.sol";

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract ZKEVM is CircuitConfig,OwnableUpgradeable {
    struct BatchData {
        bytes blockWitness;
        bytes32 stateRoot;
        uint256 timestamp;
        bytes transactions; // RLP encode
        bytes32 globalExitRoot; // withdraw merkle tree
    }

    struct Staker {
        address stakeAddress;
        uint256 stakeAmount;
    }

    OptimismPortal public PORTAL;
    address public submitter;
    address public challenger;

    // Last batch sent by the sequencers
    uint64 public lastBatchSequenced;

    uint256 public constant _WINDOW = 100;

    uint256 public constant MIN_DEPOSIT = 1000000000000000000; // 1 eth

    mapping(address => uint256) public deposits;
    mapping(uint64 => bytes32) public commitments;
    mapping(uint64 => bytes32) public stateRoots;
    mapping(uint64 => uint256) public originTimestamps;
    mapping(uint64 => uint256) public challenges; // batch => amount

    event SubmitBatches(uint64 indexed numBatch);

    modifier onlySubmitter() {
        require(msg.sender == submitter, "Caller not submitter");
        require(deposits[submitter] >=  MIN_DEPOSIT, "Do not have enought deposit");
        _;
    }

    function initialize(OptimismPortal _portal, address _submitter) external payable initializer{
        __Ownable_init();
        lastBatchSequenced = 0;
        PORTAL = _portal;
        submitter = _submitter;
        deposits[submitter] += msg.value;
    }

    function stake() external payable onlySubmitter{
        require(deposits[submitter] + msg.value >= MIN_DEPOSIT, "Do not have enought deposit");
        deposits[submitter] += msg.value;
    }

    function submitBatches(BatchData[] calldata batches) external onlySubmitter{
        require(deposits[submitter] > MIN_DEPOSIT);
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
        require(challenger == address(0), "Already in challenge");
        // check challenge window
        bool insideChallengeWindow = originTimestamps[batch] + _WINDOW > block.timestamp;
        require(insideChallengeWindow, "Cannot challenge batche outside the challenge window");

        challenger = msg.sender;

        // check challenge amount
        require(msg.value >= MIN_DEPOSIT);
        deposits[challenger] += msg.value;
        challenges[batch] = msg.value;
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
