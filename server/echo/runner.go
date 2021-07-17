package webserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	oculiContext "github.com/ravielze/oculi/context"
)

type ServiceInfo struct {
	Name       string `json:"service_name"`
	Identifier string `json:"identifier"`
}

func (w *WebServer) Run() error {
	if err := w.start(); err != nil {
		return err
	}

	sig := make(chan os.Signal, 3)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	if err := w.apply(w.beforeRun); err != nil {
		return err
	}

	go w.serve(sig)

	if err := w.apply(w.afterRun); err != nil {
		return err
	}
	<-sig

	if err := w.apply(w.beforeExit); err != nil {
		return err
	}
	w.stop()
	if err := w.apply(w.afterExit); err != nil {
		return err
	}
	return nil
}

func (w *WebServer) start() error {
	w.resource.Echo().Use(middleware.Recover())
	w.resource.Echo().Validator = w.resource.Validator()
	w.resource.Echo().Logger = w.resource.Logger()
	w.resource.Echo().Logger.SetLevel(log.INFO)

	w.resource.Echo().Use(middleware.Gzip())
	w.resource.Echo().Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	w.resource.Echo().Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, ok := c.(*oculiContext.Context)
			if !ok {
				ctx = oculiContext.New(c)
			}
			c = ctx
			return nil
		}
	})

	echo.NotFoundHandler = func(c echo.Context) error {
		ctx := oculiContext.New(c)
		ctx.AddError(http.StatusNotFound, errors.New("not found"))
		//TODO return response
		return nil
	}

	w.resource.Echo().GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, ServiceInfo{
			Name:       w.resource.ServiceName(),
			Identifier: w.resource.Identifier(),
		})
	})
	if err := w.infrastructure.Register(w.resource.Echo()); err != nil {
		w.resource.Logger().Error("error on register http")
		return err
	}

	w.resource.Echo().GET("/health", w.infrastructure.Health())
	return nil
}

func (w *WebServer) serve(sig chan os.Signal) {
	if err := w.resource.Echo().Start(fmt.Sprintf(":%d", w.resource.ServerPort())); err != nil {
		w.resource.Logger().Errorf("http server interrupted %s", err.Error())
		sig <- syscall.SIGINT
	} else {
		w.resource.Logger().Info("starting apps")
	}
}

func (w *WebServer) stop() {
	ctx, cancel := context.WithTimeout(context.Background(), w.resource.ServerGracefullyDuration())
	defer cancel()

	if err := w.resource.Echo().Shutdown(ctx); err != nil {
		w.resource.Logger().Error("failed to shutdown http server %s", err)
	}

	w.resource.Logger().Info("closing resource")
	if err := w.resource.Close(); err != nil {
		w.resource.Logger().Error("failed to close resource %s", err)
	}
}
