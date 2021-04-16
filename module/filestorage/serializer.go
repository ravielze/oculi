package filestorage

type FileSerializer struct {
	RealFilename string `json:"file_name" binding:"required,lte=512"`
	Link         string `json:"link" binding:"required,lte=512"`
	FileType     string `json:"type" binding:"required,lte=128"`
	Size         uint64 `json:"size" binding:"required,gte=0"`
}
