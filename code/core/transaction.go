package core

type Transaction struct {
	ID          TxID
	ReadSet     []ObjectID
	WriteSet    []ObjectID
	Fee         FeeEnvelope
	Commitment  Commitment
	Ciphertext  []byte
}
