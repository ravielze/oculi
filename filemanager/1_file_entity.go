package filemanager

import (
	"mime/multipart"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/oculi/common"
	"gorm.io/gorm"
)

type File struct {
	common.UUIDBase `gorm:"embedded;embeddedPrefix:file_"`
	common.InfoBase `gorm:"embedded"`
	FileGroup       string `gorm:"type:VARCHAR(256);index:,type:hash;"`
	FileType        string `gorm:"type:VARCHAR(256);" json:"file_type"`
	FileExt         string `gorm:"type:VARCHAR(16);" json:"file_ext"`
	RealFilename    string `gorm:"type:VARCHAR(512);index:,sort:asc,type:btree" json:"-"`
	Path            string `gorm:"type:VARCHAR(1024)" json:"path"`
	Size            uint64 `json:"size"`
}

func (f *File) BeforeSave(db *gorm.DB) error {
	f.FileType = strings.ToLower(f.FileType)
	f.FileExt = strings.ToLower(f.FileExt)
	return nil
}

func (f *File) BeforeUpdate(db *gorm.DB) error {
	f.FileType = strings.ToLower(f.FileType)
	f.FileExt = strings.ToLower(f.FileExt)
	return nil
}

func (File) TableName() string {
	return "file"
}

type IController interface {
	GetFile(ctx *gin.Context)
	GetFilesByGroup(ctx *gin.Context)
}

type IUsecase interface {
	GetFile(idFile string) (FileResponse, error)
	GetFilesByGroup(fileGroup string) ([]FileResponse, error)
	AddFile(item common.FileAttachment) (FileResponse, error)
	DeleteFile(idFile string) error
}

type IRepo interface {
	GetFile(idFile string) (File, error)
	GetFilesByGroup(fileGroup string) ([]File, error)
	AddFile(file File, attachment *multipart.FileHeader) (File, error)
	DeleteFile(idFile string) error
}
