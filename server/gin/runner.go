package webserver

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func (w *WebServer) Run() error {
	if err := w.infrastructure.Register(w.resource.Gin()); err != nil {
		return err
	}
	w.resource.Gin().GET("/health", w.infrastructure.Health())

	sig := make(chan os.Signal, 1)
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

	if err := w.exit(); err != nil {
		return err
	}

	if err := w.apply(w.afterExit); err != nil {
		return err
	}
	return nil
}

func (w *WebServer) exit() error {
	if err := w.resource.Close(); err != nil {
		return err
	}
	return nil
}

func (w *WebServer) serve(sig chan os.Signal) {
	if err := w.resource.Gin().Run(fmt.Sprintf(":%d", w.resource.ServerPort())); err != nil {
		sig <- syscall.SIGINT
	}
}
