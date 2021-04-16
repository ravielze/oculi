package dvlp

import (
	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/middleware"
	"gorm.io/gorm"
)

type DevModule struct {
}

func (DevModule) Name() string {
	return "Development Module"
}

func NewDevModule(db *gorm.DB, g *gin.Engine) DevModule {
	devGroup := g.Group("/development")
	{
		devGroup.GET("/test", middleware.GetStaticTokenMiddleware())
	}
	return DevModule{}
}
