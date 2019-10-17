package base

import (
	"encoding/hex"
	"github.com/elliotcourant/timber"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRow_EncodeDecode(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		initial := Row{
			TableId: 123,
			PrimaryKey: []Datum{
				Value{1}.Datum(),
			},
			Datums: map[uint8]Datum{
				1: Value{1}.Datum(),
				2: Value{"email@email.com"}.Datum(),
				3: Value{"password"}.Datum(),
			},
		}
		key, value := initial.EncodeKey(), initial.EncodeValue()
		timber.Debugf("key\n%s", hex.Dump(key))
		timber.Debugf("value\n%s", hex.Dump(value))
		decoded := Row{}
		decoded.DecodeKey(key)
		decoded.DecodeValue(value)
		assert.Equal(t, initial, decoded)
	})
}
