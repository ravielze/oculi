package auth

import (
	"errors"
	"net/http"

	code "github.com/ravielze/oculi/common/code"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi/common/utils"
)

type Usecase struct {
	repo IRepo
}

func NewUsecase(repo IRepo) IUsecase {
	return Usecase{repo: repo}
}

func (uc Usecase) GetRawUser(userId uint) (User, error) {
	u, err := uc.repo.GetByID(userId)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (uc Usecase) GetByID(userId uint) (UserResponse, error) {
	u, err := uc.repo.GetByID(userId)
	if err != nil {
		return UserResponse{}, err
	}
	return u.Convert(), nil
}

func (uc Usecase) Login(item LoginRequest) (UserTokenResponse, error) {
	user, err := uc.repo.GetByEmail(item.Email)
	if err != nil {
		return UserTokenResponse{}, err
	}

	token, err := CreateToken(user.ID)
	if err != nil {
		return UserTokenResponse{}, err
	}

	if err := VerifyPassword(user.Password, item.Password); err != nil {
		return UserTokenResponse{}, errors.New("password not match")
	}
	return user.ConvertToken(token), nil
}

func (uc Usecase) Register(item RegisterRequest) (UserResponse, error) {
	user := item.Convert()
	role := Role(user.Role)
	if !role.IsExist() {
		return UserResponse{}, errors.New("role not exist")
	}
	if role.IsRestricted() {
		return UserResponse{}, errors.New("registering with that role is restricted")
	}
	if _, err := uc.repo.GetByEmail(user.Email); err == nil {
		return UserResponse{}, errors.New("account with that email is already exist")
	}
	u, errc := uc.repo.Create(user)
	if errc != nil {
		return UserResponse{}, errc
	}
	return u.Convert(), nil
}

func (uc Usecase) RegisterAdmin(item RegisterRequest) (UserResponse, error) {
	user := item.Convert()
	user.Role = int16(ROLE_ADMIN)
	if _, err := uc.repo.GetByEmail(user.Email); err == nil {
		return UserResponse{}, errors.New("account with that email is already exist")
	}
	u, errc := uc.repo.Create(user)
	if errc != nil {
		return UserResponse{}, errc
	}
	return u.Convert(), nil
}

func (uc Usecase) Update(user User, item UpdateRequest) error {
	return uc.repo.Update(item.Convert(user))
}

func (u Usecase) GetUser(ctx *gin.Context) User {
	if user, ok := ctx.Keys["user"].(User); ok {
		return user
	}
	return User{}
}

func (u Usecase) AllowedRole(allowedRole ...Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if user, ok := ctx.Keys["user"].(User); ok {
			userRole := Role(user.Role)
			for _, role := range allowedRole {
				if userRole == role {
					return
				}
			}
			utils.AbortAndResponseData(ctx, http.StatusUnauthorized, code.UNAUTHORIZED, code.ROLE_NO_PERMISSION)
			return
		}
		utils.AbortAndResponse(ctx, http.StatusUnauthorized, code.UNAUTHORIZED)
	}
}

func (u Usecase) AuthenticationNeeded() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := ExtractTokenID(ctx.Request)
		if err != nil {
			utils.AbortAndResponse(ctx, http.StatusUnauthorized, code.UNAUTHORIZED)
			return
		}
		user, err2 := u.GetRawUser(userId)
		if err2 != nil {
			utils.AbortAndResponseData(ctx, http.StatusUnauthorized, code.UNAUTHORIZED, err2.Error())
			return
		}
		if ctx.Keys == nil {
			ctx.Keys = map[string]interface{}{}
		}
		ctx.Keys["user"] = user
	}
}
