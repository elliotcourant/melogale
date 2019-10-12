package base

import (
	"encoding/hex"
	"github.com/elliotcourant/timber"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTable_EncodeDecode(t *testing.T) {
	initial := Table{
		TableId: 1,
		Name:    "users",
		Columns: map[string]Column{
			"id": {
				ColumnId: 1,
				Name:     "id",
				Type: Type{
					Family: IntFamily,
					Size:   8,
				},
				Flags: ColumnPrimaryKey,
			},
			"email": {
				ColumnId: 2,
				Name:     "email",
				Type: Type{
					Family: StringFamily,
					Size:   255,
				},
				Flags: ColumnUnique | ColumnIndexed,
			},
			"password": {
				ColumnId: 3,
				Name:     "password",
				Type: Type{
					Family: StringFamily,
					Size:   50,
				},
				Flags: ColumnIndexed,
			},
		},
		Indexes: map[string]Index{
			"ix_users_email_password": {
				IndexId: 1,
				Name:    "ix_users_email_password",
				Columns: map[uint8]uint8{
					2: 0,
					3: 1,
				},
			},
		},
	}
	encoded := initial.EncodeValue()
	timber.Debugf("\n%s", hex.Dump(encoded))
	decoded := Table{}
	decoded.DecodeValue(encoded)
	assert.Equal(t, initial, decoded)
}
