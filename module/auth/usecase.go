package auth

type UserUsecase struct {
	userrepo IUserRepo
}

func NewUserUsecase(repo IUserRepo) IUserUsecase {
	return UserUsecase{
		userrepo: repo,
	}
}

func (usecase UserUsecase) Login(item LoginSerializer) (User, string, error) {
	user, err := usecase.userrepo.Login(item.Email, item.Password)
	if err != nil {
		return User{}, "", err
	}
	token, err := CreateToken(uint64(user.ID))
	if err != nil {
		return User{}, "", err
	}
	return user, token, nil
}
func (usecase UserUsecase) Register(item RegisterSerializer) (User, error) {
	user, err := usecase.userrepo.Register(item.Email, item.Password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
