package filemanager

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	storage "github.com/ramadani/go-filestorage"
	"github.com/ravielze/oculi/common"
	"gorm.io/gorm"
)

type Module struct {
	controller IController
	usecase    IUsecase
	repository IRepo
}

func (Module) Name() string {
	return "filemanager"
}

func (Module) Reset(db *gorm.DB) {
	db.Migrator().DropTable(&File{})
}

func NewModule(db *gorm.DB, g *gin.Engine) Module {
	storageRoot := os.Getenv("FILE_MANAGER_ROOT")
	storageComponent := storage.NewStorage(&storage.Config{
		Root: storageRoot,
		URL:  "",
	})
	fileStorage := &FileStorage{
		storageRoot: storageRoot,
		storage:     storageComponent,
	}

	repo := NewRepository(db, fileStorage)
	uc := NewUsecase(repo)
	cont := NewController(g, uc)

	db.AutoMigrate(&File{})

	if _, err := os.Stat(fmt.Sprintf("./%s/default", storageRoot)); os.IsNotExist(err) {
		storageComponent.MakeDir("default")
		fileName, errd := DownloadFile("https://i.ibb.co/TBjtryF/default.jpg", fmt.Sprintf("./%s/default", storageRoot))
		if errd != nil {
			panic(errd.Error())
		}
		db.FirstOrCreate(&File{
			UUIDBase:     common.UUIDBase{ID: "default"},
			FileGroup:    "default",
			FileType:     "image/jpg",
			FileExt:      ".jpg",
			RealFilename: "default.jpg",
			Path:         fileName,
			Size:         7007,
			OwnerID:      1,
		})
		//TODO generator env, filemanager, auth
		//add admin default to user
	}

	return Module{
		controller: cont,
		usecase:    uc,
		repository: repo,
	}
}
