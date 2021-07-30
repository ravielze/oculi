package errors

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDetailedErrors(t *testing.T) {
	t.Run("with details", func(t *testing.T) {
		err := NewDetailedErrors("test", 1, "abc", 50.2)

		assert.Equal(t, "test", err.Error())
		assert.Len(t, err.(DetailedErrors).Details, 3)
	})
	t.Run("with no details", func(t *testing.T) {
		err := NewDetailedErrors("test")

		assert.Equal(t, "test", err.Error())
		assert.Len(t, err.(DetailedErrors).Details, 0)
	})
}

func TestInjectDetails(t *testing.T) {
	t.Run("fmt Error with details", func(t *testing.T) {
		raw := fmt.Errorf("test %d", 1)
		err := InjectDetails(raw, 5, "abc", 10.15)

		assert.Equal(t, "test 1", err.Error())
		assert.Equal(t, []interface{}{5, "abc", 10.15}, err.(DetailedErrors).Details)
		assert.Len(t, err.(DetailedErrors).Details, 3)

		errType := reflect.TypeOf(err).Name()
		assert.Equal(t, reflect.TypeOf(DetailedErrors{}).Name(), errType)
	})
	t.Run("fmt Error with no details", func(t *testing.T) {
		raw := fmt.Errorf("test %d", 3)
		err := InjectDetails(raw)

		assert.Equal(t, "test 3", err.Error())
		assert.Equal(t, []interface{}(nil), err.(DetailedErrors).Details)
		assert.Len(t, err.(DetailedErrors).Details, 0)

		errType := reflect.TypeOf(err).Name()
		assert.Equal(t, reflect.TypeOf(DetailedErrors{}).Name(), errType)
	})
	t.Run("errors new with details", func(t *testing.T) {
		raw := errors.New("gatau ah males, mau beli truck")
		err := InjectDetails(raw, 5, "abc", 10.15)

		assert.Equal(t, "gatau ah males, mau beli truck", err.Error())
		assert.Equal(t, []interface{}{5, "abc", 10.15}, err.(DetailedErrors).Details)
		assert.Len(t, err.(DetailedErrors).Details, 3)

		errType := reflect.TypeOf(err).Name()
		assert.Equal(t, reflect.TypeOf(DetailedErrors{}).Name(), errType)
	})
	t.Run("errors new with no details", func(t *testing.T) {
		raw := errors.New("gatau ah males, mau beli truck")
		err := InjectDetails(raw)

		assert.Equal(t, "gatau ah males, mau beli truck", err.Error())
		assert.Equal(t, []interface{}(nil), err.(DetailedErrors).Details)
		assert.Len(t, err.(DetailedErrors).Details, 0)

		errType := reflect.TypeOf(err).Name()
		assert.Equal(t, reflect.TypeOf(DetailedErrors{}).Name(), errType)
	})
}

func TestDetails(t *testing.T) {
	t.Run("success convert, with details", func(t *testing.T) {
		err := NewDetailedErrors("au ah puyeng", 5, "abc", 10.15)
		details := Details(err)

		assert.Equal(t, []interface{}{5, "abc", 10.15}, details)
		assert.Len(t, err.(DetailedErrors).Details, 3)
	})
	t.Run("success convert, no details", func(t *testing.T) {
		err := NewDetailedErrors("au ah puyeng")
		details := Details(err)

		assert.Equal(t, []interface{}(nil), details)
		assert.Len(t, err.(DetailedErrors).Details, 0)
	})
	t.Run("failed convert", func(t *testing.T) {
		err := errors.New("blabla")
		details := Details(err)

		assert.Equal(t, []interface{}(nil), details)
	})
}
