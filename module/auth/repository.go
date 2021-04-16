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

func (repo UserRepository) GetOneByEmail(email string) (User, error) {
	var user User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, errors.New("No user with that email")
		}
		return User{}, err
	}
	return user, nil
}

func (repo UserRepository) GetOneByID(userId uint64) (User, error) {
	var user User
	if err := repo.db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, errors.New("No user with that id")
		}
		return User{}, err
	}
	return user, nil
}

func (repo UserRepository) CreateOne(email, password string) (User, error) {
	var user User
	user.Email = email
	user.Password = password
	// repo.db.Where("email = ?", email).First(&dummyUser)
	// if (User{} != dummyUser) {
	// 	return User{}, errors.New("User with that email already exist")
	// }
	if err := repo.db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
