package filestorage

import (
	"net/http"
	"strings"

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
		fileGroup.GET("/:id", cont.GetFile)
		fileGroup.POST("/", middleware.GetAuthMiddleware(), cont.AddFile)
		fileGroup.GET("/", middleware.GetAuthMiddleware(), cont.GetUserFiles)
	}
	return cont
}

func (cont FileController) GetFile(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		utils.AbortAndResponseData(ctx, http.StatusUnprocessableEntity, code.UNCOMPATIBLE_JSON, "Parameter ID is missing")
		return
	}
	result, err := cont.usecase.GetFile(id)
	if err != nil {
		utils.AbortAndResponseData(ctx, http.StatusBadRequest, code.LOGIC_ERROR, err.Error())
		return
	}
	getParam := ctx.Query("method")
	download := false
	if strings.EqualFold(getParam, "download") {
		download = true
	}
	srlzr, ok := result.(LocalStorageFile)
	if ok {
		ctx.Header("Content-Description", "File Transfer")
		ctx.Header("Content-Transfer-Encoding", "binary")
		if download {
			ctx.Header("Content-Disposition", "attachment; filename="+srlzr.RealFilename)
		} else {
			ctx.Header("Content-Disposition", "inline; filename="+srlzr.RealFilename)
		}
		ctx.Header("Content-Type", srlzr.FileBase.FileType)
		ctx.File("./storage/" + srlzr.Path)
		return
	} else {
		othersrlzr, ok2 := result.(LinkFile)
		if ok2 {
			ctx.JSON(http.StatusOK, othersrlzr)
			return
		}
	}
	utils.AbortAndResponseData(ctx, http.StatusBadRequest, code.UNKNOWN, "An unknown error has occured")
}

func (cont FileController) AddFile(ctx *gin.Context) {
	switch ctx.ContentType() {
	case "multipart/form-data":
		var srlzr LocalStorageFileSerializer
		if err := ctx.ShouldBind(&srlzr); err != nil {
			utils.AbortAndResponseData(ctx, http.StatusUnprocessableEntity, code.UNCOMPATIBLE_JSON, err.Error())
			return
		}
		result, err := cont.usecase.AddLocalStorage(middleware.GetUser(ctx), srlzr)
		if err != nil {
			utils.AbortAndResponseData(ctx, http.StatusBadRequest, code.LOGIC_ERROR, err.Error())
			return
		}
		utils.OKAndResponseData(ctx, result)
		return
	case "application/json":
		var srlzr LinkFileSerializer
		if err := ctx.ShouldBindJSON(&srlzr); err != nil {
			utils.AbortAndResponseData(ctx, http.StatusUnprocessableEntity, code.UNCOMPATIBLE_JSON, err.Error())
			return
		}
		result, err := cont.usecase.AddLink(middleware.GetUser(ctx), srlzr)
		if err != nil {
			utils.AbortAndResponseData(ctx, http.StatusBadRequest, code.LOGIC_ERROR, err.Error())
			return
		}
		utils.OKAndResponseData(ctx, result)
		return
	default:
		utils.AbortAndResponseData(ctx, http.StatusUnprocessableEntity, code.UNKNOWN_CONTENT_TYPE, "Content type is not supported. Allowed: application/json or multipart/form-data")
		return
	}
}

func (cont FileController) GetUserFiles(ctx *gin.Context) {
	result, err := cont.usecase.GetUserFiles(middleware.GetUser(ctx))
	if err != nil {
		utils.AbortAndResponseData(ctx, http.StatusBadRequest, code.LOGIC_ERROR, err.Error())
		return
	}
	utils.OKAndResponseData(ctx, result)
}
