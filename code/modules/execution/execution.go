package execution

import "harmonic-ledger/code/core"

func Execute(tx core.Transaction, state map[core.ObjectID]core.StateObject) error {
	// Decrypt AFTER ordering (mocked)
	// Check conflicts
	// Apply state transition
	return nil
}
