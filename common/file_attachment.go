package common

import "mime/multipart"

type FileAttachment struct {
	Attachment *multipart.FileHeader `json:"-" form:"attachment" binding:"required"`
}
