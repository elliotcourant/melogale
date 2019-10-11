package store

import (
	"github.com/dgraph-io/badger"
)

var _ Transaction = &transactionBase{}
var _ Iterator = &iteratorBase{}

type Transaction interface {
	Iterator() Iterator
	Set(key, value []byte) error
	Get(key []byte) ([]byte, error)
	NewObjectId(objectName []byte) (uint64, error)
}

type Iterator interface {
	Seek(prefix []byte)
	Next()
	Rewind()
	Valid() bool
	ValidForPrefix(prefix []byte) bool
	Value() (key []byte, value []byte, err error)
}

type transactionBase struct {
	db  *storeBase
	txn *badger.Txn
	itr *iteratorBase
}

func (t *transactionBase) Set(key, value []byte) error {
	return t.txn.Set(key, value)
}

func (t *transactionBase) Get(key []byte) ([]byte, error) {
	item, err := t.txn.Get(key)
	if err == badger.ErrKeyNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return item.ValueCopy(nil)
}

func (t *transactionBase) Iterator() Iterator {
	itr := t.itr
	if itr == nil {
		itr = newIterator(t.txn.NewIterator(badger.IteratorOptions{
			PrefetchValues: true,
			PrefetchSize:   10,
			Reverse:        false,
			AllVersions:    false,
			Prefix:         nil,
			InternalAccess: false,
		}))
		t.itr = itr
	}
	return itr
}

func (t *transactionBase) NewObjectId(objectName []byte) (uint64, error) {
	panic("implement me")
}

type iteratorBase struct {
	itr *badger.Iterator
}

func newIterator(itr *badger.Iterator) *iteratorBase {
	return &iteratorBase{itr: itr}
}

func (i *iteratorBase) Seek(prefix []byte) {
	i.itr.Seek(prefix)
}

func (i *iteratorBase) Next() {
	i.itr.Next()
}

func (i *iteratorBase) Rewind() {
	i.itr.Rewind()
}

func (i *iteratorBase) Valid() bool {
	return i.itr.Valid()
}

func (i *iteratorBase) ValidForPrefix(prefix []byte) bool {
	return i.itr.ValidForPrefix(prefix)
}

func (i *iteratorBase) Value() ([]byte, []byte, error) {
	itm := i.itr.Item()
	if itm == nil {
		return nil, nil, nil
	}
	key := itm.KeyCopy(nil)
	value, err := itm.ValueCopy(nil)
	return key, value, err
}
