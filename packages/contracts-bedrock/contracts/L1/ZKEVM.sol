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
        address challenger;
        uint64 batchIndex;
        uint256 challengeDeposit;
        uint256 startTime;
        bool finished;
    }

    OptimismPortal public PORTAL;
    address public submitter;
    address public challenger;

    // Last batch sent by the sequencers
    uint64 public lastBatchSequenced;
    uint64 public challengeNum;

    uint256 public constant PROOF_WINDOW = 100;
    uint256 public constant FINALIZATION_PERIOD_SECONDS = 100000;
    uint256 public constant MIN_DEPOSIT = 1000000000000000000; // 1 eth

    mapping(address => uint256) public deposits;
    mapping(uint64 => bytes32) public commitments;
    mapping(uint64 => bytes32) public stateRoots;
    mapping(uint64 => uint256) public originTimestamps;
    mapping(uint64 => BatchChallenge) public challenges; // batchIndex  => Batch

    event ChallengeState(uint256 indexed challengeIndex, address challenger, uint256 batchIndex, uint256 challengeDeposit);
    event SubmitBatches(uint64 indexed numBatch);
    event ChallengeRes(uint256 indexed challengeIndex, address winner, string res);

    // todo submitter may change to validators
    modifier onlySubmitter() {
        require(submitter != address(0),"Submitter cannot be address(0)");
        require(msg.sender == submitter, "Caller not submitter");
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
        require(deposits[submitter] + msg.value >= MIN_DEPOSIT, "Do not have enough deposit");
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
            originTimestamps[currentBatchSequenced] = block.timestamp;
        }
        lastBatchSequenced = currentBatchSequenced;

        emit SubmitBatches(lastBatchSequenced);
    }

    // challengeState challenges a batch by submitting a deposit.
    function challengeState(uint64 batchIndex) external payable {
        require(challenger == address(0), "Already in challenge");
        // check challenge window
        // todo get finalization period from output oracle
        bool insideChallengeWindow = originTimestamps[batchIndex] + FINALIZATION_PERIOD_SECONDS > block.timestamp;
        require(insideChallengeWindow, "Cannot challenge batch outside the challenge window");

        challenger = msg.sender;
        // check challenge amount
        require(msg.value >= MIN_DEPOSIT);
        deposits[challenger] += msg.value;
        uint64 challengeIndex = ++challengeNum;
        challenges[challengeIndex] = BatchChallenge(msg.sender, batchIndex, msg.value, block.timestamp, false);
        emit ChallengeState(challengeIndex, msg.sender, batchIndex, msg.value);
    }

    // proveState proves a batch by submitting a proof.
    function proveState(uint64 challengeIndex, bytes calldata proof) external{
        // check challenge exists
        require(challenges[challengeIndex].challenger != address(0));
        require(!challenges[challengeIndex].finished);
        bool insideChallengeWindow = challenges[challengeIndex].startTime + PROOF_WINDOW > block.timestamp;
        if (insideChallengeWindow) {
            _challengerWin(challengeIndex,"timeout");
            // todo pause PORTAL contracts
            // PORTAL.pause();
        } else {
            // check proof
            require(proof.length > 0);
            require(msg.sender == submitter);
            require(_verifyProof(proof, commitments[challenges[challengeIndex].batchIndex]),"Prove failed");
            _defenderWin(challengeIndex, "prove success");
        }
        challenges[challengeIndex].finished = true;
    }

    function _defenderWin(uint64 challengeIndex, string memory  _type) internal {
        address challengerAddr = challenges[challengeIndex].challenger;
        uint256 challengeDeposit = challenges[challengeIndex].challengeDeposit;
        uint256 submitterDeposit = deposits[submitter];
        deposits[challengerAddr] -= challengeDeposit;
        deposits[submitter] += challengeDeposit;
        emit ChallengeRes(challengeIndex, submitter, _type);
    }

    function _challengerWin(uint64 challengeIndex, string memory _type) internal {
        address challengerAddr = challenges[challengeIndex].challenger;
        uint256 submitterDeposit = deposits[submitter];
        deposits[submitter] -= MIN_DEPOSIT;
        deposits[challengerAddr] += MIN_DEPOSIT;
        emit ChallengeRes(challengeIndex, challengerAddr, _type);
    }

    function withdraw(uint256 amount) public {
        require(!isInChallenge(msg.sender));
        uint256 value = deposits[msg.sender];
        require(amount > 0);
        uint256 withdrawAmount = amount;
        if (amount + MIN_DEPOSIT > value){
            // 1. check if in challenge
            // 2. check weather sender is submitter
            // 2.1 submitter should confirm all batch then withdraw
            if (msg.sender == submitter){
                require(
                    originTimestamps[lastBatchSequenced]+FINALIZATION_PERIOD_SECONDS <= block.timestamp,
                    "submitter should wait batch to be confirm"
                );
            }
        }
        if (amount > value){
            withdrawAmount = value;
        }
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

    function isInChallenge(address user) public view returns (bool) {
        if (user == submitter){
            if (!challenges[challengeNum-1].finished){
                return true;
            }
        }
        for(uint64 i=0;i<challengeNum;i++){
            if(challenges[i].challenger == user && !challenges[i].finished){
                return true;
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
