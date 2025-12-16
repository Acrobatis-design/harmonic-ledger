package node

import (
	"harmonic-ledger/code/core"
	"harmonic-ledger/code/modules/admission"
	"harmonic-ledger/code/modules/ordering"
	"harmonic-ledger/code/modules/execution"
	"harmonic-ledger/code/modules/finality"
)

func Process(tx core.Transaction) error {
	receipt, _ := admission.Admit(tx)

	domain := loadDomain(receipt.Domain)

	ordering.Order(domain)

	for _, tx := range domain.OrderedTxs {
		execution.Execute(tx, loadState())
	}

	return finality.Finalize(*domain)
}
