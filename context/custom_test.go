package context

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestContext_Process(t *testing.T) {
	var (
		anyErr = errors.New("any error")
	)
	t.Run("single return with error (1)", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() error {
			return nil
		}

		oculiCtx.Process(fn())

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
	})
	t.Run("single return with error (2)", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() error {
			return anyErr
		}

		oculiCtx.Process(fn())

		assert.Equal(t, oculiCtx.httpCode, 400)
		assert.Len(t, oculiCtx.errors, 1)
	})

	t.Run("single return with no error (1)", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() int {
			return 10
		}

		oculiCtx.Process(fn())

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
	})

	t.Run("single return with no error (2)", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() string {
			return "abc"
		}

		oculiCtx.Process(fn())

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
	})

	t.Run("single return with no error (3)", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() []uint {
			return []uint{0, 2, 4}
		}

		oculiCtx.Process(fn())

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
	})

	t.Run("one return and error (1)", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (string, error) {
			return "abc", anyErr
		}

		oculiCtx.Process(fn())

		assert.Equal(t, oculiCtx.httpCode, 400)
		assert.Len(t, oculiCtx.errors, 1)
	})

	t.Run("one return and error (2)", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (int, error) {
			return 5, anyErr
		}

		oculiCtx.Process(fn())

		assert.Equal(t, oculiCtx.httpCode, 400)
		assert.Len(t, oculiCtx.errors, 1)
	})

	t.Run("one return with no error (1)", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (int, error) {
			return 5, nil
		}

		oculiCtx.Process(fn())

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
	})

	t.Run("one return with no error (2)", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (float64, error) {
			return 18273.22, nil
		}

		oculiCtx.Process(fn())

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
	})

	t.Run("one return with no error (3)", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() ([]float64, error) {
			return []float64{1, 2.19, 3, 4.5}, nil
		}

		oculiCtx.Process(fn())

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
	})

	t.Run("more than two return with error (1)", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (float64, string, error) {
			return 18273.22, "abc", anyErr
		}

		oculiCtx.Process(fn())

		assert.Equal(t, oculiCtx.httpCode, 400)
		assert.Len(t, oculiCtx.errors, 1)
	})

	t.Run("more than two return with error (2)", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (float64, []string, error) {
			return 18273.22, []string{"abc", "def"}, anyErr
		}

		oculiCtx.Process(fn())

		assert.Equal(t, oculiCtx.httpCode, 400)
		assert.Len(t, oculiCtx.errors, 1)
	})

	t.Run("more any with no error (1)", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (float64, []string, error) {
			return 18273.22, []string{"abc", "def"}, nil
		}

		oculiCtx.Process(fn())

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
	})

	t.Run("more than two return with error (2)", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (float64, string, error) {
			return 18273.22, "abc", nil
		}

		oculiCtx.Process(fn())

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
	})
}
