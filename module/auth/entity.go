package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common"
	"gorm.io/gorm"
)

type User struct {
	common.IDBase   `gorm:"embedded;embeddedPrefix:user_"`
	common.InfoBase `gorm:"embedded"`
	Email           string `gorm:"type:VARCHAR(320);uniqueIndex:,sort:asc,type:btree" json:"email"`
	Password        string `gorm:"type:VARCHAR(1024)" json:"password"`
}

func (u *User) BeforeSave(db *gorm.DB) error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

type IUserController interface {
	Register(ctx *gin.Context) // Register
	Login(ctx *gin.Context)    // Send jwt token with response
	Check(ctx *gin.Context)    //Buat midldeware cek token.
}

type IUserUsecase interface {
	Login(item LoginSerializer) (User, string, error)
	Register(item RegisterSerializer) (User, error)
}

type IUserRepo interface {
	Login(email, password string) (User, error)
	Register(email, password string) (User, error)
}

// func (u *User) Prepare() {
// 	u.ID = 0
// 	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
// 	u.CreatedAt = time.Now()
// 	u.UpdatedAt = time.Now()
// }

// func (u *User) Validate(action string) error {

// 	switch strings.ToLower(action) {
// 	case "update":
// 		if u.Password == "" {
// 			return errors.New("Required Password")
// 		}
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if err := checkmail.ValidateFormat(u.Email); err != nil {
// 			return errors.New("Invalid Email")
// 		}

// 		return nil
// 	case "login":
// 		if u.Password == "" {
// 			return errors.New("Required Password")
// 		}
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if err := checkmail.ValidateFormat(u.Email); err != nil {
// 			return errors.New("Invalid Email")
// 		}
// 		return nil

// 	default:
// 		if u.Password == "" {
// 			return errors.New("Required Password")
// 		}
// 		if u.Email == "" {
// 			return errors.New("Required Email")
// 		}
// 		if err := checkmail.ValidateFormat(u.Email); err != nil {
// 			return errors.New("Invalid Email")
// 		}
// 		return nil
// 	}
// }

// func (u *User) SaveUser(db *gorm.DB) (*User, error) {

// 	var err error
// 	err = db.Debug().Create(&u).Error
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	return u, nil
// }
