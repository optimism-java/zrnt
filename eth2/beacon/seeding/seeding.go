package seeding

import (
	"encoding/binary"
	. "github.com/protolambda/zrnt/eth2/core"
	. "github.com/protolambda/zrnt/eth2/meta"
	. "github.com/protolambda/zrnt/eth2/util/hashing"
)

type SeedFeature struct {
	Meta interface {
		RandomnessMeta
		ActiveIndexRootMeta
	}
}

// Generate a seed for the given epoch
func (feat *SeedFeature) GetSeed(epoch Epoch) Root {
	buf := make([]byte, 32*3)
	mix := feat.Meta.GetRandomMix(epoch + EPOCHS_PER_HISTORICAL_VECTOR - MIN_SEED_LOOKAHEAD) // Avoid underflow
	copy(buf[0:32], mix[:])
	activeIndexRoot := feat.Meta.GetActiveIndexRoot(epoch)
	copy(buf[32:64], activeIndexRoot[:])
	binary.LittleEndian.PutUint64(buf[64:], uint64(epoch))
	return Hash(buf)
}
