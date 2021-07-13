package context

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

// Request returns `*http.Request`.
func (ctx *Context) Request() *http.Request {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.Request()
}

// SetRequest sets `*http.Request`.
func (ctx *Context) SetRequest(r *http.Request) {
	if ctx.ec == nil {
		return
	}
	ctx.ec.SetRequest(r)
}

// SetResponse sets `*Response`.
func (ctx *Context) SetResponse(r *echo.Response) {
	if ctx.ec == nil {
		return
	}
	ctx.ec.SetResponse(r)
}

// Response returns `*Response`.
func (ctx *Context) Response() *echo.Response {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.Response()
}

// IsTLS returns true if HTTP connection is TLS otherwise false.
func (ctx *Context) IsTLS() bool {
	if ctx.ec == nil {
		return false
	}
	return ctx.ec.IsTLS()
}

// IsWebSocket returns true if HTTP connection is WebSocket otherwise false.
func (ctx *Context) IsWebSocket() bool {
	if ctx.ec == nil {
		return false
	}
	return ctx.ec.IsWebSocket()
}

// Scheme returns the HTTP protocol scheme, `http` or `https`.
func (ctx *Context) Scheme() string {
	if ctx.ec == nil {
		return ""
	}
	return ctx.ec.Scheme()
}

// RealIP returns the client's network address based on `X-Forwarded-For`
// or `X-Real-IP` request header.
// The behavior can be configured using `Echo#IPExtractor`.
func (ctx *Context) RealIP() string {
	if ctx.ec == nil {
		return ""
	}
	return ctx.ec.RealIP()
}

// Path returns the registered path for the handler.
func (ctx *Context) Path() string {
	if ctx.ec == nil {
		return ""
	}
	return ctx.ec.Path()
}

// SetPath sets the registered path for the handler.
func (ctx *Context) SetPath(p string) {
	if ctx.ec == nil {
		return
	}
	ctx.ec.SetPath(p)
}

// Param returns path parameter by name.
func (ctx *Context) Param(name string) string {
	if ctx.ec == nil {
		return ""
	}
	return ctx.ec.Param(name)
}

// ParamNames returns path parameter names.
func (ctx *Context) ParamNames() []string {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.ParamNames()
}

// SetParamNames sets path parameter names.
func (ctx *Context) SetParamNames(names ...string) {
	if ctx.ec == nil {
		return
	}
	ctx.ec.SetParamNames(names...)
}

// ParamValues returns path parameter values.
func (ctx *Context) ParamValues() []string {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.ParamValues()
}

// SetParamValues sets path parameter values.
func (ctx *Context) SetParamValues(values ...string) {
	if ctx.ec == nil {
		return
	}
	ctx.ec.SetParamValues(values...)
}

// QueryParam returns the query param for the provided name.
func (ctx *Context) QueryParam(name string) string {
	if ctx.ec == nil {
		return ""
	}
	return ctx.ec.QueryParam(name)
}

// QueryParams returns the query parameters as `url.Values`.
func (ctx *Context) QueryParams() url.Values {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.QueryParams()
}

// QueryString returns the URL query string.
func (ctx *Context) QueryString() string {
	if ctx.ec == nil {
		return ""
	}
	return ctx.ec.QueryString()
}

// FormValue returns the form field value for the provided name.
func (ctx *Context) FormValue(name string) string {
	if ctx.ec == nil {
		return ""
	}
	return ctx.ec.FormValue(name)
}

// FormParams returns the form parameters as `url.Values`.
func (ctx *Context) FormParams() (url.Values, error) {
	if ctx.ec == nil {
		return nil, nil
	}
	return ctx.ec.FormParams()
}

// FormFile returns the multipart form file for the provided name.
func (ctx *Context) FormFile(name string) (*multipart.FileHeader, error) {
	if ctx.ec == nil {
		return nil, nil
	}
	return ctx.ec.FormFile(name)
}

// MultipartForm returns the multipart form.
func (ctx *Context) MultipartForm() (*multipart.Form, error) {
	if ctx.ec == nil {
		return nil, nil
	}
	return ctx.ec.MultipartForm()
}

// Cookie returns the named cookie provided in the request.
func (ctx *Context) Cookie(name string) (*http.Cookie, error) {
	if ctx.ec == nil {
		return nil, nil
	}
	return ctx.ec.Cookie(name)
}

// SetCookie adds a `Set-Cookie` header in HTTP response.
func (ctx *Context) SetCookie(cookie *http.Cookie) {
	if ctx.ec == nil {
		return
	}
	ctx.ec.SetCookie(cookie)
}

// Cookies returns the HTTP cookies sent with the request.
func (ctx *Context) Cookies() []*http.Cookie {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.Cookies()
}

// Get retrieves data from the context.
func (ctx *Context) Get(key string) interface{} {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.Get(key)
}

// Set saves data in the context.
func (ctx *Context) Set(key string, val interface{}) {
	if ctx.ec == nil {
		return
	}
	ctx.ec.Set(key, val)
}

// Bind binds the request body into provided type `i`. The default binder
// does it based on Content-Type header.
func (ctx *Context) Bind(i interface{}) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.Bind(i)
}

// Validate validates provided `i`. It is usually called after `Context#Bind()`.
// Validator must be registered using `Echo#Validator`.
func (ctx *Context) Validate(i interface{}) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.Validate(i)
}

// Render renders a template with data and sends a text/html response with status
// code. Renderer must be registered using `Echo.Renderer`.
func (ctx *Context) Render(code int, name string, data interface{}) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.Render(code, name, data)
}

// HTML sends an HTTP response with status code.
func (ctx *Context) HTML(code int, html string) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.HTML(code, html)
}

// HTMLBlob sends an HTTP blob response with status code.
func (ctx *Context) HTMLBlob(code int, b []byte) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.HTMLBlob(code, b)
}

// String sends a string response with status code.
func (ctx *Context) String(code int, s string) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.String(code, s)
}

// JSON sends a JSON response with status code.
func (ctx *Context) JSON(code int, i interface{}) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.JSON(code, i)
}

// JSONPretty sends a pretty-print JSON with status code.
func (ctx *Context) JSONPretty(code int, i interface{}, indent string) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.JSONPretty(code, i, indent)
}

// JSONBlob sends a JSON blob response with status code.
func (ctx *Context) JSONBlob(code int, b []byte) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.JSONBlob(code, b)
}

// JSONP sends a JSONP response with status code. It uses `callback` to construct
// the JSONP payload.
func (ctx *Context) JSONP(code int, callback string, i interface{}) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.JSONP(code, callback, i)
}

// JSONPBlob sends a JSONP blob response with status code. It uses `callback`
// to construct the JSONP payload.
func (ctx *Context) JSONPBlob(code int, callback string, b []byte) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.JSONPBlob(code, callback, b)
}

// XML sends an XML response with status code.
func (ctx *Context) XML(code int, i interface{}) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.XML(code, i)
}

// XMLPretty sends a pretty-print XML with status code.
func (ctx *Context) XMLPretty(code int, i interface{}, indent string) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.XMLPretty(code, i, indent)
}

// XMLBlob sends an XML blob response with status code.
func (ctx *Context) XMLBlob(code int, b []byte) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.XMLBlob(code, b)
}

// Blob sends a blob response with status code and content type.
func (ctx *Context) Blob(code int, contentType string, b []byte) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.Blob(code, contentType, b)
}

// Stream sends a streaming response with status code and content type.
func (ctx *Context) Stream(code int, contentType string, r io.Reader) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.Stream(code, contentType, r)
}

// File sends a response with the content of the file.
func (ctx *Context) File(file string) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.File(file)
}

// Attachment sends a response as attachment, prompting client to save the
// file.
func (ctx *Context) Attachment(file string, name string) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.Attachment(file, name)
}

// Inline sends a response as inline, opening the file in the browser.
func (ctx *Context) Inline(file string, name string) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.Inline(file, name)
}

// NoContent sends a response with no body and a status code.
func (ctx *Context) NoContent(code int) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.NoContent(code)
}

// Redirect redirects the request to a provided URL with status code.
func (ctx *Context) Redirect(code int, url string) error {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.Redirect(code, url)
}

// Error invokes the registered HTTP error handler. Generally used by middleware.
func (ctx *Context) Error(err error) {
	if ctx.ec == nil {
		return
	}
	ctx.ec.Error(err)
}

// Handler returns the matched handler by router.
func (ctx *Context) Handler() echo.HandlerFunc {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.Handler()
}

// SetHandler sets the matched handler by router.
func (ctx *Context) SetHandler(h echo.HandlerFunc) {
	if ctx.ec == nil {
		return
	}
	ctx.ec.SetHandler(h)
}

// Logger returns the `Logger` instance.
func (ctx *Context) Logger() echo.Logger {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.Logger()
}

// Set the logger
func (ctx *Context) SetLogger(l echo.Logger) {
	if ctx.ec == nil {
		return
	}
	ctx.ec.SetLogger(l)
}

// Echo returns the `Echo` instance.
func (ctx *Context) Echo() *echo.Echo {
	if ctx.ec == nil {
		return nil
	}
	return ctx.ec.Echo()
}

// Reset resets the context after request completes. It must be called along
// with `Echo#AcquireContext()` and `Echo#ReleaseContext()`.
// See `Echo#ServeHTTP()`
func (ctx *Context) Reset(r *http.Request, w http.ResponseWriter) {
	if ctx.ec == nil {
		return
	}
	ctx.ec.Reset(r, w)
}
