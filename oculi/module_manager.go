package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi/common/essentials"
	"github.com/ravielze/oculi/common/middleware"
	"github.com/ravielze/oculi/common/module"
	"gorm.io/gorm"
)

func NewModule(db *gorm.DB, g *gin.Engine, devMode bool) {
	middleware.InstallCors(g, []string{"http://localhost:3000", "https://example.com"})
	middleware.InstallDefaultLimiter(g)

	essentialsModule := essentials.NewModule(db, g)

	module.AddModule(essentialsModule)
}
