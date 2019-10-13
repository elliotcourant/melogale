package sql

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewUniqueColumn(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		id := NewUniqueColumn(1, 3)
		assert.Equal(t, id.TableId(), uint8(1))
		assert.Equal(t, id.ColumnId(), uint8(3))
	})
}
