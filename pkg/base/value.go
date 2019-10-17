package base

import (
	"encoding/binary"
	"reflect"
)

type Value struct {
	val interface{}
}

func (v Value) Datum() Datum {
	value := reflect.ValueOf(v.val)
	datum := Datum{}
	switch value.Kind() {
	case reflect.String:
		datum = Datum(value.String())
	case reflect.Bool:
		b := 0
		if value.Bool() {
			b = 1
		}
		datum = Datum{uint8(b)}
	case reflect.Int8:
		datum = Datum{uint8(value.Int())}
	case reflect.Int16:
		binary.BigEndian.PutUint16(datum, uint16(value.Int()))
	case reflect.Int, reflect.Int32:
		binary.BigEndian.PutUint32(datum, uint32(value.Int()))
	case reflect.Int64:
		binary.BigEndian.PutUint64(datum, uint64(value.Int()))
	case reflect.Uint16:
		binary.BigEndian.PutUint16(datum, uint16(value.Uint()))
	case reflect.Uint, reflect.Uint32:
		binary.BigEndian.PutUint32(datum, uint32(value.Uint()))
	case reflect.Uint64:
		binary.BigEndian.PutUint64(datum, value.Uint())
	default:
		panic("not implemented")
	}
	return datum
}
