package external

import "github.com/gin-gonic/gin"

func NewGin() (*gin.Engine, error) {
	return gin.Default(), nil
}
