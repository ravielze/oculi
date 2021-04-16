package filestorage

import (
	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common"
	"github.com/ravielze/fuzzy-broccoli/module/auth"
)

type File struct {
	common.UUIDBase `gorm:"embedded;embeddedPrefix:file_"`
	common.InfoBase `gorm:"embedded"`
	RealFilename    string    `gorm:"VARCHAR(512);uniqueIndex:,sort:asc,type:btree" json:"-"`
	Link            string    `gorm:"VARCHAR(512)" json:"link"`
	FileType        string    `gorm:"VARCHAR(128)" json:"type"`
	Size            uint64    `json:"size"`
	UserID          uint      `json:"user_id"`
	User            auth.User `json:"-"`
}

func (File) TableName() string {
	return "file"
}

type IFileController interface {
	GetFile(ctx *gin.Context)
	AddFile(ctx *gin.Context)
}

type IFileUsecase interface {
	GetFile(idFile string) (File, error)
	AddFile(userId uint, item FileSerializer) (File, error)
}

type IFileRepo interface {
	GetOneByID(idFile string) (File, error)
	AddOne(userId uint, realFilename, link, fileType string, size uint64) (File, error)
}
