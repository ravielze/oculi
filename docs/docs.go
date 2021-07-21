package docs

import (
	_ "embed"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
)

//go:embed template.html
var docsRaw string

//go:embed example.json
var Data string

var l sync.Mutex

func SetData(data string) {
	l.Lock()
	defer l.Unlock()
	Data = data
}

type (
	documentation struct {
		docsRendered string
	}

	Documentation interface {
		Get(echo.Context) error
		Data(echo.Context) error
	}
)

func New(e *echo.Echo, serviceName, host string, port int) Documentation {
	d := &documentation{}
	d.docsRendered = strings.ReplaceAll(docsRaw, "{{ serviceName }}", serviceName)
	address := host
	if port != 80 {
		address = host + ":" + strconv.Itoa(port)
	}
	d.docsRendered = strings.ReplaceAll(d.docsRendered, "{{ source }}", address+"/docs/data")

	e.GET("/docs", d.Get)
	e.GET("/docs/data", d.Data)
	return d
}

func (d *documentation) Get(ec echo.Context) error {
	return ec.HTMLBlob(http.StatusOK, []byte(d.docsRendered))
}

func (d *documentation) Data(ec echo.Context) error {
	return ec.String(http.StatusOK, Data)
}
