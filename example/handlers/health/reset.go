package health

import (
	consts "github.com/ravielze/oculi/constant/key"
	"github.com/ravielze/oculi/example/constants"
	"github.com/ravielze/oculi/request"
)

func (h *handler) Reset(ctx request.ReqContext) error {
	if ctx.GetOrDefault(consts.QueryPrefix("key"), "").(string) != h.resource.Config.DatabaseResetKey {
		return constants.ErrResetUnauthorized
	}
	h.resource.DBManager.Reset()
	h.resource.DBManager.Install()
	return nil
}
