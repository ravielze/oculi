package auth

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepo {
	return Repository{db: db}
}

func (repo Repository) A(a int32, b float32) (error, int) {

	return nil, 10
}

func (r Repository) D(b string) int {
	if true {

	}
	return 30
}
