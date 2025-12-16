package core

type StateObject struct {
	ID      ObjectID
	Version uint64
	Data    []byte
}
