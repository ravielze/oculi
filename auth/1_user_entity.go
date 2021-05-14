package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi/common"
)

type User struct {
	common.IntIDBase      `gorm:"embedded;embeddedPrefix:user_"`
	common.InfoBase       `gorm:"embedded"`
	common.SoftDeleteBase `gorm:"embedded"`
	Email                 string `gorm:"type:VARCHAR(320);uniqueIndex:,sort:asc,type:btree"`
	Name                  string `gorm:"type:VARCHAR(512);"`
	Password              string `gorm:"type:VARCHAR(128)"`
	Role                  int16  `gorm:"<-:create;type:SMALLINT;"`
}

func (User) TableName() string {
	return "user"
}

type IController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Update(ctx *gin.Context)
	Check(ctx *gin.Context)
}

type IUsecase interface {
	Login(item LoginRequest) (UserTokenResponse, error)
	Register(item RegisterRequest) (UserResponse, error)
	Update(user User) error
	RegisterAdmin(item RegisterRequest) (UserResponse, error)
	GetByID(userId uint) (UserResponse, error)
}

type IRepo interface {
	GetByID(userId uint) (User, error)
	GetByEmail(email string) (User, error)
	Create(item User) (User, error)
	Update(item User) error
}
