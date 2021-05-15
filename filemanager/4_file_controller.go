package filemanager

import (
	"github.com/gin-gonic/gin"
	cutils "github.com/ravielze/oculi/common/controller_utils"
	"github.com/ravielze/oculi/common/utils"
)

type Controller struct {
	uc IUsecase
}

func NewController(g *gin.Engine, uc IUsecase) IController {
	cont := Controller{
		uc: uc,
	}
	filemanagerGroup := g.Group("/file")
	{
		filemanagerGroup.GET("/id/:fileid", cont.GetFile)
		filemanagerGroup.GET("/group/:filegroup", cont.GetFilesByGroup)
	}
	return cont
}

func (cont Controller) GetFile(ctx *gin.Context) {
	ok, params, _ := cutils.NewControlChain(ctx).ParamBase36ToUUID("fileid").End()
	if ok {
		result, err := cont.uc.GetFile(params["fileid"])
		if err != nil {
			utils.AbortUsecaseError(ctx, err)
			return
		}
		utils.OKAndResponseData(ctx, result)
	}
}

func (cont Controller) GetFilesByGroup(ctx *gin.Context) {
	ok, params, _ := cutils.NewControlChain(ctx).Param("filegroup").End()
	if ok {
		result, err := cont.uc.GetFilesByGroup(params["filegroup"])
		if err != nil {
			utils.AbortUsecaseError(ctx, err)
			return
		}
		utils.OKAndResponseData(ctx, result)
	}
}
