package filestorage

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FileModule struct {
	controller IFileController
	Usecase    IFileUsecase
	repo       IFileRepo
}

func (FileModule) Name() string {
	return "File Storage Module"
}

func NewFileModule(db *gorm.DB, g *gin.Engine) FileModule {
	r := NewFileRepository(db)
	u := NewFileUsecase(r)
	c := NewFileController(g, u)

	db.AutoMigrate(&File{})

	return FileModule{
		controller: c,
		Usecase:    u,
		repo:       r,
	}
}
