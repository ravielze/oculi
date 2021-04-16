package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthModule struct {
	controller IUserController
	Usecase    IUserUsecase
	repo       IUserRepo
}

func NewAuthModule(db *gorm.DB, g *gin.Engine) AuthModule {
	userrepo := NewUserRepository(db)
	userusecase := NewUserUsecase(userrepo)
	usercontroller := NewUserController(g, userusecase)

	db.AutoMigrate(&User{})

	return AuthModule{
		controller: usercontroller,
		Usecase:    userusecase,
		repo:       userrepo,
	}
}
