package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common/code"
	"github.com/ravielze/fuzzy-broccoli/common/serializer"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetUserID(ctx *gin.Context) uint64 {
	userId, err := strconv.ParseUint(ctx.Request.Header.Get("userId"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, serializer.NewResponse(http.StatusUnauthorized, code.UNAUTHORIZED))
		return 0
	}
	return userId
}
