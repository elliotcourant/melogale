// +build badger

package engine

import (
	"github.com/dgraph-io/badger"
	"github.com/elliotcourant/timber"
	"time"
)

var _ Store = &badgerStore{}
var _ Transaction = &badgerTransaction{}
var _ Iterator = &badgerIterator{}

func newStore(options Options) (Store, error) {
	bOptions := badger.DefaultOptions(options.Directory)
	db, err := badger.Open(bOptions)
	return &badgerStore{
		db: db,
	}, err
}

type badgerStore struct {
	db *badger.DB
}

func (b *badgerStore) Begin() (Transaction, error) {
	return &badgerTransaction{
		db:  b,
		txn: b.db.NewTransaction(true),
	}, nil
}

type badgerTransaction struct {
	db  *badgerStore
	txn *badger.Txn
}

func (b *badgerTransaction) Get(key []byte) ([]byte, error) {
	itm, err := b.txn.Get(key)
	if err != nil {
		return nil, err
	}
	return itm.ValueCopy(nil)
}

func (b *badgerTransaction) Set(key, value []byte) error {
	return b.txn.Set(key, value)
}

func (b *badgerTransaction) Iterator() Iterator {
	return &badgerIterator{
		itr: b.txn.NewIterator(badger.IteratorOptions{
			PrefetchValues: true,
			PrefetchSize:   10,
			Reverse:        false,
			AllVersions:    false,
			Prefix:         nil,
			InternalAccess: false,
		}),
	}
}

func (b *badgerTransaction) Commit() error {
	start := time.Now()
	defer timber.Tracef("commit took %s", time.Since(start))
	return b.txn.Commit()
}

func (b *badgerTransaction) Rollback() error {
	b.txn.Discard()
	return nil
}

type badgerIterator struct {
	itr *badger.Iterator
}

func (b *badgerIterator) Close() {
	b.itr.Close()
}

func (b *badgerIterator) Next() {
	b.itr.Next()
}

func (b *badgerIterator) Rewind() {
	b.itr.Rewind()
}

func (b *badgerIterator) Item() (key, value []byte, err error) {
	itm := b.itr.Item()
	key = itm.KeyCopy(nil)
	val, err := itm.ValueCopy(nil)
	return key, val, err
}

func (b *badgerIterator) Seek(prefix []byte) {
	b.itr.Seek(prefix)
}

func (b *badgerIterator) ValidForPrefix(prefix []byte) bool {
	return b.itr.ValidForPrefix(prefix)
}

func (b *badgerIterator) Valid() bool {
	return b.itr.Valid()
}
