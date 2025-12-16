package admission

import "harmonic-ledger/code/core"

type Receipt struct {
	TxID       core.TxID
	Commitment core.Commitment
	Domain     core.DomainID
}

func Admit(tx core.Transaction) (Receipt, error) {
	// Validate declared read/write sets
	// Validate fee envelope
	// DO NOT decrypt
	return Receipt{
		TxID:       tx.ID,
		Commitment: tx.Commitment,
		Domain:     assignDomain(tx),
	}, nil
}

func assignDomain(tx core.Transaction) core.DomainID {
	// Deterministic hash over write set
	return core.DomainID("domain-hash-placeholder")
}
