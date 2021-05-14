package filemanager

import "github.com/ravielze/oculi/common"

type Usecase struct {
	repo IRepo
}

func NewUsecase(repo IRepo) IUsecase {
	return Usecase{repo: repo}
}

func (uc Usecase) AddFile(item common.FileAttachment) (FileResponse, error) {
	panic("not implemented")
}

func (uc Usecase) DeleteFile(idFile string) error {
	panic("not implemented")
}

func (uc Usecase) GetFile(idFile string) (FileResponse, error) {
	panic("not implemented")
}

func (uc Usecase) GetFilesByGroup(fileGroup string) ([]FileResponse, error) {
	panic("not implemented")
}
