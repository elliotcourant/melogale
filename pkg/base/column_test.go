package base

import (
	"encoding/hex"
	"github.com/elliotcourant/timber"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColumn_EncodeDecode(t *testing.T) {
	initial := Column{
		ColumnId: 123,
		Name:     "account_id",
		Type: Type{
			Family: IntFamily,
			Size:   8,
		},
		Flags: ColumnPrimaryKey | ColumnIndexed,
	}
	encoded := initial.Encode()
	timber.Debugf("\n%s", hex.Dump(encoded))
	decoded := Column{}
	decoded.Decode(encoded)
	assert.Equal(t, initial, decoded)
}
