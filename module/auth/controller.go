package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common/code"
	"github.com/ravielze/fuzzy-broccoli/common/utils"
)

type UserController struct {
	usecase IUserUsecase
}

func NewUserController(router *gin.Engine, usecase IUserUsecase) IUserController {
	cont := UserController{
		usecase: usecase,
	}
	userGroup := router.Group("/auth")
	{
		userGroup.POST("/login", cont.Login)
		userGroup.POST("/register", cont.Register)
		userGroup.GET("/check", func(ctx *gin.Context) {
			userId, err := ExtractTokenID(ctx.Request)
			if err != nil {
				utils.AbortAndResponse(ctx, http.StatusUnauthorized, code.UNAUTHORIZED)
				return
			}
			user, err2 := usecase.GetID(userId)
			if err2 != nil {
				utils.AbortAndResponseData(ctx, http.StatusUnauthorized, code.UNAUTHORIZED, err2.Error())
				return
			}
			ctx.Keys = map[string]interface{}{}
			ctx.Keys["user"] = user
		}, cont.Check)
	}
	return cont
}

func (u UserController) Register(ctx *gin.Context) {
	var srlzr RegisterSerializer
	if err := ctx.ShouldBindJSON(&srlzr); err != nil {
		utils.AbortAndResponseData(ctx, http.StatusUnprocessableEntity, code.UNCOMPATIBLE_JSON, err.Error())
		return
	}
	user, err := u.usecase.Register(srlzr)
	if err != nil {
		utils.AbortAndResponseData(ctx, http.StatusBadRequest, code.LOGIC_ERROR, err.Error())
		return
	}
	utils.OKAndResponseData(ctx, user)
}

func (u UserController) Login(ctx *gin.Context) {

	var srlzr LoginSerializer
	if err := ctx.ShouldBindJSON(&srlzr); err != nil {
		utils.AbortAndResponseData(ctx, http.StatusUnprocessableEntity, code.UNCOMPATIBLE_JSON, err.Error())
		return
	}
	user, token, err := u.usecase.Login(srlzr)
	if err != nil {
		utils.AbortAndResponseData(ctx, http.StatusBadRequest, code.LOGIC_ERROR, err.Error())
		return
	}
	utils.OKAndResponseData(ctx, struct {
		User  User   `json:"user_data"`
		Token string `json:"token"`
	}{
		User:  user,
		Token: token,
	})
}

func (u UserController) Check(ctx *gin.Context) {
	if user, ok := ctx.Keys["user"].(User); ok {
		utils.OKAndResponseData(ctx, user)
		return
	}
	utils.OKAndResponseData(ctx, User{})
}
