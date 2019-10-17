package base

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestState(t *testing.T) {
	t.Run("none", func(t *testing.T) {
		s := None
		assert.False(t, s.IsReadable())
		assert.False(t, s.IsWritable())
		assert.False(t, s.IsDeletable())
	})

	t.Run("readable", func(t *testing.T) {
		s := Readable
		assert.True(t, s.IsReadable())
		assert.False(t, s.IsWritable())
		assert.False(t, s.IsDeletable())
	})

	t.Run("readable-writable", func(t *testing.T) {
		s := Readable | Writable
		assert.True(t, s.IsReadable())
		assert.True(t, s.IsWritable())
		assert.False(t, s.IsDeletable())
	})

	t.Run("deletable", func(t *testing.T) {
		s := Deletable
		assert.False(t, s.IsReadable())
		assert.False(t, s.IsWritable())
		assert.True(t, s.IsDeletable())
	})
}
