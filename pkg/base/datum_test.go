package base

import (
	"encoding/hex"
	"github.com/elliotcourant/timber"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestDatum_EncodeDecode(t *testing.T) {
	t.Run("int to string", func(t *testing.T) {
		initial := Datum{
			Type:  reflect.String,
			Value: 123,
		}
		encoded := initial.Encode()
		timber.Debugf("\n%s", hex.Dump(encoded))
		decoded := Datum{}
		decoded.Decode(encoded)
		assert.Equal(t, Datum{
			Type:  reflect.String,
			Value: "123",
		}, decoded)
	})

	t.Run("int to int", func(t *testing.T) {
		initial := Datum{
			Type:  reflect.Int64,
			Value: int64(123),
		}
		encoded := initial.Encode()
		timber.Debugf("\n%s", hex.Dump(encoded))
		decoded := Datum{}
		decoded.Decode(encoded)
		assert.Equal(t, initial, decoded)
	})
}
