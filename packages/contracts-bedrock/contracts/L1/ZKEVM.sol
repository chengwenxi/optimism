// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { CircuitConfig } from "./CircuitConfig.sol";
import { OptimismPortal } from "./OptimismPortal.sol";
import { SafeCall } from "../libraries/SafeCall.sol";

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

    struct BatchChallenge{
        uint64 batchIndex;
        address challenger;
        uint256 challengeDeposit;
        uint256 startTime;
        bool finished;
    }

    OptimismPortal public PORTAL;
    address public submitter;
    address public challenger;

    // Last batch sent by the sequencers
    uint64 public lastBatchSequenced;

    uint256 public constant PROOF_WINDOW = 100;
    uint256 public constant FINALIZATION_PERIOD_SECONDS = 100000;
    uint256 public constant MIN_DEPOSIT = 1000000000000000000; // 1 eth

    mapping(address => uint256) public deposits;
    mapping(uint64 => bytes32) public commitments;
    mapping(uint64 => bytes32) public stateRoots;
    mapping(uint64 => uint256) public originTimestamps;
    mapping(uint64 => bool) public confirmBatchIndex;
    mapping(uint64 => BatchChallenge) public challenges; // batchIndex  => Batch

    event ChallengeState(uint256 indexed batchIndex, address challenger, uint256 challengeDeposit);
    event SubmitBatches(uint64 indexed numBatch);
    event ChallengeRes(uint256 indexed batchIndex, address winner, string res);

    // todo submitter may change to validators
    modifier onlySubmitter() {
        require(submitter != address(0),"Submitter cannot be address(0)");
        require(msg.sender == submitter, "Caller not submitter");
        _;
    }

    // todo challenger may change to validators
    modifier onlyChallenger() {
        require(challenger != address(0),"Challenger cannot be address(0)");
        require(msg.sender == challenger, "Caller not challenger");
        _;
    }

    constructor(OptimismPortal _portal, address _submitter, address _challenger){
        initialize(_portal, _submitter, _challenger);
    }

    function initialize(
        OptimismPortal _portal,
        address _submitter,
        address _challenger
    ) public payable initializer{
        __Ownable_init();
        lastBatchSequenced = 0;
        PORTAL = _portal;
        submitter = _submitter;
        challenger = _challenger;
        deposits[submitter] += msg.value;
    }

    function stake() external payable onlySubmitter{
        require(deposits[submitter] + msg.value >= MIN_DEPOSIT, "Do not have enough deposit");
        deposits[submitter] += msg.value;
    }

    function submitBatches(BatchData[] calldata batches) external onlySubmitter{
        require(deposits[submitter] >= MIN_DEPOSIT,"Insufficient staking amount");
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

            commitments[currentBatchSequenced] = hash;
            stateRoots[currentBatchSequenced] = batches[i].stateRoot;
            originTimestamps[currentBatchSequenced] = block.timestamp;
            currentBatchSequenced++;
        }
        lastBatchSequenced = currentBatchSequenced;

        emit SubmitBatches(lastBatchSequenced);
    }

    function confirmBatch(uint64 batchIndex) public {
        require(!isBatchInChallenge(batchIndex));
        bool insideChallengeWindow =
            originTimestamps[batchIndex] + FINALIZATION_PERIOD_SECONDS > block.timestamp;
        require(!insideChallengeWindow, "Cannot confirm batch inside the challenge window");
        // todo confirm one by one?
        confirmBatchIndex[batchIndex] = true;
    }

    // todo : challenge by validator
    // challengeState challenges a batch by submitting a deposit.
    function challengeState(uint64 batchIndex)
        external
        payable
        onlyChallenger
    {
        require(originTimestamps[batchIndex] != 0,"Batch not exist");
        require(challenges[batchIndex].challenger == address(0), "Already has challenge");
        // check challenge window
        // todo get finalization period from output oracle
        bool insideChallengeWindow =
            originTimestamps[batchIndex] + FINALIZATION_PERIOD_SECONDS > block.timestamp;
        require(insideChallengeWindow, "Cannot challenge batch outside the challenge window");

        // check challenge amount
        require(msg.value >= MIN_DEPOSIT);
        deposits[challenger] += msg.value;
        challenges[batchIndex] =
            BatchChallenge(batchIndex, msg.sender, msg.value, block.timestamp, false);
        emit ChallengeState(batchIndex, msg.sender, msg.value);
    }

    // proveState proves a batch by submitting a proof.
    function proveState(uint64 batchIndex, bytes calldata proof) external{
        // check challenge exists
        require(challenges[batchIndex].challenger != address(0));
        require(!challenges[batchIndex].finished, "Challenge already finished");
        bool insideChallengeWindow =
            challenges[batchIndex].startTime + PROOF_WINDOW > block.timestamp;
        if (!insideChallengeWindow) {
            _challengerWin(batchIndex,"timeout");
            // todo pause PORTAL contracts
            // PORTAL.pause();
        } else {
            // check proof
            require(proof.length > 0, "Invalid proof");
            require(
                _verifyProof(proof, commitments[batchIndex]),
                "Prove failed"
                );
            _defenderWin(batchIndex, "prove success");
        }
        challenges[batchIndex].finished = true;
    }

    function _defenderWin(uint64 batchIndex, string memory  _type) internal {
        address challengerAddr = challenges[batchIndex].challenger;
        uint256 challengeDeposit = challenges[batchIndex].challengeDeposit;
        uint256 submitterDeposit = deposits[submitter];
        deposits[challengerAddr] -= challengeDeposit;
        deposits[submitter] += challengeDeposit;
        emit ChallengeRes(batchIndex, submitter, _type);
    }

    function _challengerWin(uint64 batchIndex, string memory _type) internal {
        address challengerAddr = challenges[batchIndex].challenger;
        uint256 submitterDeposit = deposits[submitter];
        deposits[submitter] -= MIN_DEPOSIT;
        deposits[challengerAddr] += MIN_DEPOSIT;
        emit ChallengeRes(batchIndex, challengerAddr, _type);
    }

    function withdraw(uint256 amount) public {
        require(!isUserInChallenge(msg.sender),"User is in challenge");
        uint256 value = deposits[msg.sender];
        require(amount > 0);
        uint256 withdrawAmount = amount;
        if (amount + MIN_DEPOSIT > value){
            // 1. check weather sender is submitter
            // 1.1 submitter should confirm all batch then withdraw
            if (msg.sender == submitter && lastBatchSequenced != 0){
                require(
                    originTimestamps[lastBatchSequenced]+
                        FINALIZATION_PERIOD_SECONDS <= block.timestamp,
                    "submitter should wait batch to be confirm"
                );
            }
        }
        if (amount > value){
            withdrawAmount = value;
        }
        deposits[msg.sender] -= amount;
        _transfer(msg.sender, withdrawAmount);
    }

    function _transfer(address _to,uint256 _amount) internal {
        bool success = SafeCall.call(_to, gasleft(), _amount, hex"");
        require(success, "StandardBridge: ETH transfer failed");
    }

    function _verifyProof(bytes calldata proof, bytes32 commitment) internal returns (bool) {
        // TODO
        if (proof.length == 0){
            return false;
        }
        return true;
    }

    function isStaked(address addr) public view returns (bool) {
        return deposits[addr] != 0;
    }

    function isBatchInChallenge(uint64 batchIndex) public view returns (bool) {
        return challenges[batchIndex].challenger != address(0) && !challenges[batchIndex].finished;
    }

    function isUserInChallenge(address user) public view returns (bool) {
        if (lastBatchSequenced == 0){
            return false;
        }
        if (user == submitter){
                return !confirmBatchIndex[lastBatchSequenced-1];
        }else {
            for(uint64 i=0;i<lastBatchSequenced;i++){
                if(challenges[i].challenger == user && !challenges[i].finished){
                    return true;
                }
            }
        }
        return false;
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
