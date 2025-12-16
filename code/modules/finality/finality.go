package finality

import (
	"errors"
	"harmonic-ledger/code/core"
	"harmonic-ledger/code/modules/da"
)

func Finalize(domain core.CausalDomain) error {
	if !da.IsAvailable(domain) {
		return errors.New("data unavailable: finality forbidden")
	}
	// Seal state
	return nil
}
