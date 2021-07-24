package user

import (
	"encoding/json"

	userCommonDto "github.com/ravielze/oculi/common/model/dto/user"
	consts "github.com/ravielze/oculi/constant/key"
	"github.com/ravielze/oculi/example/constants"
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/example/model/dto/user"
	"github.com/ravielze/oculi/request"
)

func (h *handler) Check(req request.EchoContext) (user.UserResponse, error) {
	k := req.Echo().Get(consts.KeyCredentials)
	if k == nil {
		return user.UserResponse{}, constants.ErrNotLoggedIn
	}
	credentialsData := k.(userCommonDto.CredentialsDTO)
	userDataBuff, _ := json.Marshal(credentialsData.Metadata)
	var userData dao.User
	json.Unmarshal(userDataBuff, &userData)
	return user.NewUserResponse(userData), nil
}
