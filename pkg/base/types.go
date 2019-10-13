package base

import (
	"fmt"
	"github.com/elliotcourant/buffers"
	"github.com/elliotcourant/melogale/pkg/ast"
	"reflect"
	"strings"
)

type Type struct {
	Family reflect.Kind
	Size   uint8
}

func (t Type) Encode() []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendUint8(uint8(t.Family))
	buf.AppendUint8(t.Size)
	return buf.Bytes()
}

func (t *Type) Decode(src []byte) {
	buf := buffers.NewBytesReader(src)
	t.Family = reflect.Kind(buf.NextUint8())
	t.Size = buf.NextUint8()
}

func GetType(t ast.TypeName) Type {
	typ := Type{
		Family: 0,
		Size:   0,
	}
	name := getName(t)
	switch T(name) {
	case PgInt8:
		typ.Family = reflect.Int64
		typ.Size = 8
	case Text:
		typ.Family = reflect.String
		typ.Size = 255
	default:
		typ.Family = reflect.String
		typ.Size = 255
	}
	return typ
}

func getName(t ast.TypeName) string {
	s := make([]string, len(t.Names.Items))
	for i, n := range t.Names.Items {
		switch t := n.(type) {
		case ast.String:
			s[i] = strings.ToLower(t.Str)
		default:
			panic(fmt.Sprintf("cannot parse type name with item type [%T] in list", t))
		}
	}
	return strings.Join(s, ".")
}
