package auth

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStandardCredentials_MapMetadata(t *testing.T) {
	t.Run("when metadata is nil", func(t *testing.T) {
		cred := StandardCredentials{ID: 1}
		m, err := cred.MapMetadata()

		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{}, m)
		assert.Len(t, m, 0)
	})
	t.Run("when metadata is int", func(t *testing.T) {
		cred := StandardCredentials{ID: 1, Metadata: 1}
		m, err := cred.MapMetadata()

		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"metadata": 1}, m)
		assert.Len(t, m, 1)
	})
	t.Run("when metadata is array of int", func(t *testing.T) {
		cred := StandardCredentials{ID: 1, Metadata: []int{1, 2, 3}}
		m, err := cred.MapMetadata()

		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"metadata": []int{1, 2, 3}}, m)
		assert.Len(t, m, 1)
	})
	t.Run("when metadata is string", func(t *testing.T) {
		cred := StandardCredentials{ID: 1, Metadata: "123"}
		m, err := cred.MapMetadata()

		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"metadata": "123"}, m)
		assert.Len(t, m, 1)
	})
	t.Run("when metadata is pointer", func(t *testing.T) {
		x := 55.4
		cred := StandardCredentials{ID: 1, Metadata: &x}
		m, err := cred.MapMetadata()

		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"metadata": &x}, m)
		assert.Len(t, m, 1)
	})
	t.Run("when metadata is error", func(t *testing.T) {
		cred := StandardCredentials{ID: 1, Metadata: errors.New("123")}
		m, err := cred.MapMetadata()

		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"metadata": errors.New("123")}, m)
		assert.Len(t, m, 1)
	})
	t.Run("when metadata is struct", func(t *testing.T) {
		type A struct {
			babiHutan    int
			CacingAlaska string
			Dadu         error
			EeeeeE       []float32
		}
		cred := StandardCredentials{ID: 1, Metadata: A{babiHutan: 1, CacingAlaska: "asd", Dadu: nil, EeeeeE: []float32{1.1, 3.2}}}
		m, err := cred.MapMetadata()

		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"cacing_alaska": "asd", "dadu": nil, "eeeee_e": []float32{1.1, 3.2}}, m)
		assert.Len(t, m, 3)
	})
	t.Run("when metadata is map[string]interface", func(t *testing.T) {
		cred := StandardCredentials{
			ID:       1,
			Metadata: map[string]interface{}{"cacing_alaska": "asd", "dadu": nil, "eeeee_e": []float32{1.1, 3.2}},
		}
		m, err := cred.MapMetadata()

		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"cacing_alaska": "asd", "dadu": nil, "eeeee_e": []interface{}{1.1, 3.2}}, m)
		assert.Len(t, m, 3)
	})
	t.Run("when metadata is map[int]float32", func(t *testing.T) {
		cred := StandardCredentials{
			ID:       1,
			Metadata: map[int]float32{1: 0.5, 2: 2.5},
		}
		m, err := cred.MapMetadata()

		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"1": 0.5, "2": 2.5}, m)
		assert.Len(t, m, 2)
	})
	t.Run("when metadata is map[int][]int", func(t *testing.T) {
		cred := StandardCredentials{
			ID:       1,
			Metadata: map[int][]int{1: {1, 2, 3}, 2: {4, 5, 6}},
		}
		m, err := cred.MapMetadata()

		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"1": []interface{}{1.0, 2.0, 3.0}, "2": []interface{}{4.0, 5.0, 6.0}}, m)
		assert.Len(t, m, 2)
	})
	t.Run("when metadata is map[struct]int", func(t *testing.T) {
		type Key struct {
			A string
			B int
		}
		cred := StandardCredentials{
			ID:       1,
			Metadata: map[Key]int{{A: "abc", B: 1}: 1, {A: "def", B: 2}: 3},
		}
		m, err := cred.MapMetadata()

		assert.Error(t, err)
		assert.Equal(t, map[string]interface{}{}, m)
		assert.Len(t, m, 0)
	})
}
