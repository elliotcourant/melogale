package engine

import (
	"reflect"
)

type Type interface {
	Kind() reflect.Kind
	Size() uint8
}
