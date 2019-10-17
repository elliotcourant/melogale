package sql

type Explanation struct {
	Action      int
	Name        string
	Description string
	Key         []byte
	Value       []byte
	Cost        int
}
