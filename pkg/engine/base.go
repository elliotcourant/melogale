package engine

type Options struct {
	Directory string
}

func NewStore(options Options) (Store, error) {
	return newStore(options)
}

type Store interface {
	Begin() (Transaction, error)
}

type Transaction interface {
	Get(key []byte) ([]byte, error)
	Set(key, value []byte) error
	Iterator() Iterator

	Commit() error
	Rollback() error
}

type Iterator interface {
	Close()
	Next()
	Rewind()
	Item() (key, value []byte, err error)
	Seek(prefix []byte)
	ValidForPrefix(prefix []byte) bool
	Valid() bool
}
