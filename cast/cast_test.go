package cast

import (
	"testing"
	"time"

	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
)

func TestCast_Fallback(t *testing.T) {
	t.Run("ToInt64F", func(t *testing.T) {
		assert.Equal(t, ToInt64F("1", 2), int64(1))
		assert.Equal(t, ToInt64F("1s", 2), int64(2))
	})

	t.Run("ToValueF", func(t *testing.T) {
		assert.Equal(t, ToValueF(cast.ToStringSliceE, []int{1, 2, 3}, []string{"4", "5", "6"}), []string{"1", "2", "3"})
		assert.Equal(t, ToValueF(cast.ToIntSliceE, []string{"4s", "5s", "6s"}, []int{1, 2, 3}), []int{1, 2, 3})
	})

	t.Run("ToXxxxF", func(t *testing.T) {
		assert.True(t, ToBoolF("true", false))
		assert.False(t, ToBoolF("false", true))

		assert.False(t, ToBoolF(0, true))
		assert.True(t, ToBoolF(1, false))

		assert.True(t, ToBoolF("kkk", true))

		assert.Equal(t, ToDurationF("1s", 2), 1*time.Second)
	})
}
