package base

import (
	"encoding/hex"
	"github.com/elliotcourant/timber"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestRow_EncodeDecode(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		initial := Row{
			TableId: 123,
			PrimaryKey: []Datum{
				{
					Type:  reflect.Int64,
					Value: int64(1),
				},
			},
			Datums: map[uint8]Datum{
				1: {
					Type:  reflect.Int64,
					Value: int64(1),
				},
				2: {
					Type:  reflect.String,
					Value: "email@email.com",
				},
				3: {
					Type:  reflect.String,
					Value: "password",
				},
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
