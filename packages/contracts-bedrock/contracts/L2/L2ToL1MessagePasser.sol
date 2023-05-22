// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Types } from "../libraries/Types.sol";
import { Hashing } from "../libraries/Hashing.sol";
import { Encoding } from "../libraries/Encoding.sol";
import { Burn } from "../libraries/Burn.sol";
import { Semver } from "../universal/Semver.sol";

/**
 * @custom:proxied
 * @custom:predeploy 0x5300000000000000000000000000000000000000
 * @title L2ToL1MessagePasser
 * @notice The L2ToL1MessagePasser is a dedicated contract where messages that are being sent from
 *         L2 to L1 can be stored. The storage root of this contract is pulled up to the top level
 *         of the L2 output to reduce the cost of proving the existence of sent messages.
 */
contract L2ToL1MessagePasser is Semver {
    /**
     * @notice The L1 gas limit set when eth is withdrawn using the receive() function.
     */
    uint256 internal constant RECEIVE_DEFAULT_GAS_LIMIT = 100_000;

    /**
     * @notice Current message version identifier.
     */
    uint16 public constant MESSAGE_VERSION = 1;

    /**
     * @notice Includes the message hashes for all withdrawals
     */
    mapping(bytes32 => bool) public sentMessages;

    /**
     * @notice A unique value hashed with each withdrawal.
     */
    uint240 internal msgNonce;

    /// @dev The maximum height of the withdraw merkle tree.
    uint256 private constant MAX_TREE_HEIGHT = 40;

    /// @notice The merkle root of the current merkle tree.
    /// @dev This is actual equal to `branches[n]`.
    bytes32 public messageRoot;

    /// @notice The list of zero hash in each height.
    bytes32[MAX_TREE_HEIGHT] private zeroHashes;

    bytes32[MAX_TREE_HEIGHT] public branches;

    /**
     * @notice Emitted any time a withdrawal is initiated.
     *
     * @param nonce          Unique value corresponding to each withdrawal.
     * @param sender         The L2 account address which initiated the withdrawal.
     * @param target         The L1 account address the call will be send to.
     * @param value          The ETH value submitted for withdrawal, to be forwarded to the target.
     * @param gasLimit       The minimum amount of gas that must be provided when withdrawing.
     * @param data           The data to be forwarded to the target on L1.
     * @param withdrawalHash The hash of the withdrawal.
     * @param rootHash       The root hash of the merkle tree after the withdrawal is included.
     */
    event MessagePassed(
        uint256 indexed nonce,
        address indexed sender,
        address indexed target,
        uint256 value,
        uint256 gasLimit,
        bytes data,
        bytes32 withdrawalHash,
        bytes32 rootHash
    );

    /**
     * @notice Emitted when the balance of this contract is burned.
     *
     * @param amount Amount of ETh that was burned.
     */
    event WithdrawerBalanceBurnt(uint256 indexed amount);

    /**
     * @custom:semver 1.0.0
     */
    constructor() Semver(1, 0, 0) {}

    /**
     * @notice Allows users to withdraw ETH by sending directly to this contract.
     */
    receive() external payable {
        initiateWithdrawal(msg.sender, RECEIVE_DEFAULT_GAS_LIMIT, bytes(""));
    }

    /**
     * @notice Removes all ETH held by this contract from the state. Used to prevent the amount of
     *         ETH on L2 inflating when ETH is withdrawn. Currently only way to do this is to
     *         create a contract and self-destruct it to itself. Anyone can call this function. Not
     *         incentivized since this function is very cheap.
     */
    function burn() external {
        uint256 balance = address(this).balance;
        Burn.eth(balance);
        emit WithdrawerBalanceBurnt(balance);
    }

    /**
     * @notice Sends a message from L2 to L1.
     *
     * @param _target   Address to call on L1 execution.
     * @param _gasLimit Minimum gas limit for executing the message on L1.
     * @param _data     Data to forward to L1 target.
     */
    function initiateWithdrawal(
        address _target,
        uint256 _gasLimit,
        bytes memory _data
    ) public payable {
        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: messageNonce(),
                sender: msg.sender,
                target: _target,
                value: msg.value,
                gasLimit: _gasLimit,
                data: _data
            })
        );

        sentMessages[withdrawalHash] = true;

        emit MessagePassed(
            messageNonce(),
            msg.sender,
            _target,
            msg.value,
            _gasLimit,
            _data,
            withdrawalHash,
            rootHash
        );
    }

    /**
     * @notice Retrieves the next message nonce. Message version will be added to the upper two
     *         bytes of the message nonce. Message version allows us to treat messages as having
     *         different structures.
     *
     * @return Nonce of the next message to be sent, with added message version.
     */
    function messageNonce() public view returns (uint256) {
        return Encoding.encodeVersionedNonce(msgNonce, MESSAGE_VERSION);
    }

    function _initializeMerkleTree() internal {
        // Compute hashes in empty sparse Merkle tree
        for (uint256 height = 0; height + 1 < MAX_TREE_HEIGHT; height++) {
            zeroHashes[height + 1] = _efficientHash(zeroHashes[height], zeroHashes[height]);
        }
    }

    function _appendMessageHash(bytes32 _messageHash) internal returns (bytes32) {
        uint240 _currentMessageIndex = msgNonce;
        bytes32 _hash = _messageHash;
        uint256 _height = 0;
        // @todo it can be optimized, since we only need the newly added branch.
        while (_currentMessageIndex != 0) {
            if (_currentMessageIndex % 2 == 0) {
                // it may be used in next round.
                branches[_height] = _hash;
                // it's a left child, the right child must be null
                _hash = _efficientHash(_hash, zeroHashes[_height]);
            } else {
                // it's a right child, use previously computed hash
                _hash = _efficientHash(branches[_height], _hash);
            }
            unchecked {
                _height += 1;
            }
            _currentMessageIndex >>= 1;
        }

        branches[_height] = _hash;
        messageRoot = _hash;

        _currentMessageIndex = msgNonce;
        unchecked {
            msgNonce = _currentMessageIndex + 1;
        }
        return _hash;
    }

    function _efficientHash(bytes32 a, bytes32 b) private pure returns (bytes32 value) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            mstore(0x00, a)
            mstore(0x20, b)
            value := keccak256(0x00, 0x40)
        }
    }
}
