package filemanager

import "github.com/ravielze/oculi/common/radix36"

type FileResponse struct {
	ID       string `json:"file_id"`
	FileType string `json:"file_type"`
}

func (f File) Convert() FileResponse {
	return FileResponse{
		ID:       radix36.MustEncodeUUID(f.ID),
		FileType: f.FileType,
	}
}
