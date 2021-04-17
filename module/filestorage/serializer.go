package filestorage

import "mime/multipart"

type LinkFileSerializer struct {
	Link      string `json:"link" binding:"required,lte=512"`
	FileType  string `json:"file_type" binding:"required,lte=128"`
	FileGroup string `json:"file_group" binding:"required,lte=128"`
}

type LocalStorageFileSerializer struct {
	Attachment *multipart.FileHeader `json:"-" form:"attachment" binding:"required"`
	FileGroup  string                `json:"-" form:"file_group" binding:"required,lte=128"`
}
