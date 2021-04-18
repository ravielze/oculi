package filestorage

import "github.com/ravielze/fuzzy-broccoli/module/auth"

type FileUsecase struct {
	repo IFileRepo
}

func NewFileUsecase(repo IFileRepo) IFileUsecase {
	return FileUsecase{repo: repo}
}

func (uc FileUsecase) GetLink(idFile string) (LinkFile, error) {
	return uc.repo.GetOneLinkByID(idFile)
}

func (uc FileUsecase) AddLink(user auth.User, item LinkFileSerializer) (LinkFile, error) {
	return uc.repo.AddOneLink(user.ID, item.Link, item.FileType, item.FileGroup)
}

func (uc FileUsecase) GetLocalStorage(idFile string) (LocalStorageFile, error) {
	return uc.repo.GetOneLocalStorageByID(idFile)
}

func (uc FileUsecase) AddLocalStorage(user auth.User, item LocalStorageFileSerializer) (LocalStorageFile, error) {
	return uc.repo.AddOneLocalStorage(user.ID, item.Attachment, item.FileGroup)
}

func (uc FileUsecase) GetUserFiles(user auth.User) ([]interface{}, error) {
	return uc.repo.GetUserFiles(user.ID)
}

func (uc FileUsecase) GetFile(idFile string) (interface{}, error) {
	base, err := uc.repo.GetFileBase(idFile)
	if err != nil {
		return nil, err
	}
	switch base.FileMethod {
	case LOCAL_STORAGE:
		local, err := uc.repo.GetOneLocalStorageByID(idFile)
		if err != nil {
			return nil, err
		}
		return local, nil
	case LINK_STORAGE:
		link, err := uc.repo.GetOneLinkByID(idFile)
		if err != nil {
			return nil, err
		}
		return link, nil
	default:
		return nil, nil
	}
}
