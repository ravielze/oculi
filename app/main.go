package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi/oculi"
	"gorm.io/gorm"
)

func main() {
	oculi.New("Testing", func(db *gorm.DB, g *gin.Engine) {

	}, func(db *gorm.DB, g *gin.Engine) {

	})
}
