package filemanager

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepo {
	return Repository{db: db}
}

func (repo Repository) AddFile(file File) (File, error) {
	panic("not implemented")
}

func (repo Repository) DeleteFile(idFile string) error {
	panic("not implemented")
}

func (repo Repository) GetFile(idFile string) (File, error) {
	panic("not implemented")
}

func (repo Repository) GetFilesByGroup(fileGroup string) ([]File, error) {
	panic("not implemented")
}
