package docs

import (
	_ "embed"
	"net/http"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/encoding/jsoniter"
)

//go:embed template.html
var docsRaw string

//go:embed example.json
var Data string
var ProcessedData map[string]interface{} = nil

var l sync.Mutex

var json = jsoniter.New()

func SetData(data string) {
	l.Lock()
	defer l.Unlock()
	Data = data
	if err := json.Unmarshal([]byte(Data), &ProcessedData); err != nil {
		panic(err.Error())
	}
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

func New(e *echo.Echo, serviceName, baselink string) Documentation {
	d := &documentation{}
	d.docsRendered = strings.ReplaceAll(docsRaw, "{{ serviceName }}", serviceName)
	d.docsRendered = strings.ReplaceAll(d.docsRendered, "{{ source }}", baselink+"/docs/data")

	e.GET("/docs", d.Get)
	e.GET("/docs/data", d.Data)
	return d
}

func (d *documentation) Get(ec echo.Context) error {
	return ec.HTMLBlob(http.StatusOK, []byte(d.docsRendered))
}

func (d *documentation) Data(ec echo.Context) error {
	if ProcessedData == nil {
		SetData(Data)
	}
	return ec.JSONPretty(http.StatusOK, ProcessedData, " ")
}
