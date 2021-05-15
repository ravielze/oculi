package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi"
	"github.com/ravielze/oculi/auth"
	"github.com/ravielze/oculi/common/essentials"
	"github.com/ravielze/oculi/common/middleware"
	mm "github.com/ravielze/oculi/common/module"
	"github.com/ravielze/oculi/filemanager"
	"gorm.io/gorm"
)

func main() {
	oculi.New("App_Name", func(db *gorm.DB, g *gin.Engine) {
		middleware.InstallCors(g, []string{"http://localhost:3000", "https://example.com"})
		middleware.InstallDefaultLimiter(g)
		// Add your middleware here
	}, func(db *gorm.DB, g *gin.Engine) {
		mm.AddModule(essentials.NewModule(db, g))
		mm.AddModule(auth.NewModule(db, g))
		mm.AddModule(filemanager.NewModule(db, g))
	})
}
