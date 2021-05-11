package auth

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepo {
	return Repository{db: db}
}

func (r Repository) D(b string) int {
	if true {

	}
	return 30
}

func (repo Repository) A(a int, b float32) (error, int) {

	return nil, 10

}

func (repo Repository) B(b string, d User) User {
	panic("not implemented")
}

func (repo Repository) E(b string) (error, int) {
	panic("not implemented")
}

func (repo Repository) F(c int) {
	panic("not implemented")
}

func (repo Repository) G() User {
	panic("not implemented")
}

func (repo Repository) H() (User, error) {
	panic("not implemented")
}

func (repo Repository) I(a string) {
	panic("not implemented")
}

func (repo Repository) C(b string, c int) {
	return

}
