package auth

type User struct {
}

func (User) TableName() string {
	return "user"
}

type IController interface {
}

type IUsecase interface {
}

type IRepo interface {
	A(a int, b float32) (error, int)
	C(b string, c int)
	B(b string, d User) User
	G() User
	F(c int)
	D(b string) int
	E(b string) (error, int)
	H() (User, error)
	I(a string)
}
