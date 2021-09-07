package webserver

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

func InfoRoutes(ec *echo.Echo) func() {
	return func() {
		routes := make(map[string][]string)
		for _, route := range ec.Routes() {
			routes[route.Method] = append(routes[route.Method], route.Path)
		}

		methodList := []string{http.MethodGet, http.MethodPost, http.MethodPatch,
			http.MethodPut, http.MethodDelete, http.MethodConnect, http.MethodHead,
			http.MethodOptions, http.MethodTrace}

		for method := range routes {
			sort.Strings(routes[method])
		}

		fmt.Printf("[%s Oculi %s] %v | Showing registered routes:\n",
			magenta, reset, time.Now().Format("15:04:05 02 Jan 2006"))

		lastRouteMethod := ""
		for _, method := range methodList {
			for _, route := range routes[method] {
				if lastRouteMethod != method {
					fmt.Printf("[%s Oculi %s] %-20v | %s %-7s %s %#v\n",
						magenta, reset, time.Now().Format("15:04:05 02 Jan 2006"),
						methodColor(method), method, reset,
						route,
					)
				} else {
					fmt.Printf("[%s Oculi %s] %-20v | %s %-7s %s %#v\n",
						magenta, reset, "",
						methodColor(method), "", reset,
						route,
					)
				}
				lastRouteMethod = method
			}
		}
	}
}

func formatRequest(ec echo.Context, start time.Time) string {
	now := time.Now()
	statusCode := ec.Response().Status
	method := ec.Request().Method
	latency := now.Sub(start)
	path := ec.Path()
	if ec.QueryString() != "" {
		path = path + "?" + ec.QueryString()
	}
	return fmt.Sprintf("[%s Oculi %s] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v",
		magenta, reset,
		now.Format("15:04:05 02 Jan 2006"),
		statusCodeColor(statusCode), statusCode, reset,
		latency,
		ec.RealIP(),
		methodColor(method), method, reset,
		path,
	)
}

func statusCodeColor(code int) string {
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red
	}
}

func methodColor(method string) string {

	switch method {
	case http.MethodGet:
		return blue
	case http.MethodPost:
		return cyan
	case http.MethodPut:
		return yellow
	case http.MethodDelete:
		return red
	case http.MethodPatch:
		return green
	case http.MethodHead:
		return magenta
	case http.MethodOptions:
		return white
	default:
		return reset
	}
}
