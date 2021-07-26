package radix36

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("OK (1)", func(t *testing.T) {
		result := Validate("ABCDEF")
		assert.True(t, result)
	})
	t.Run("OK (2)", func(t *testing.T) {
		result := Validate("ABCDEF12345")
		assert.True(t, result)
	})
	t.Run("OK (3)", func(t *testing.T) {
		result := Validate("12345")
		assert.True(t, result)
	})
	t.Run("OK (4)", func(t *testing.T) {
		result := Validate("PQRSTU")
		assert.True(t, result)
	})
	t.Run("OK (5)", func(t *testing.T) {
		result := Validate("12345PQRSTU")
		assert.True(t, result)
	})
	t.Run("OK (6)", func(t *testing.T) {
		result := Validate("PQRSTUmnop123")
		assert.True(t, result)
	})
	t.Run("OK (7)", func(t *testing.T) {
		result := Validate("mnop12345PQRSTU")
		assert.True(t, result)
	})
	t.Run("NOT OK (1)", func(t *testing.T) {
		result := Validate("ABCDEF!")
		assert.False(t, result)
	})
	t.Run("NOT OK (2)", func(t *testing.T) {
		result := Validate("ABCDEF,123")
		assert.False(t, result)
	})
	t.Run("NOT OK (3)", func(t *testing.T) {
		result := Validate("abcdef@123")
		assert.False(t, result)
	})
	t.Run("NOT OK (4)", func(t *testing.T) {
		result := Validate("abcdef@123😊")
		assert.False(t, result)
	})
}
