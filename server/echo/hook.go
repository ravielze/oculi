package webserver

import "github.com/ravielze/oculi/server"

func (w *WebServer) BeforeRun(hf server.HookFunction) server.Server {
	w.beforeRun = append(w.beforeRun, hf)
	return w
}

func (w *WebServer) AfterRun(hf server.HookFunction) server.Server {
	w.afterRun = append(w.afterRun, hf)
	return w
}

func (w *WebServer) BeforeExit(hf server.HookFunction) server.Server {
	w.beforeExit = append(w.beforeExit, hf)
	return w
}

func (w *WebServer) AfterExit(hf server.HookFunction) server.Server {
	w.afterExit = append(w.afterExit, hf)
	return w
}

func (w *WebServer) apply(hooks []server.HookFunction) error {
	for i := range hooks {
		if err := hooks[i](w.resource); err != nil {
			return err
		}
	}
	return nil
}
