package base

import (
	"fmt"
	"github.com/elliotcourant/buffers"
	"reflect"
)

type Datum struct {
	Type  reflect.Kind
	Value interface{}
}

func (d Datum) Encode() []byte {
	buf := buffers.NewBytesBuffer()
	buf.AppendUint8(uint8(d.Type))
	val := reflect.ValueOf(d.Value)
	switch d.Type {
	case reflect.String:
		buf.AppendString(fmt.Sprintf("%v", d.Value))
	case reflect.Int64:
		buf.AppendInt64(val.Int())
	}
	return buf.Bytes()
}

func (d *Datum) Decode(src []byte) {
	buf := buffers.NewBytesReader(src)
	d.Type = reflect.Kind(buf.NextUint8())
	switch d.Type {
	case reflect.String:
		d.Value = buf.NextString()
	case reflect.Int64:
		d.Value = buf.NextInt64()
	}
}
