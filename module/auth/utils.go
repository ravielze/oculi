package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common/code"
	"github.com/ravielze/fuzzy-broccoli/common/serializer"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetUserID(ctx *gin.Context, uc *IUserUsecase) User {
	userId, err := strconv.ParseUint(ctx.Request.Header.Get("userId"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, serializer.NewResponse(http.StatusUnauthorized, code.UNAUTHORIZED))
		return User{}
	}
	user, err2 := (*uc).GetID(userId)
	if err2 != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, serializer.NewResponseData(http.StatusUnauthorized, code.UNAUTHORIZED, err2.Error()))
		return User{}
	}
	return user
}
