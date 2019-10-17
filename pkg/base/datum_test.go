package base

import (
	"encoding/hex"
	"github.com/elliotcourant/timber"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDatum_EncodeDecode(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		initial := Value{val: 123}.Datum()
		encoded := initial.Encode()
		timber.Debugf("\n%s", hex.Dump(encoded))
		decoded := Datum{}
		decoded.Decode(encoded)
		assert.Equal(t, initial, decoded)
	})

	t.Run("string", func(t *testing.T) {
		initial := Value{val: "123"}.Datum()
		encoded := initial.Encode()
		timber.Debugf("\n%s", hex.Dump(encoded))
		decoded := Datum{}
		decoded.Decode(encoded)
		assert.Equal(t, initial, decoded)
	})
}
