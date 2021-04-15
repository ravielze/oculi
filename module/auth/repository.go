package auth

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepo {
	return UserRepository{
		db: db,
	}
}

func (repo UserRepository) Login(email, password string) (User, error) {
	var user User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, errors.New("No user with that email")
		}
		return User{}, err
	}

	if err := VerifyPassword(user.Password, password); err != nil {
		return User{}, errors.New("Password not match")
	}
	return user, nil
}
func (repo UserRepository) Register(email, password string) (User, error) {
	var user User
	var dummyUser User
	user.Email = email
	hashedPassword, err := Hash(password)
	if err != nil {
		return User{}, err
	}
	user.Password = string(hashedPassword)
	repo.db.Where("email = ?", email).First(&dummyUser)
	if (User{} != dummyUser) {
		return User{}, errors.New("User with that email already exist")
	}
	if err := repo.db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
