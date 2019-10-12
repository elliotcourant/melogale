package base

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewColumnFlag(t *testing.T) {
	t.Run("pkey-unique", func(t *testing.T) {
		flag := NewColumnFlag(ColumnPrimaryKey, ColumnUnique)
		assert.True(t, flag.IsPrimaryKey())
		assert.True(t, flag.IsUnique())
		assert.True(t, flag.IsIndexed())
	})

	t.Run("unique-indexed", func(t *testing.T) {
		flag := NewColumnFlag(ColumnUnique, ColumnIndexed)
		assert.False(t, flag.IsPrimaryKey())
		assert.True(t, flag.IsUnique())
		assert.True(t, flag.IsIndexed())
	})

	t.Run("primary key is unique", func(t *testing.T) {
		flag := NewColumnFlag(ColumnPrimaryKey)
		assert.True(t, flag.IsPrimaryKey())
		assert.True(t, flag.IsUnique())
		assert.True(t, flag.IsIndexed())
	})

	t.Run("unique is not primary key", func(t *testing.T) {
		flag := NewColumnFlag(ColumnUnique)
		assert.False(t, flag.IsPrimaryKey())
		assert.True(t, flag.IsUnique())
		assert.False(t, flag.IsIndexed())
	})
}
