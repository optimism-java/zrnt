//go:generate go run ../../presets/cmd/main.go --presets-dir=../../presets/eth2.0-presets --output-dir=../../presets/generated

package core

import (
	generated "github.com/protolambda/zrnt/presets/generated"
)

// Helper to know what preset is active
const PRESET_NAME = generated.PRESET_NAME

// Misc.
const SHARD_COUNT Shard = generated.SHARD_COUNT
const TARGET_COMMITTEE_SIZE = generated.TARGET_COMMITTEE_SIZE
const MAX_VALIDATORS_PER_COMMITTEE = generated.MAX_VALIDATORS_PER_COMMITTEE
const MIN_PER_EPOCH_CHURN_LIMIT = generated.MIN_PER_EPOCH_CHURN_LIMIT
const CHURN_LIMIT_QUOTIENT = generated.CHURN_LIMIT_QUOTIENT
const SHUFFLE_ROUND_COUNT uint8 = generated.SHUFFLE_ROUND_COUNT

// Genesis
const MIN_GENESIS_ACTIVE_VALIDATOR_COUNT = generated.MIN_GENESIS_ACTIVE_VALIDATOR_COUNT
const MIN_GENESIS_TIME = generated.MIN_GENESIS_TIME

// Gwei values
const MIN_DEPOSIT_AMOUNT Gwei = generated.MIN_DEPOSIT_AMOUNT
const MAX_EFFECTIVE_BALANCE Gwei = generated.MAX_EFFECTIVE_BALANCE

// unused const FORK_CHOICE_BALANCE_INCREMENT Gwei = generated.FORK_CHOICE_BALANCE_INCREMENT
const EJECTION_BALANCE Gwei = generated.EJECTION_BALANCE

const EFFECTIVE_BALANCE_INCREMENT Gwei = generated.EFFECTIVE_BALANCE_INCREMENT
const HALF_INCREMENT = generated.EFFECTIVE_BALANCE_INCREMENT / 2

// Initial values
const GENESIS_SLOT Slot = generated.GENESIS_SLOT
const GENESIS_EPOCH = Epoch(GENESIS_SLOT / SLOTS_PER_EPOCH)

const BLS_WITHDRAWAL_PREFIX byte = generated.BLS_WITHDRAWAL_PREFIX

// Time parameters
const SECONDS_PER_SLOT Timestamp = generated.SECONDS_PER_SLOT
const MIN_ATTESTATION_INCLUSION_DELAY Slot = generated.MIN_ATTESTATION_INCLUSION_DELAY
const SLOTS_PER_EPOCH Slot = generated.SLOTS_PER_EPOCH
const MIN_SEED_LOOKAHEAD Epoch = generated.MIN_SEED_LOOKAHEAD
const ACTIVATION_EXIT_DELAY Epoch = generated.ACTIVATION_EXIT_DELAY
const SLOTS_PER_ETH1_VOTING_PERIOD Slot = generated.SLOTS_PER_ETH1_VOTING_PERIOD
const SLOTS_PER_HISTORICAL_ROOT Slot = generated.SLOTS_PER_HISTORICAL_ROOT
const MIN_VALIDATOR_WITHDRAWABILITY_DELAY Epoch = generated.MIN_VALIDATOR_WITHDRAWABILITY_DELAY
const PERSISTENT_COMMITTEE_PERIOD Epoch = generated.PERSISTENT_COMMITTEE_PERIOD
const MAX_EPOCHS_PER_CROSSLINK Epoch = generated.MAX_EPOCHS_PER_CROSSLINK
const MIN_EPOCHS_TO_INACTIVITY_PENALTY = generated.MIN_EPOCHS_TO_INACTIVITY_PENALTY

// State list lengths
const EPOCHS_PER_HISTORICAL_VECTOR Epoch = generated.EPOCHS_PER_HISTORICAL_VECTOR
const EPOCHS_PER_SLASHINGS_VECTOR Epoch = generated.EPOCHS_PER_SLASHINGS_VECTOR
const HISTORICAL_ROOTS_LIMIT = generated.HISTORICAL_ROOTS_LIMIT
const VALIDATOR_REGISTRY_LIMIT = generated.VALIDATOR_REGISTRY_LIMIT

// Reward and penalty quotients
const BASE_REWARD_FACTOR = generated.BASE_REWARD_FACTOR
const WHISTLEBLOWER_REWARD_QUOTIENT = generated.WHISTLEBLOWER_REWARD_QUOTIENT
const PROPOSER_REWARD_QUOTIENT = generated.PROPOSER_REWARD_QUOTIENT
const INACTIVITY_PENALTY_QUOTIENT = generated.INACTIVITY_PENALTY_QUOTIENT
const MIN_SLASHING_PENALTY_QUOTIENT = generated.MIN_SLASHING_PENALTY_QUOTIENT

// Max transactions per block
const MAX_PROPOSER_SLASHINGS = generated.MAX_PROPOSER_SLASHINGS
const MAX_ATTESTER_SLASHINGS = generated.MAX_ATTESTER_SLASHINGS
const MAX_ATTESTATIONS = generated.MAX_ATTESTATIONS
const MAX_DEPOSITS = generated.MAX_DEPOSITS
const MAX_VOLUNTARY_EXITS = generated.MAX_VOLUNTARY_EXITS
const MAX_TRANSFERS = generated.MAX_TRANSFERS

// Signature domains
const (
	DOMAIN_BEACON_PROPOSER BLSDomainType = generated.DOMAIN_BEACON_PROPOSER
	DOMAIN_RANDAO          BLSDomainType = generated.DOMAIN_RANDAO
	DOMAIN_ATTESTATION     BLSDomainType = generated.DOMAIN_ATTESTATION
	DOMAIN_DEPOSIT         BLSDomainType = generated.DOMAIN_DEPOSIT
	DOMAIN_VOLUNTARY_EXIT  BLSDomainType = generated.DOMAIN_VOLUNTARY_EXIT
	DOMAIN_TRANSFER        BLSDomainType = generated.DOMAIN_TRANSFER
)

const maxCommitteesPerSlot = uint64(SHARD_COUNT) / uint64(SLOTS_PER_EPOCH)

// Return the number of committees in one epoch.
func CommitteeCount(activeValidators uint64) uint64 {
	validatorsPerSlot := activeValidators / uint64(SLOTS_PER_EPOCH)
	committeesPerSlot := validatorsPerSlot / TARGET_COMMITTEE_SIZE
	if maxCommitteesPerSlot < committeesPerSlot {
		committeesPerSlot = maxCommitteesPerSlot
	}
	if committeesPerSlot == 0 {
		committeesPerSlot = 1
	}
	return committeesPerSlot * uint64(SLOTS_PER_EPOCH)
}
