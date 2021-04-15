package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common/code"
	"github.com/ravielze/fuzzy-broccoli/common/serializer"
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
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, serializer.NewResponse(http.StatusUnauthorized, code.UNAUTHORIZED))
				return
			}
			ctx.Request.Header.Set("userId", strconv.FormatUint(userId, 10))
		}, cont.Check)
	}
	return cont
}

func (u UserController) Register(ctx *gin.Context) {
	var j RegisterSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, serializer.NewResponse(http.StatusBadRequest, code.UNCOMPATIBLE_JSON))
		return
	}
	user, err := u.usecase.Register(j)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, serializer.NewResponse(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, serializer.NewResponseData(http.StatusOK, code.OK, user))
}

func (u UserController) Login(ctx *gin.Context) {

	var j LoginSerializer
	if err := ctx.ShouldBindJSON(&j); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, serializer.NewResponse(http.StatusBadRequest, code.UNCOMPATIBLE_JSON))
		return
	}
	user, token, err := u.usecase.Login(j)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, serializer.NewResponse(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, serializer.NewResponseData(http.StatusOK, code.OK, struct {
		user  User
		token string
	}{
		user:  user,
		token: token,
	}))
}

func (u UserController) Check(ctx *gin.Context) {
	userId := GetUserID(ctx)
	fmt.Println(userId)
}
