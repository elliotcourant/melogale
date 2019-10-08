package store

import (
	"github.com/dgraph-io/badger"
)

var _ Store = &storeBase{}

type Store interface {
	Begin() Transaction
}

type storeBase struct {
	db *badger.DB
}

func (s *storeBase) Begin() Transaction {
	return nil
}
