package auth

import (
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepo {
	return Repository{db: db}
}

func (repo Repository) Create(item User) (User, error) {
	user := item
	if err := repo.db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (repo Repository) GetByEmail(email string) (User, error) {
	var user User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, errors.New("no user with that email")
		}
		return User{}, err
	}
	return user, nil
}

func (repo Repository) GetByID(userId uint) (User, error) {
	var user User
	if err := repo.db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, errors.New("no user with that id")
		}
		return User{}, err
	}
	return user, nil
}

func (repo Repository) Update(item User) error {
	updateUser := item
	if err := repo.db.
		Model(&User{}).
		Where("user_id = ?", item.ID).
		Updates(&updateUser).
		Error; err != nil {
		return err
	}
	return nil
}
