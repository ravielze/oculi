package webserver

import (
	"github.com/ravielze/oculi/server"
)

type (
	WebServer struct {
		infrastructure server.Infrastructure
		resource       server.Resource
		useDefaultCors bool
		useDefaultGZip bool

		afterRun   []server.HookFunction
		beforeRun  []server.HookFunction
		beforeExit []server.HookFunction
		afterExit  []server.HookFunction
	}

	useDefaultCors bool
	useDefaultGZip bool

	Option interface {
		Apply(w *WebServer)
	}
)

func DefaultCors(use bool) Option {
	return useDefaultCors(use)
}

func (o useDefaultCors) Apply(w *WebServer) {
	w.useDefaultCors = bool(o)
}

func DefaultGZip(use bool) Option {
	return useDefaultGZip(use)
}

func (o useDefaultGZip) Apply(w *WebServer) {
	w.useDefaultGZip = bool(o)
}

func New(infrastructure server.Infrastructure, resource server.Resource, options ...Option) server.Server {
	ws := &WebServer{
		useDefaultCors: true,
		useDefaultGZip: true,
		infrastructure: infrastructure,
		resource:       resource,
		afterRun:       make([]server.HookFunction, 0),
		beforeRun:      make([]server.HookFunction, 0),
		afterExit:      make([]server.HookFunction, 0),
		beforeExit:     make([]server.HookFunction, 0),
	}

	for _, o := range options {
		o.Apply(ws)
	}

	return ws
}
