//go:generate go run ../../../presets/cmd/main.go --presets-dir=../../../presets/configs --output-dir=../../presets/generated

package common

import (
	"github.com/protolambda/ztyp/codec"
	"github.com/protolambda/ztyp/tree"
	"math/big"
)

type Phase0Config struct {

	// Misc.
	MAX_COMMITTEES_PER_SLOT      uint64 `yaml:"MAX_COMMITTEES_PER_SLOT" json:"MAX_COMMITTEES_PER_SLOT"`
	TARGET_COMMITTEE_SIZE        uint64 `yaml:"TARGET_COMMITTEE_SIZE" json:"TARGET_COMMITTEE_SIZE"`
	MAX_VALIDATORS_PER_COMMITTEE uint64 `yaml:"MAX_VALIDATORS_PER_COMMITTEE" json:"MAX_VALIDATORS_PER_COMMITTEE"`
	MIN_PER_EPOCH_CHURN_LIMIT    uint64 `yaml:"MIN_PER_EPOCH_CHURN_LIMIT" json:"MIN_PER_EPOCH_CHURN_LIMIT"`
	CHURN_LIMIT_QUOTIENT         uint64 `yaml:"CHURN_LIMIT_QUOTIENT" json:"CHURN_LIMIT_QUOTIENT"`
	SHUFFLE_ROUND_COUNT          uint8  `yaml:"SHUFFLE_ROUND_COUNT" json:"SHUFFLE_ROUND_COUNT"`

	// Genesis.
	MIN_GENESIS_ACTIVE_VALIDATOR_COUNT uint64    `yaml:"MIN_GENESIS_ACTIVE_VALIDATOR_COUNT" json:"MIN_GENESIS_ACTIVE_VALIDATOR_COUNT"`
	MIN_GENESIS_TIME                   Timestamp `yaml:"MIN_GENESIS_TIME" json:"MIN_GENESIS_TIME"`

	// Balance math
	HYSTERESIS_QUOTIENT            uint64 `yaml:"HYSTERESIS_QUOTIENT" json:"HYSTERESIS_QUOTIENT"`
	HYSTERESIS_DOWNWARD_MULTIPLIER uint64 `yaml:"HYSTERESIS_DOWNWARD_MULTIPLIER" json:"HYSTERESIS_DOWNWARD_MULTIPLIER"`
	HYSTERESIS_UPWARD_MULTIPLIER   uint64 `yaml:"HYSTERESIS_UPWARD_MULTIPLIER" json:"HYSTERESIS_UPWARD_MULTIPLIER"`

	// Fork Choice
	SAFE_SLOTS_TO_UPDATE_JUSTIFIED uint64 `yaml:"SAFE_SLOTS_TO_UPDATE_JUSTIFIED" json:"SAFE_SLOTS_TO_UPDATE_JUSTIFIED"`

	// Validator
	ETH1_FOLLOW_DISTANCE                  uint64 `yaml:"ETH1_FOLLOW_DISTANCE" json:"ETH1_FOLLOW_DISTANCE"`
	TARGET_AGGREGATORS_PER_COMMITTEE      uint64 `yaml:"TARGET_AGGREGATORS_PER_COMMITTEE" json:"TARGET_AGGREGATORS_PER_COMMITTEE"`
	RANDOM_SUBNETS_PER_VALIDATOR          uint64 `yaml:"RANDOM_SUBNETS_PER_VALIDATOR" json:"RANDOM_SUBNETS_PER_VALIDATOR"`
	EPOCHS_PER_RANDOM_SUBNET_SUBSCRIPTION uint64 `yaml:"EPOCHS_PER_RANDOM_SUBNET_SUBSCRIPTION" json:"EPOCHS_PER_RANDOM_SUBNET_SUBSCRIPTION"`
	SECONDS_PER_ETH1_BLOCK                uint64 `yaml:"SECONDS_PER_ETH1_BLOCK" json:"SECONDS_PER_ETH1_BLOCK"`

	// Deposit contract
	DEPOSIT_CHAIN_ID         uint64      `yaml:"DEPOSIT_CHAIN_ID" json:"DEPOSIT_CHAIN_ID"`
	DEPOSIT_NETWORK_ID       uint64      `yaml:"DEPOSIT_NETWORK_ID" json:"DEPOSIT_NETWORK_ID"`
	DEPOSIT_CONTRACT_ADDRESS Eth1Address `yaml:"DEPOSIT_CONTRACT_ADDRESS" json:"DEPOSIT_CONTRACT_ADDRESS"`

	// Gwei values
	MIN_DEPOSIT_AMOUNT          Gwei `yaml:"MIN_DEPOSIT_AMOUNT" json:"MIN_DEPOSIT_AMOUNT"`
	MAX_EFFECTIVE_BALANCE       Gwei `yaml:"MAX_EFFECTIVE_BALANCE" json:"MAX_EFFECTIVE_BALANCE"`
	EJECTION_BALANCE            Gwei `yaml:"EJECTION_BALANCE" json:"EJECTION_BALANCE"`
	EFFECTIVE_BALANCE_INCREMENT Gwei `yaml:"EFFECTIVE_BALANCE_INCREMENT" json:"EFFECTIVE_BALANCE_INCREMENT"`

	// Initial values
	GENESIS_FORK_VERSION  Version          `yaml:"GENESIS_FORK_VERSION" json:"GENESIS_FORK_VERSION"`
	BLS_WITHDRAWAL_PREFIX WithdrawalPrefix `yaml:"BLS_WITHDRAWAL_PREFIX" json:"BLS_WITHDRAWAL_PREFIX"`

	// Time parameters
	GENESIS_DELAY                       Timestamp `yaml:"GENESIS_DELAY" json:"GENESIS_DELAY"`
	SECONDS_PER_SLOT                    Timestamp `yaml:"SECONDS_PER_SLOT" json:"SECONDS_PER_SLOT"`
	MIN_ATTESTATION_INCLUSION_DELAY     Slot      `yaml:"MIN_ATTESTATION_INCLUSION_DELAY" json:"MIN_ATTESTATION_INCLUSION_DELAY"`
	SLOTS_PER_EPOCH                     Slot      `yaml:"SLOTS_PER_EPOCH" json:"SLOTS_PER_EPOCH"`
	MIN_SEED_LOOKAHEAD                  Epoch     `yaml:"MIN_SEED_LOOKAHEAD" json:"MIN_SEED_LOOKAHEAD"`
	MAX_SEED_LOOKAHEAD                  Epoch     `yaml:"MAX_SEED_LOOKAHEAD" json:"MAX_SEED_LOOKAHEAD"`
	EPOCHS_PER_ETH1_VOTING_PERIOD       Epoch     `yaml:"EPOCHS_PER_ETH1_VOTING_PERIOD" json:"EPOCHS_PER_ETH1_VOTING_PERIOD"`
	SLOTS_PER_HISTORICAL_ROOT           Slot      `yaml:"SLOTS_PER_HISTORICAL_ROOT" json:"SLOTS_PER_HISTORICAL_ROOT"`
	MIN_VALIDATOR_WITHDRAWABILITY_DELAY Epoch     `yaml:"MIN_VALIDATOR_WITHDRAWABILITY_DELAY" json:"MIN_VALIDATOR_WITHDRAWABILITY_DELAY"`
	SHARD_COMMITTEE_PERIOD              Epoch     `yaml:"SHARD_COMMITTEE_PERIOD" json:"SHARD_COMMITTEE_PERIOD"`
	MIN_EPOCHS_TO_INACTIVITY_PENALTY    Epoch     `yaml:"MIN_EPOCHS_TO_INACTIVITY_PENALTY" json:"MIN_EPOCHS_TO_INACTIVITY_PENALTY"`

	// State vector lengths
	EPOCHS_PER_HISTORICAL_VECTOR Epoch  `yaml:"EPOCHS_PER_HISTORICAL_VECTOR" json:"EPOCHS_PER_HISTORICAL_VECTOR"`
	EPOCHS_PER_SLASHINGS_VECTOR  Epoch  `yaml:"EPOCHS_PER_SLASHINGS_VECTOR" json:"EPOCHS_PER_SLASHINGS_VECTOR"`
	HISTORICAL_ROOTS_LIMIT       uint64 `yaml:"HISTORICAL_ROOTS_LIMIT" json:"HISTORICAL_ROOTS_LIMIT"`
	VALIDATOR_REGISTRY_LIMIT     uint64 `yaml:"VALIDATOR_REGISTRY_LIMIT" json:"VALIDATOR_REGISTRY_LIMIT"`

	// Reward and penalty quotients
	BASE_REWARD_FACTOR               uint64 `yaml:"BASE_REWARD_FACTOR" json:"BASE_REWARD_FACTOR"`
	WHISTLEBLOWER_REWARD_QUOTIENT    uint64 `yaml:"WHISTLEBLOWER_REWARD_QUOTIENT" json:"WHISTLEBLOWER_REWARD_QUOTIENT"`
	PROPOSER_REWARD_QUOTIENT         uint64 `yaml:"PROPOSER_REWARD_QUOTIENT" json:"PROPOSER_REWARD_QUOTIENT"`
	INACTIVITY_PENALTY_QUOTIENT      uint64 `yaml:"INACTIVITY_PENALTY_QUOTIENT" json:"INACTIVITY_PENALTY_QUOTIENT"`
	MIN_SLASHING_PENALTY_QUOTIENT    uint64 `yaml:"MIN_SLASHING_PENALTY_QUOTIENT" json:"MIN_SLASHING_PENALTY_QUOTIENT"`
	PROPORTIONAL_SLASHING_MULTIPLIER uint64 `yaml:"PROPORTIONAL_SLASHING_MULTIPLIER" json:"PROPORTIONAL_SLASHING_MULTIPLIER"`

	// Max operations per block
	MAX_PROPOSER_SLASHINGS uint64 `yaml:"MAX_PROPOSER_SLASHINGS" json:"MAX_PROPOSER_SLASHINGS"`
	MAX_ATTESTER_SLASHINGS uint64 `yaml:"MAX_ATTESTER_SLASHINGS" json:"MAX_ATTESTER_SLASHINGS"`
	MAX_ATTESTATIONS       uint64 `yaml:"MAX_ATTESTATIONS" json:"MAX_ATTESTATIONS"`
	MAX_DEPOSITS           uint64 `yaml:"MAX_DEPOSITS" json:"MAX_DEPOSITS"`
	MAX_VOLUNTARY_EXITS    uint64 `yaml:"MAX_VOLUNTARY_EXITS" json:"MAX_VOLUNTARY_EXITS"`

	// Signature domains
	DOMAIN_BEACON_PROPOSER     BLSDomainType `yaml:"DOMAIN_BEACON_PROPOSER" json:"DOMAIN_BEACON_PROPOSER"`
	DOMAIN_BEACON_ATTESTER     BLSDomainType `yaml:"DOMAIN_BEACON_ATTESTER" json:"DOMAIN_BEACON_ATTESTER"`
	DOMAIN_RANDAO              BLSDomainType `yaml:"DOMAIN_RANDAO" json:"DOMAIN_RANDAO"`
	DOMAIN_DEPOSIT             BLSDomainType `yaml:"DOMAIN_DEPOSIT" json:"DOMAIN_DEPOSIT"`
	DOMAIN_VOLUNTARY_EXIT      BLSDomainType `yaml:"DOMAIN_VOLUNTARY_EXIT" json:"DOMAIN_VOLUNTARY_EXIT"`
	DOMAIN_SELECTION_PROOF     BLSDomainType `yaml:"DOMAIN_SELECTION_PROOF" json:"DOMAIN_SELECTION_PROOF"`
	DOMAIN_AGGREGATE_AND_PROOF BLSDomainType `yaml:"DOMAIN_AGGREGATE_AND_PROOF" json:"DOMAIN_AGGREGATE_AND_PROOF"`
}

type AltairConfig struct {
	// Updated penalty values
	INACTIVITY_PENALTY_QUOTIENT_ALTAIR      uint64 `yaml:"INACTIVITY_PENALTY_QUOTIENT_ALTAIR" json:"INACTIVITY_PENALTY_QUOTIENT_ALTAIR"`
	MIN_SLASHING_PENALTY_QUOTIENT_ALTAIR    uint64 `yaml:"MIN_SLASHING_PENALTY_QUOTIENT_ALTAIR" json:"MIN_SLASHING_PENALTY_QUOTIENT_ALTAIR"`
	PROPORTIONAL_SLASHING_MULTIPLIER_ALTAIR uint64 `yaml:"PROPORTIONAL_SLASHING_MULTIPLIER_ALTAIR" json:"PROPORTIONAL_SLASHING_MULTIPLIER_ALTAIR"`

	// Misc
	SYNC_COMMITTEE_SIZE    uint64 `yaml:"SYNC_COMMITTEE_SIZE" json:"SYNC_COMMITTEE_SIZE"`
	SYNC_SUBCOMMITTEE_SIZE uint64 `yaml:"SYNC_SUBCOMMITTEE_SIZE" json:"SYNC_SUBCOMMITTEE_SIZE"`

	INACTIVITY_SCORE_BIAS uint64 `yaml:"INACTIVITY_SCORE_BIAS" json:"INACTIVITY_SCORE_BIAS"`

	// Time parameters
	EPOCHS_PER_SYNC_COMMITTEE_PERIOD Epoch `yaml:"EPOCHS_PER_SYNC_COMMITTEE_PERIOD" json:"EPOCHS_PER_SYNC_COMMITTEE_PERIOD"`

	// Sync committees and light clients
	MIN_SYNC_COMMITTEE_PARTICIPANTS uint64 `yaml:"MIN_SYNC_COMMITTEE_PARTICIPANTS" json:"MIN_SYNC_COMMITTEE_PARTICIPANTS"`
	MAX_VALID_LIGHT_CLIENT_UPDATES uint64 `yaml:"MAX_VALID_LIGHT_CLIENT_UPDATES" json:"MAX_VALID_LIGHT_CLIENT_UPDATES"`
	LIGHT_CLIENT_UPDATE_TIMEOUT uint64 `yaml:"LIGHT_CLIENT_UPDATE_TIMEOUT" json:"LIGHT_CLIENT_UPDATE_TIMEOUT"`

	// Domain types
	DOMAIN_SYNC_COMMITTEE BLSDomainType `yaml:"DOMAIN_SYNC_COMMITTEE" json:"DOMAIN_SYNC_COMMITTEE"`

	ALTAIR_FORK_SLOT    Slot    `yaml:"ALTAIR_FORK_SLOT" json:"ALTAIR_FORK_SLOT"`
	ALTAIR_FORK_VERSION Version `yaml:"ALTAIR_FORK_VERSION" json:"ALTAIR_FORK_VERSION"`
}

type Phase1Config struct {
	// phase1-fork
	PHASE_1_FORK_VERSION  Version `yaml:"PHASE_1_FORK_VERSION" json:"PHASE_1_FORK_VERSION"`
	PHASE_1_FORK_SLOT     Slot    `yaml:"PHASE_1_FORK_SLOT" json:"PHASE_1_FORK_SLOT"`
	INITIAL_ACTIVE_SHARDS uint64  `yaml:"INITIAL_ACTIVE_SHARDS" json:"INITIAL_ACTIVE_SHARDS"`

	// beacon-chain
	MAX_SHARDS                      uint64 `yaml:"MAX_SHARDS" json:"MAX_SHARDS"`
	LIGHT_CLIENT_COMMITTEE_SIZE     uint64 `yaml:"LIGHT_CLIENT_COMMITTEE_SIZE" json:"LIGHT_CLIENT_COMMITTEE_SIZE"`
	GASPRICE_ADJUSTMENT_COEFFICIENT uint64 `yaml:"GASPRICE_ADJUSTMENT_COEFFICIENT" json:"GASPRICE_ADJUSTMENT_COEFFICIENT"`

	// Shard block configs
	MAX_SHARD_BLOCK_SIZE             uint64   `yaml:"MAX_SHARD_BLOCK_SIZE" json:"MAX_SHARD_BLOCK_SIZE"`
	TARGET_SHARD_BLOCK_SIZE          uint64   `yaml:"TARGET_SHARD_BLOCK_SIZE" json:"TARGET_SHARD_BLOCK_SIZE"`
	SHARD_BLOCK_OFFSETS              []uint64 `yaml:"SHARD_BLOCK_OFFSETS" json:"SHARD_BLOCK_OFFSETS"`
	MAX_SHARD_BLOCKS_PER_ATTESTATION uint64   `yaml:"MAX_SHARD_BLOCKS_PER_ATTESTATION" json:"MAX_SHARD_BLOCKS_PER_ATTESTATION"`
	BYTES_PER_CUSTODY_CHUNK          uint64   `yaml:"BYTES_PER_CUSTODY_CHUNK" json:"BYTES_PER_CUSTODY_CHUNK"`
	CUSTODY_RESPONSE_DEPTH           uint64   `yaml:"CUSTODY_RESPONSE_DEPTH" json:"CUSTODY_RESPONSE_DEPTH"`

	// Gwei values
	MAX_GASPRICE uint64 `yaml:"MAX_GASPRICE" json:"MAX_GASPRICE"`
	MIN_GASPRICE uint64 `yaml:"MIN_GASPRICE" json:"MIN_GASPRICE"`

	// Time parameters
	ONLINE_PERIOD                 uint64 `yaml:"ONLINE_PERIOD" json:"ONLINE_PERIOD"`
	LIGHT_CLIENT_COMMITTEE_PERIOD uint64 `yaml:"LIGHT_CLIENT_COMMITTEE_PERIOD" json:"LIGHT_CLIENT_COMMITTEE_PERIOD"`

	// Max operations per block
	MAX_CUSTODY_CHUNK_CHALLENGE_RECORDS uint64 `yaml:"MAX_CUSTODY_CHUNK_CHALLENGE_RECORDS" json:"MAX_CUSTODY_CHUNK_CHALLENGE_RECORDS"`

	// Domain types
	DOMAIN_SHARD_PROPOSAL  BLSDomainType `yaml:"DOMAIN_SHARD_PROPOSAL" json:"DOMAIN_SHARD_PROPOSAL"`
	DOMAIN_SHARD_COMMITTEE BLSDomainType `yaml:"DOMAIN_SHARD_COMMITTEE" json:"DOMAIN_SHARD_COMMITTEE"`
	DOMAIN_LIGHT_CLIENT    BLSDomainType `yaml:"DOMAIN_LIGHT_CLIENT" json:"DOMAIN_LIGHT_CLIENT"`

	// custody-game domains
	DOMAIN_CUSTODY_BIT_SLASHING      BLSDomainType `yaml:"DOMAIN_CUSTODY_BIT_SLASHING" json:"DOMAIN_CUSTODY_BIT_SLASHING"`
	DOMAIN_LIGHT_SELECTION_PROOF     BLSDomainType `yaml:"DOMAIN_LIGHT_SELECTION_PROOF" json:"DOMAIN_LIGHT_SELECTION_PROOF"`
	DOMAIN_LIGHT_AGGREGATE_AND_PROOF BLSDomainType `yaml:"DOMAIN_LIGHT_AGGREGATE_AND_PROOF" json:"DOMAIN_LIGHT_AGGREGATE_AND_PROOF"`

	// Custody game
	RANDAO_PENALTY_EPOCHS                          uint64 `yaml:"RANDAO_PENALTY_EPOCHS" json:"RANDAO_PENALTY_EPOCHS"`
	EARLY_DERIVED_SECRET_PENALTY_MAX_FUTURE_EPOCHS uint64 `yaml:"EARLY_DERIVED_SECRET_PENALTY_MAX_FUTURE_EPOCHS" json:"EARLY_DERIVED_SECRET_PENALTY_MAX_FUTURE_EPOCHS"`
	EPOCHS_PER_CUSTODY_PERIOD                      uint64 `yaml:"EPOCHS_PER_CUSTODY_PERIOD" json:"EPOCHS_PER_CUSTODY_PERIOD"`
	CUSTODY_PERIOD_TO_RANDAO_PADDING               uint64 `yaml:"CUSTODY_PERIOD_TO_RANDAO_PADDING" json:"CUSTODY_PERIOD_TO_RANDAO_PADDING"`
	MAX_CHUNK_CHALLENGE_DELAY                      uint64 `yaml:"MAX_CHUNK_CHALLENGE_DELAY" json:"MAX_CHUNK_CHALLENGE_DELAY"`

	CUSTODY_PRIME                *big.Int `yaml:"CUSTODY_PRIME" json:"CUSTODY_PRIME"`
	CUSTODY_SECRETS              uint64   `yaml:"CUSTODY_SECRETS" json:"CUSTODY_SECRETS"`
	BYTES_PER_CUSTODY_ATOM       uint64   `yaml:"BYTES_PER_CUSTODY_ATOM" json:"BYTES_PER_CUSTODY_ATOM"`
	CUSTODY_PROBABILITY_EXPONENT uint64   `yaml:"CUSTODY_PROBABILITY_EXPONENT" json:"CUSTODY_PROBABILITY_EXPONENT"`

	// Max operations
	MAX_CUSTODY_KEY_REVEALS          uint64 `yaml:"MAX_CUSTODY_KEY_REVEALS" json:"MAX_CUSTODY_KEY_REVEALS"`
	MAX_EARLY_DERIVED_SECRET_REVEALS uint64 `yaml:"MAX_EARLY_DERIVED_SECRET_REVEALS" json:"MAX_EARLY_DERIVED_SECRET_REVEALS"`
	MAX_CUSTODY_CHUNK_CHALLENGES     uint64 `yaml:"MAX_CUSTODY_CHUNK_CHALLENGES" json:"MAX_CUSTODY_CHUNK_CHALLENGES"`
	MAX_CUSTODY_CHUNK_CHALLENGE_RESP uint64 `yaml:"MAX_CUSTODY_CHUNK_CHALLENGE_RESP" json:"MAX_CUSTODY_CHUNK_CHALLENGE_RESP"`
	MAX_CUSTODY_SLASHINGS            uint64 `yaml:"MAX_CUSTODY_SLASHINGS" json:"MAX_CUSTODY_SLASHINGS"`

	// Reward and penalty quotients
	EARLY_DERIVED_SECRET_REVEAL_SLOT_REWARD_MULTIPLE uint64 `yaml:"EARLY_DERIVED_SECRET_REVEAL_SLOT_REWARD_MULTIPLE" json:"EARLY_DERIVED_SECRET_REVEAL_SLOT_REWARD_MULTIPLE"`
	MINOR_REWARD_QUOTIENT                            uint64 `yaml:"MINOR_REWARD_QUOTIENT" json:"MINOR_REWARD_QUOTIENT"`
}

type SpecObj interface {
	Deserialize(spec *Spec, dr *codec.DecodingReader) error
	Serialize(spec *Spec, w *codec.EncodingWriter) error
	ByteLength(spec *Spec) uint64
	HashTreeRoot(spec *Spec, h tree.HashFn) Root
	FixedLength(spec *Spec) uint64
}

type SSZObj interface {
	codec.Serializable
	codec.Deserializable
	codec.FixedLength
	tree.HTR
}

type specObj struct {
	spec *Spec
	des  SpecObj
}

func (s specObj) Deserialize(dr *codec.DecodingReader) error {
	return s.des.Deserialize(s.spec, dr)
}

func (s specObj) Serialize(w *codec.EncodingWriter) error {
	return s.des.Serialize(s.spec, w)
}

func (s specObj) ByteLength() uint64 {
	return s.des.ByteLength(s.spec)
}

func (s specObj) HashTreeRoot(h tree.HashFn) Root {
	return s.des.HashTreeRoot(s.spec, h)
}

func (s specObj) FixedLength() uint64 {
	return s.des.FixedLength(s.spec)
}

type Spec struct {
	CONFIG_NAME  string `yaml:"CONFIG_NAME,omitempty"`
	Phase0Config `yaml:",inline"`
	AltairConfig `yaml:",inline"`
	Phase1Config `yaml:",inline"`
}

func (spec *Spec) Wrap(des SpecObj) SSZObj {
	return specObj{spec, des}
}

func (spec *Spec) ForkVersion(slot Slot) Version {
	return spec.GENESIS_FORK_VERSION
	// TODO more forks
	//if slot < spec.ALTAIR_FORK_SLOT {
	//	return spec.GENESIS_FORK_VERSION
	//}
	//if slot < spec.PHASE_1_FORK_SLOT {
	//	return spec.PHASE_1_FORK_VERSION
	//}
	//return spec.ALTAIR_FORK_VERSION
}
