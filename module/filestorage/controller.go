package filestorage

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common/code"
	"github.com/ravielze/fuzzy-broccoli/common/utils"
	"github.com/ravielze/fuzzy-broccoli/middleware"
)

type FileController struct {
	usecase IFileUsecase
}

func NewFileController(router *gin.Engine, usecase IFileUsecase) IFileController {
	cont := FileController{
		usecase: usecase,
	}
	fileGroup := router.Group("/file")
	{
		fileGroup.POST("/add", middleware.GetAuthMiddleware(), cont.AddFile)
		fileGroup.GET("/:id", cont.GetFile)
	}
	return cont
}

func (c FileController) GetFile(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		utils.AbortAndResponseData(ctx, http.StatusUnprocessableEntity, code.UNCOMPATIBLE_JSON, "Parameter ID is missing")
		return
	}
	result, err := c.usecase.GetFile(id)
	if err != nil {
		utils.AbortAndResponseData(ctx, http.StatusBadRequest, code.LOGIC_ERROR, err.Error())
		return
	}
	utils.OKAndResponseData(ctx, result)
}

func (c FileController) AddFile(ctx *gin.Context) {
	var srlzr FileSerializer
	if err := ctx.ShouldBindJSON(&srlzr); err != nil {
		utils.AbortAndResponseData(ctx, http.StatusUnprocessableEntity, code.UNCOMPATIBLE_JSON, err.Error())
		return
	}
	user := middleware.GetUser(ctx)
	result, err := c.usecase.AddFile(user.ID, srlzr)
	if err != nil {
		utils.AbortAndResponseData(ctx, http.StatusBadRequest, code.LOGIC_ERROR, err.Error())
		return
	}
	utils.OKAndResponseData(ctx, result)
}
