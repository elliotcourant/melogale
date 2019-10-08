package nitrodb

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestDatum_Encode(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		datum := Datum{
			TableId: 123,
			PrimaryKey: [][]byte{
				[]byte("SKU1234"),
			},
			CommitTs: 35612345,
			Values: map[uint8][]byte{
				1: []byte("This is a short description"),
				2: []byte(time.Now().String()),
			},
		}
		encoded := datum.Encode()
		fmt.Println(hex.Dump(encoded))

		decode := Datum{}
		decode.Decode(encoded)
		if !reflect.DeepEqual(datum, decode) {
			panic("not equal")
		}
	})
}
