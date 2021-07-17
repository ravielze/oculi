package webserver

import (
	"github.com/ravielze/oculi/server"
)

type (
	WebServer struct {
		infrastructure server.Infrastructure
		resource       server.Resource

		afterRun   []server.HookFunction
		beforeRun  []server.HookFunction
		beforeExit []server.HookFunction
		afterExit  []server.HookFunction
	}
)

func New(infrastructure server.Infrastructure, resource server.Resource) server.Server {
	return &WebServer{
		infrastructure: infrastructure,
		resource:       resource,
		afterRun:       make([]server.HookFunction, 0),
		beforeRun:      make([]server.HookFunction, 0),
		afterExit:      make([]server.HookFunction, 0),
		beforeExit:     make([]server.HookFunction, 0),
	}
}
