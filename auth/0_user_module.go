package auth

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi/common"
	"gorm.io/gorm"
)

type Module struct {
	controller IController
	usecase    IUsecase
	repository IRepo
}

func (Module) Name() string {
	return "auth"
}

func (m Module) Usecase() IUsecase {
	return m.usecase
}

func (Module) Reset(db *gorm.DB) {
	db.Migrator().DropTable(&User{})
}

func NewModule(db *gorm.DB, g *gin.Engine) Module {
	repo := NewRepository(db)
	uc := NewUsecase(repo)
	cont := NewController(g, uc)

	db.AutoMigrate(&User{})

	hashedPassword, err := Hash(os.Getenv("ADMIN_DEFAULT_PASSWORD"))
	if err != nil {
		panic(err)
	}
	db.FirstOrCreate(&User{
		IntIDBase: common.IntIDBase{ID: 1},
		Email:     os.Getenv("ADMIN_DEFAULT_EMAIL"),
		Name:      "admin",
		Password:  hashedPassword,
		Role:      int16(ROLE_ADMIN),
	})

	return Module{
		controller: cont,
		usecase:    uc,
		repository: repo,
	}
}
