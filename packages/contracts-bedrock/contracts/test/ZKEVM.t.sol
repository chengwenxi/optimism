// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./CommonTest.t.sol";

contract ZKEVMTest is ZKEVM_Initializer {
    uint256 public constant MIN_DEPOSIT = 1 ether;
    uint256 public constant PROOF_WINDOW = 100;

    function test_stake_withdraw() external {
        assertEq(alice, zkevm.submitter());
        vm.deal(alice, 5 * MIN_DEPOSIT);
        vm.startPrank(alice);
        zkevm.stake{value:MIN_DEPOSIT}();
        assertEq(MIN_DEPOSIT, zkevm.deposits(alice));

        zkevm.withdraw(MIN_DEPOSIT);
        assertEq(0, zkevm.deposits(alice));
    }

    function test_stake_revertCallerNotSubmitter() external {
        // bob submit batch: revert with Caller not submitter
        assertEq(alice, zkevm.submitter());
        vm.deal(bob, 5 * MIN_DEPOSIT);
        vm.startPrank(bob);
        vm.expectRevert("Caller not submitter");
        zkevm.stake{value:MIN_DEPOSIT}();
    }

    function test_roullUp_revertSubmitterDeposit() external {
        // batch init
        ZKEVM.BatchData[] memory batches = new ZKEVM.BatchData[](1);
        batches[0] = ZKEVM.BatchData(
                bytes("0"),
                bytes32(0),
                0,
                bytes("0"),
                bytes32(0)
            );

        assertEq(alice, zkevm.submitter());
        // bob submit batch: revert with Caller not submitter
        vm.prank(bob);
        vm.deal(bob, 5 * MIN_DEPOSIT);
        vm.expectRevert("Caller not submitter");
        zkevm.submitBatches(batches);

        // alice submit batch: revert with Insufficient staking amount
        vm.deal(alice, 5 * MIN_DEPOSIT);
        vm.startPrank(alice);
        assertEq(0, zkevm.deposits(alice));
        vm.expectRevert("Insufficient staking amount");
        zkevm.submitBatches(batches);
    }

    function test_challenge_revert() external {
        // alice submit batch 0
        ZKEVM.BatchData[] memory batches = new ZKEVM.BatchData[](1);
        batches[0] = ZKEVM.BatchData(
                bytes("0"),
                bytes32(0),
                0,
                bytes("0"),
                bytes32(0)
            );
        vm.deal(alice, 5 * MIN_DEPOSIT);
        vm.startPrank(alice);
        zkevm.stake{value:MIN_DEPOSIT}();
        zkevm.submitBatches(batches);
        vm.stopPrank();

        // bob challenge batch 1 : revert with Batch not exist
        vm.prank(bob);
        vm.deal(bob, 5 * MIN_DEPOSIT);
        vm.expectRevert("Batch not exist");
        zkevm.challengeState{value:MIN_DEPOSIT}(1);
        // chloe challenge batch 0 : revert with Caller not challenger
        address chloe = address(1024);
        vm.deal(chloe, 5 * MIN_DEPOSIT);
        vm.expectRevert("Caller not challenger");
        vm.prank(chloe);
        zkevm.challengeState{value:MIN_DEPOSIT}(1);
        // bob challenge btach 0 twice: revert with Already has challenge
        vm.prank(bob);
        zkevm.challengeState{value:MIN_DEPOSIT}(0);
        vm.expectRevert("Already has challenge");
        vm.prank(bob);
        zkevm.challengeState{value:MIN_DEPOSIT}(0);

        // withdraw: User is in challenge
        vm.prank(alice);
        vm.expectRevert("User is in challenge");
        zkevm.withdraw(2 ether);
        vm.prank(bob);
        vm.expectRevert("User is in challenge");
        zkevm.withdraw(2 ether);

        // prove time out
        (
            uint64 batchIndex_,
            address challenger_,
            uint256 challengeDeposit_,
            uint256 startTime_,
            bool finished_
        ) = zkevm.challenges(0);
        vm.warp(startTime_+PROOF_WINDOW*2);
        vm.expectEmit(true, true, true, true);
        emit ChallengeRes(0, bob, "timeout");
        vm.prank(bob);
        zkevm.proveState(0,"");
    }

    function test_rollup_challenge() external {
        // alice stake
        assertEq(alice, zkevm.submitter());
        vm.deal(alice, 5 * MIN_DEPOSIT);
        vm.startPrank(alice);
        zkevm.stake{value:MIN_DEPOSIT}();
        assertEq(MIN_DEPOSIT, zkevm.deposits(alice));

        // alice roll up batch 1,2,3,4,5
        ZKEVM.BatchData[] memory batches = new ZKEVM.BatchData[](5);
        for (uint256 i=0 ;i < 5;i++) {
            batches[i] = ZKEVM.BatchData(
                abi.encodePacked(i),
                bytes32(i),
                0,
                abi.encodePacked(i),
                bytes32(i)
            );
        }
        vm.expectEmit(true, true, true, true);
        emit SubmitBatches(uint64(batches.length));
        zkevm.submitBatches(batches);
        // check batch submit success
        assertEq(batches.length, zkevm.lastBatchSequenced());
        uint64 challengeBatchIndex = 3;
        assertFalse(zkevm.isBatchInChallenge(challengeBatchIndex));
        assertTrue(zkevm.isUserInChallenge(alice));
        vm.stopPrank();

        vm.deal(bob, 5 * MIN_DEPOSIT);
        vm.startPrank(bob);
        vm.expectEmit(true, true, true, true);
        emit ChallengeState(challengeBatchIndex, bob, MIN_DEPOSIT);
        zkevm.challengeState{value:MIN_DEPOSIT}(challengeBatchIndex);
        assertEq(MIN_DEPOSIT, zkevm.deposits(bob));
        // bob challenge batch 3
        (
            uint64 batchIndex_,
            address challenger_,
            uint256 challengeDeposit_,
            uint256 startTime_,
            bool finished_
        ) = zkevm.challenges(challengeBatchIndex);
        // chech challenge data
        assertEq(batchIndex_, challengeBatchIndex);
        assertEq(challenger_, bob);
        assertEq(challengeDeposit_, MIN_DEPOSIT);
        assertEq(finished_, false);
        // alice prove challenge
        uint256 bobDeposit = zkevm.deposits(bob);
        uint256 aliceDeposit = zkevm.deposits(alice);

        // submit proof: revert with Invalid proof
        vm.expectRevert("Invalid proof");
        zkevm.proveState(challengeBatchIndex,"");
        // submit proof
        vm.expectEmit(true, true, true, true);
        emit ChallengeRes(challengeBatchIndex, alice, "prove success");
        zkevm.proveState(challengeBatchIndex,bytes("nil"));
        assertEq(zkevm.deposits(bob), 0);
        assertEq(zkevm.deposits(alice), bobDeposit+aliceDeposit);
        // submit proof twice: revert with Challenge already finished
        vm.expectRevert("Challenge already finished");
        zkevm.proveState(challengeBatchIndex,bytes("nil"));
        vm.stopPrank();
    }

}
