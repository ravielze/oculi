package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InstallGZipCompressor(g *gin.Engine) {
	g.Use(gzip.Gzip(gzip.DefaultCompression))
}
