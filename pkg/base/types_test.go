package base

import (
	"encoding/hex"
	"github.com/elliotcourant/timber"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestType_EncodeDecode(t *testing.T) {
	initial := Type{
		Family: IntFamily,
		Size:   8,
	}
	encoded := initial.Encode()
	timber.Debugf("\n%s", hex.Dump(encoded))
	decoded := Type{}
	decoded.Decode(encoded)
	assert.Equal(t, initial, decoded)
}