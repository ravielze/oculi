package context

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestContext_ProcessV2_OnErrorFunction(t *testing.T) {
	var (
		anyErr  = errors.New("any error")
		singleF = func(i int, x []string) error {
			if i%2 == 0 {
				return nil
			}
			return anyErr
		}
	)
	t.Run("single return with error, onError does run", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)

		outer := "something"
		result := oculiCtx.Process(
			NewFunction(singleF, 1, []string{"a", "b", "c", "d", "e"}),
			func() {
				outer = "baba"
			},
			nil,
		)

		assert.Equal(t, oculiCtx.httpCode, 500)
		assert.Len(t, oculiCtx.errors, 1)
		assert.Equal(t, nil, result)
		assert.Equal(t, "baba", outer)
	})
	t.Run("single return with error, onError does not run", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)

		outer := "something"
		result := oculiCtx.Process(
			NewFunction(singleF, 2, []string{"a", "b", "c", "d", "e"}),
			func() {
				outer = "baba"
			},
			nil,
		)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, nil, result)
		assert.Equal(t, "something", outer)
	})
}

func TestContext_ProcessV2_WithParam(t *testing.T) {
	var (
		anyErr  = errors.New("any error")
		zeroF   = func(i int, x []string) {}
		singleF = func(i int, x []string) string {
			return x[i]
		}
		singleFerror = func(i int, x []string) error {
			if i%2 == 0 {
				return nil
			}
			return anyErr
		}
		doubleF = func(i int, x []string) (string, error) {
			if i < 0 {
				return "", anyErr
			}
			return x[i], nil
		}
		tripleF = func(i int, x []string) (string, string, error) {
			if i < 0 {
				return "", "", anyErr
			}
			return x[i], x[i-1], nil
		}
	)
	t.Run("no return", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)

		result := oculiCtx.Process(
			NewFunction(zeroF, 1, []string{"a", "b", "c", "d", "e"}),
			nil,
			nil,
		)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, nil, result)
	})

	t.Run("single return with no error, return string", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)

		result := oculiCtx.Process(
			NewFunction(singleF, 2, []string{"a", "b", "c", "d", "e"}),
			nil,
			nil,
		)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, "c", result)
	})

	t.Run("single return with error, return no error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)

		result := oculiCtx.Process(
			NewFunction(singleFerror, 2, []string{"a", "b", "c", "d", "e"}),
			nil,
			nil,
		)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, nil, result)
	})

	t.Run("single return with error, return error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		result := oculiCtx.Process(
			NewFunction(singleFerror, 1, []string{"a", "b", "c", "d", "e"}),
			nil,
			nil,
		)

		assert.Equal(t, oculiCtx.httpCode, 500)
		assert.Len(t, oculiCtx.errors, 1)
		assert.Equal(t, nil, result)
	})

	t.Run("double return with error, return string and no error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)

		result := oculiCtx.Process(
			NewFunction(doubleF, 3, []string{"a", "b", "c", "d", "e"}),
			nil,
			nil,
		)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, "d", result)
	})

	t.Run("double return with error, return 2 string and error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		result := oculiCtx.Process(
			NewFunction(doubleF, -1, []string{"a", "b", "c", "d", "e"}),
			nil,
			nil,
		)

		assert.Equal(t, oculiCtx.httpCode, 500)
		assert.Len(t, oculiCtx.errors, 1)
		assert.Equal(t, nil, result)
	})

	t.Run("triple return with error, return 2 string and no error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)

		result := oculiCtx.Process(
			NewFunction(tripleF, 3, []string{"a", "b", "c", "d", "e"}),
			nil,
			nil,
		)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, "d", result)
	})

	t.Run("triple return with error, return string and error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		result := oculiCtx.Process(
			NewFunction(tripleF, -1, []string{"a", "b", "c", "d", "e"}),
			nil,
			nil,
		)

		assert.Equal(t, oculiCtx.httpCode, 500)
		assert.Len(t, oculiCtx.errors, 1)
		assert.Equal(t, nil, result)
	})
}

func TestContext_ProcessV2_NoParam(t *testing.T) {
	var (
		anyErr = errors.New("any error")
	)
	t.Run("no return", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() {}

		result := oculiCtx.Process(NewFunction(fn), nil, nil)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, nil, result)
	})

	t.Run("single return with error, return no error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() error {
			return nil
		}

		result := oculiCtx.Process(NewFunction(fn), nil, nil)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, nil, result)
	})

	t.Run("single return with error, return error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() error {
			return anyErr
		}

		result := oculiCtx.Process(NewFunction(fn), nil, nil)

		assert.Equal(t, oculiCtx.httpCode, 500)
		assert.Len(t, oculiCtx.errors, 1)
		assert.Equal(t, nil, result)
	})

	t.Run("single return with no error, return int", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() int {
			return 10
		}

		result := oculiCtx.Process(NewFunction(fn), nil, nil)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, 10, result)
	})

	t.Run("single return with no error, return string", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() string {
			return "abc"
		}

		result := oculiCtx.Process(NewFunction(fn), nil, nil)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, "abc", result)
	})

	t.Run("single return with no error, return array uint", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() []uint {
			return []uint{0, 2, 4}
		}

		result := oculiCtx.Process(NewFunction(fn), nil, nil)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, []uint{0, 2, 4}, result)
	})

	t.Run("one return and error, return string with error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (string, error) {
			return "abc", anyErr
		}

		result := oculiCtx.Process(NewFunction(fn), nil, nil)

		assert.Equal(t, oculiCtx.httpCode, 500)
		assert.Len(t, oculiCtx.errors, 1)
		assert.Equal(t, nil, result)
	})

	t.Run("one return and error, return int with error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (int, error) {
			return 5, anyErr
		}

		result := oculiCtx.Process(NewFunction(fn), nil, nil)

		assert.Equal(t, oculiCtx.httpCode, 500)
		assert.Len(t, oculiCtx.errors, 1)
		assert.Equal(t, nil, result)
	})

	t.Run("one return with no error, return int and no error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (int, error) {
			return 5, nil
		}

		result := oculiCtx.Process(NewFunction(fn), nil, nil)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, 5, result)
	})

	t.Run("one return with no error, return float and no error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (float64, error) {
			return 18273.22, nil
		}

		result := oculiCtx.Process(NewFunction(fn), nil, nil)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, 18273.22, result)
	})

	t.Run("one return with no error, return array of float and no error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() ([]float64, error) {
			return []float64{1, 2.19, 3, 4.5}, nil
		}

		result := oculiCtx.Process(NewFunction(fn), nil, nil)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, []float64{1, 2.19, 3, 4.5}, result)
	})

	t.Run("more than two return with error, return float, string, and error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (float64, string, error) {
			return 18273.22, "abc", anyErr
		}

		result := oculiCtx.Process(NewFunction(fn), nil, nil)

		assert.Equal(t, oculiCtx.httpCode, 500)
		assert.Len(t, oculiCtx.errors, 1)
		assert.Equal(t, nil, result)
	})

	t.Run("more than two return with error, return float, array of string, and error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (float64, []string, error) {
			return 18273.22, []string{"abc", "def"}, anyErr
		}

		result := oculiCtx.Process(NewFunction(fn), nil, nil)

		assert.Equal(t, oculiCtx.httpCode, 500)
		assert.Len(t, oculiCtx.errors, 1)
		assert.Equal(t, nil, result)
	})

	t.Run("more than two return with no error, return float, array of string and no error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (float64, []string, error) {
			return 18273.22, []string{"abc", "def"}, nil
		}

		result := oculiCtx.Process(NewFunction(fn), nil, nil)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, 18273.22, result)
	})

	t.Run("more than two return with no error, return float, string, and no error", func(t *testing.T) {
		ctx := echo.New().NewContext(&http.Request{}, httptest.NewRecorder())
		oculiCtx := New(ctx)
		fn := func() (float64, string, error) {
			return 18273.22, "abc", nil
		}

		result := oculiCtx.Process(NewFunction(fn), nil, nil)

		assert.Equal(t, oculiCtx.httpCode, 200)
		assert.Len(t, oculiCtx.errors, 0)
		assert.Equal(t, 18273.22, result)
	})
}
