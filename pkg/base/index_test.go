package base

import (
	"encoding/hex"
	"github.com/elliotcourant/timber"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndex_EncodeDecode(t *testing.T) {
	t.Run("single column", func(t *testing.T) {
		initial := Index{
			IndexId: 123,
			Name:    "ix_accounts_name",
			Columns: map[uint8]uint8{
				1: 0,
			},
		}
		encoded := initial.Encode()
		timber.Debugf("\n%s", hex.Dump(encoded))
		decoded := Index{}
		decoded.Decode(encoded)
		assert.Equal(t, initial, decoded)
	})

	t.Run("multi column", func(t *testing.T) {
		initial := Index{
			IndexId: 123,
			Name:    "ix_users_email_password",
			Columns: map[uint8]uint8{
				1: 0,
				2: 1,
			},
		}
		encoded := initial.Encode()
		timber.Debugf("\n%s", hex.Dump(encoded))
		decoded := Index{}
		decoded.Decode(encoded)
		assert.Equal(t, initial, decoded)
	})
}
