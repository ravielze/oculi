package health

import (
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/pkg/errors"
	"github.com/ravielze/oculi/common/model/dto/health"
	"github.com/ravielze/oculi/request"
)

func (h *handler) Check(ctx request.Context) health.CheckResponseDTO {
	var (
		response health.CheckResponseDTO
		m        runtime.MemStats
	)

	response.Ts = time.Now()
	response.Pid = os.Getpid()
	response.Uptime = time.Since(h.resource.Uptime()).String()
	response.Status = "ok"
	runtime.ReadMemStats(&m)
	response.Memory.Alloc = m.Alloc
	response.Memory.TotalAlloc = m.TotalAlloc
	response.Memory.Sys = m.Sys
	response.Memory.NumGC = m.NumGC
	response.Memory.HeapAlloc = m.HeapAlloc
	response.Memory.HeapSys = m.HeapSys

	// NOTE: Add connection checking per service
	if err := h.resource.Database.Ping(ctx.GetContext()); err != nil {
		ctx.AddError(http.StatusBadGateway, errors.Wrap(err, "db connection: "))
	}
	return response
}
