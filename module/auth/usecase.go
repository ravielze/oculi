package auth

import "errors"

type UserUsecase struct {
	userrepo IUserRepo
}

func NewUserUsecase(repo IUserRepo) IUserUsecase {
	return UserUsecase{
		userrepo: repo,
	}
}

func (uc UserUsecase) GetID(userId uint64) (User, error) {
	user, err := uc.userrepo.GetOneByID(userId)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (uc UserUsecase) Login(item LoginSerializer) (User, string, error) {
	user, err := uc.userrepo.GetOneByEmail(item.Email)
	if err != nil {
		return User{}, "", err
	}
	token, err := CreateToken(uint64(user.ID))
	if err != nil {
		return User{}, "", err
	}

	if err := VerifyPassword(user.Password, item.Password); err != nil {
		return User{}, "", errors.New("Password not match")
	}
	return user, token, nil
}

func (uc UserUsecase) Register(item RegisterSerializer) (User, error) {
	user, err := uc.userrepo.CreateOne(item.Email, item.Password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
