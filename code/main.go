package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"time"
)

type TxID string
type ObjectID string
type DomainID string
type Commitment [32]byte

type FeeEnvelope struct{ MaxFee uint64 }

type Transaction struct {
	ID         TxID
	ReadSet    []ObjectID
	WriteSet   []ObjectID
	Fee        FeeEnvelope
	Commitment Commitment
	Ciphertext []byte // mock encrypted payload
}

type Receipt struct {
	TxID       TxID
	Commitment Commitment
	Domain     DomainID
}

type StateObject struct {
	ID      ObjectID
	Version uint64
	Balance int64
}

type CausalDomain struct {
	ID         DomainID
	Pending    []Transaction
	Ordered    []Transaction
	DACommit   string
	DAReady    bool
	Finalized  bool
	FinalizedH int64
}

// ---- Admission (ciphertext only) ----

func admit(tx Transaction) (Receipt, error) {
	if len(tx.ReadSet) == 0 || len(tx.WriteSet) == 0 {
		return Receipt{}, fmt.Errorf("missing declared read/write sets")
	}
	if tx.Fee.MaxFee == 0 {
		return Receipt{}, fmt.Errorf("fee envelope missing")
	}
	did := assignDomain(tx.WriteSet)
	return Receipt{TxID: tx.ID, Commitment: tx.Commitment, Domain: did}, nil
}

func assignDomain(writeSet []ObjectID) DomainID {
	ids := append([]ObjectID(nil), writeSet...)
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	h := sha256.New()
	for _, id := range ids {
		h.Write([]byte(id))
		h.Write([]byte{0})
	}
	return DomainID("cd_" + hex.EncodeToString(h.Sum(nil))[:10])
}

// ---- Ordering (commitment-based) ----

func order(cd *CausalDomain) {
	txs := append([]Transaction(nil), cd.Pending...)
	sort.Slice(txs, func(i, j int) bool {
		// deterministic: by commitment bytes (no plaintext)
		return hex.EncodeToString(txs[i].Commitment[:]) < hex.EncodeToString(txs[j].Commitment[:])
	})
	cd.Ordered = txs
}

// ---- Reveal (mock threshold decrypt) ----

func reveal(tx Transaction) ([]byte, error) {
	// In real system: threshold decryption AFTER ordering is fixed.
	// Here: ciphertext is just plaintext bytes for demo.
	if len(tx.Ciphertext) == 0 {
		return nil, fmt.Errorf("empty ciphertext")
	}
	return tx.Ciphertext, nil
}

// ---- Execution (deterministic, explicit sets) ----

func execute(tx Transaction, payload []byte, state map[ObjectID]*StateObject) error {
	// payload format: "Alice->Bob:10"
	var from, to string
	var amt int64
	_, err := fmt.Sscanf(string(payload), "%[^-]->%[^:]:%d", &from, &to, &amt)
	if err != nil {
		return fmt.Errorf("bad payload: %w", err)
	}

	fromID := ObjectID("acct:" + from)
	toID := ObjectID("acct:" + to)

	// enforce declared write set contains both (simple rule for demo)
	if !contains(tx.WriteSet, fromID) || !contains(tx.WriteSet, toID) {
		return fmt.Errorf("write set does not match payload objects")
	}

	sf := state[fromID]
	st := state[toID]
	if sf == nil || st == nil {
		return fmt.Errorf("unknown account")
	}
	if sf.Balance < amt {
		return fmt.Errorf("insufficient funds")
	}

	sf.Balance -= amt
	sf.Version++
	st.Balance += amt
	st.Version++
	return nil
}

func contains(set []ObjectID, x ObjectID) bool {
	for _, v := range set {
		if v == x {
			return true
		}
	}
	return false
}

// ---- DA (publish + check) ----

func daPublish(cd *CausalDomain) {
	// commit to ordered tx commitments
	h := sha256.New()
	for _, tx := range cd.Ordered {
		h.Write(tx.Commitment[:])
	}
	cd.DACommit = hex.EncodeToString(h.Sum(nil))
	cd.DAReady = true
}

func daIsAvailable(cd *CausalDomain) bool { return cd.DAReady }

// ---- Finality (gated by DA) ----

func finalize(cd *CausalDomain, height int64) error {
	if !daIsAvailable(cd) {
		return fmt.Errorf("DA unavailable: finality forbidden")
	}
	cd.Finalized = true
	cd.FinalizedH = height
	return nil
}

// ---- Demo "node" ----

func commit(txPlain string, readSet, writeSet []ObjectID, maxFee uint64) Transaction {
	c := sha256.Sum256([]byte(txPlain + "|" + time.Now().UTC().String()))
	return Transaction{
		ID:         TxID(hex.EncodeToString(c[:8])),
		ReadSet:    readSet,
		WriteSet:   writeSet,
		Fee:        FeeEnvelope{MaxFee: maxFee},
		Commitment: c,
		Ciphertext: []byte(txPlain), // mock: ciphertext=plaintext
	}
}

func main() {
	// initial state
	state := map[ObjectID]*StateObject{
		"acct:Alice": {ID: "acct:Alice", Version: 1, Balance: 100},
		"acct:Bob":   {ID: "acct:Bob", Version: 1, Balance: 5},
	}

	// build one tx
	tx := commit(
		"Alice->Bob:10",
		[]ObjectID{"acct:Alice", "acct:Bob"},
		[]ObjectID{"acct:Alice", "acct:Bob"},
		1,
	)

	// pipeline
	rcpt, err := admit(tx)
	if err != nil {
		panic(err)
	}
	cd := &CausalDomain{ID: rcpt.Domain, Pending: []Transaction{tx}}

	order(cd)

	for _, otx := range cd.Ordered {
		pl, err := reveal(otx)
		if err != nil {
			panic(err)
		}
		if err := execute(otx, pl, state); err != nil {
			fmt.Println("EXEC FAIL:", err)
		}
	}

	daPublish(cd)

	if err := finalize(cd, 1); err != nil {
		panic(err)
	}

	fmt.Println("FINALIZED:", cd.Finalized, "DACommit:", cd.DACommit[:16]+"...")
	fmt.Println("Alice:", state["acct:Alice"].Balance, "Bob:", state["acct:Bob"].Balance)
}
