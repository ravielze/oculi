package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthModule struct {
	controller IUserController
	usecase    IUserUsecase
	repo       IUserRepo
}

func NewAuthModule(db *gorm.DB, g *gin.Engine) AuthModule {
	userrepo := NewUserRepository(db)
	userusecase := NewUserUsecase(userrepo)
	usercontroller := NewUserController(g, userusecase)

	if db != nil {
		fmt.Println("Ada db")
	}
	err := db.AutoMigrate(&User{})
	fmt.Println(err.Error())

	return AuthModule{
		controller: usercontroller,
		usecase:    userusecase,
		repo:       userrepo,
	}
}
