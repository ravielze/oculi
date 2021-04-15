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
	//uc IUserUsecase
}

func NewUserController(router *gin.Engine) IUserController {
	cont := UserController{}
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
	return nil
}

func (u UserController) Register(ctx *gin.Context) {

}

func (u UserController) Login(ctx *gin.Context) {
	userId := 12
	token, err := CreateToken(uint64(userId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, serializer.NewResponse(http.StatusUnauthorized, code.UNAUTHORIZED))
		return
	}
	ctx.JSON(http.StatusOK, serializer.NewResponseData(http.StatusOK, code.OK, token))
}

func (u UserController) Check(ctx *gin.Context) {
	userId := GetUserID(ctx)
	fmt.Println(userId)
}
