package core

type CausalDomain struct {
	ID           DomainID
	PendingTxs   []Transaction
	OrderedTxs   []Transaction
}
