package auth

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
	authGroup := g.Group("/auth")
	{
		authGroup.POST("/login", cont.Login)
		authGroup.POST("/register", cont.Register)
		authGroup.GET("/", uc.AuthenticationNeeded(), cont.Check)
	}
	return cont
}

func (cont Controller) Check(ctx *gin.Context) {
	utils.OKAndResponseData(ctx, cont.uc.GetUser(ctx))
}

func (cont Controller) Login(ctx *gin.Context) {
	var obj LoginRequest
	ok, _, _ := cutils.NewControlChain(ctx).BindJSON(&obj).End()
	if ok {
		result, err := cont.uc.Login(obj)
		if err != nil {
			utils.AbortUsecaseError(ctx, err)
			return
		}
		utils.OKAndResponseData(ctx, result)
	}
}

func (cont Controller) Register(ctx *gin.Context) {
	var obj RegisterRequest
	ok, _, _ := cutils.NewControlChain(ctx).BindJSON(&obj).End()
	if ok {
		result, err := cont.uc.Register(obj)
		if err != nil {
			utils.AbortUsecaseError(ctx, err)
			return
		}
		utils.OKAndResponseData(ctx, result)
	}
}

func (cont Controller) Update(ctx *gin.Context) {
	var obj UpdateRequest
	ok, _, _ := cutils.NewControlChain(ctx).BindJSON(&obj).End()
	if ok {
		user := cont.uc.GetUser(ctx)
		err := cont.uc.Update(user, obj)
		if err != nil {
			utils.AbortUsecaseError(ctx, err)
			return
		}
		utils.OKAndResponse(ctx)
	}
}
