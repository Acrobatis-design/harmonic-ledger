package core

type TxID string
type DomainID string
type ObjectID string
type Commitment [32]byte

type FeeEnvelope struct {
	MaxFee uint64
}
