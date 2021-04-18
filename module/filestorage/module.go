package filestorage

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	storage "github.com/ramadani/go-filestorage"
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

func (FileModule) Reset(db *gorm.DB) {
	db.Migrator().DropTable(&FileBase{})
	db.Migrator().DropTable(&LinkFile{})
	db.Migrator().DropTable(&LocalStorageFile{})
	err := os.RemoveAll("storage")
	if err != nil {
		fmt.Println(err)
	}
}

func NewFileModule(db *gorm.DB, g *gin.Engine) FileModule {
	localStorage := storage.NewStorage(&storage.Config{
		Root: "storage",
	})
	r := NewFileRepository(db, localStorage)
	u := NewFileUsecase(r)
	c := NewFileController(g, u)

	db.AutoMigrate(&FileBase{})
	db.AutoMigrate(&LinkFile{})
	db.AutoMigrate(&LocalStorageFile{})

	return FileModule{
		controller: c,
		Usecase:    u,
		repo:       r,
	}
}
