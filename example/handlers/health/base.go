package health

import (
	"github.com/ravielze/oculi/common/model/dto/health"
	"github.com/ravielze/oculi/example/resources"
	"github.com/ravielze/oculi/request"
)

type (
	Handler interface {
		Check(ctx request.ReqContext) health.CheckResponseDTO
		Reset(ctx request.ReqContext) error
	}

	handler struct {
		resource resources.Resource
	}
)

func NewHandler(resource resources.Resource) (Handler, error) {
	return &handler{
		resource: resource,
	}, nil
}
