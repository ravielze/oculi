package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi/common"
	"gorm.io/gorm"
)

const TOKEN_EXPIRED_TIME time.Duration = time.Hour * 3

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

func (u *User) BeforeSave(db *gorm.DB) error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
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
	Update(user User, item UpdateRequest) error
	RegisterAdmin(item RegisterRequest) (UserResponse, error)
	GetByID(userId uint) (UserResponse, error)
	GetRawUser(userId uint) (User, error)

	//Middleware thing
	GetUser(ctx *gin.Context) User
	AllowedRole(allowedRole ...Role) gin.HandlerFunc
	AuthenticationNeeded() gin.HandlerFunc
}

type IRepo interface {
	GetByID(userId uint) (User, error)
	GetByEmail(email string) (User, error)
	Create(item User) (User, error)
	Update(item User) error
}
