package auth

type Usecase struct {
	repo IRepo
}

func NewUsecase(repo IRepo) IUsecase {
	return Usecase{repo: repo}
}
