package sql

type Explanation struct {
	Order       int
	Action      int
	Name        string
	Description string
	Key         []byte
	Value       []byte
	Cost        int
}
