package filestorage

import (
	"database/sql"
	"mime/multipart"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common"
	"github.com/ravielze/fuzzy-broccoli/module/auth"
	"gorm.io/gorm"
)

type FileBase struct {
	common.UUIDBase `gorm:"embedded;embeddedPrefix:filebase_"`
	common.InfoBase `gorm:"embedded"`
	UserID          uint           `json:"user_id"`
	User            auth.User      `json:"-"`
	FileGroup       string         `gorm:"type:VARCHAR(128);index:,type:hash;" json:"-"`
	FileType        string         `gorm:"type:VARCHAR(128);" json:"file_type"`
	FileExt         sql.NullString `gorm:"type:VARCHAR(8);" json:"file_ext"`
	FileMethod      string         `gorm:"type:VARCHAR(8);" json:"method"`
}

func (f *FileBase) BeforeSave(db *gorm.DB) error {
	f.FileType = strings.ToLower(f.FileType)
	if f.FileExt.Valid {
		f.FileExt.String = strings.ToLower(f.FileExt.String)
	}
	return nil
}

func (FileBase) TableName() string {
	return "file_base"
}

type LinkFile struct {
	FileBaseID string   `gorm:"primaryKey;uniqueIndex:,sort:asc,type:btree;"`
	FileBase   FileBase `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Link       string   `gorm:"type:VARCHAR(512);" json:"link"`
}

func (LinkFile) TableName() string {
	return "link_file"
}

type LocalStorageFile struct {
	FileBaseID   string   `gorm:"primaryKey;uniqueIndex:,sort:asc,type:btree"`
	FileBase     FileBase `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RealFilename string   `gorm:"type:VARCHAR(512);index:,sort:asc,type:btree" json:"-"`
	Path         string   `gorm:"type:VARCHAR(1024)" json:"path"`
	Size         uint64   `json:"size"`
}

func (LocalStorageFile) TableName() string {
	return "local_storage_file"
}

type IFileController interface {
	GetFile(ctx *gin.Context)
	AddFile(ctx *gin.Context)
	GetUserFiles(ctx *gin.Context)
}

type IFileUsecase interface {
	GetFile(idFile string) (interface{}, error)
	GetLink(idFile string) (LinkFile, error)
	AddLink(user auth.User, item LinkFileSerializer) (LinkFile, error)
	GetLocalStorage(idFile string) (LocalStorageFile, error)
	AddLocalStorage(user auth.User, item LocalStorageFileSerializer) (LocalStorageFile, error)
	GetUserFiles(user auth.User) ([]interface{}, error)
}

type IFileRepo interface {
	GetOneLinkByID(idFile string) (LinkFile, error)
	AddOneLink(userId uint, link, fileType, fileGroup string) (LinkFile, error)
	GetOneLocalStorageByID(idFile string) (LocalStorageFile, error)
	AddOneLocalStorage(userId uint, attachment *multipart.FileHeader, fileGroup string) (LocalStorageFile, error)
	GetUserFiles(userId uint) ([]interface{}, error)
}
