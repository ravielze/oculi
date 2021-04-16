package filestorage

type FileUsecase struct {
	repo IFileRepo
}

func NewFileUsecase(repo IFileRepo) IFileUsecase {
	return FileUsecase{repo: repo}
}

func (u FileUsecase) GetFile(idFile string) (File, error) {
	return u.repo.GetOneByID(idFile)
}

func (u FileUsecase) AddFile(userId uint, item FileSerializer) (File, error) {
	return u.repo.AddOne(userId, item.RealFilename, item.Link, item.FileType, item.Size)
}
