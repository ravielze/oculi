package user

import (
	"github.com/ravielze/oculi/example/constants"
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/example/model/dto/user"
	"github.com/ravielze/oculi/request"
)

func (h *handler) Check(req request.ReqContext) (user.UserResponse, error) {
	credentialsData := req.Identifier()
	if credentialsData.ID == 0 {
		return user.UserResponse{}, constants.ErrNotLoggedIn
	}
	userDataBuff, _ := h.resource.Json.Marshal(credentialsData.Metadata)
	var userData dao.User
	h.resource.Json.Unmarshal(userDataBuff, &userData)
	return user.NewUserResponse(userData), nil
}
