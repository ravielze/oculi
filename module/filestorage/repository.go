package filestorage

import (
	"database/sql"
	"mime/multipart"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
	storage "github.com/ramadani/go-filestorage"
	"gorm.io/gorm"
)

type FileRepository struct {
	db      *gorm.DB
	storage *storage.Storage
}

func NewFileRepository(db *gorm.DB, storage *storage.Storage) IFileRepo {
	return FileRepository{
		db:      db,
		storage: storage,
	}
}

func (repo FileRepository) GetOneLinkByID(idFile string) (LinkFile, error) {
	var result LinkFile
	if err := repo.db.Model(&LinkFile{}).Preload("FileBase").Where("file_base_id = ?", idFile).First(&result).Error; err != nil {
		return LinkFile{}, err
	}
	return result, nil
}

func (repo FileRepository) AddOneLink(userId uint, link, fileType, fileGroup string) (LinkFile, error) {
	added := LinkFile{
		FileBase: FileBase{
			UserID:     userId,
			FileGroup:  fileGroup,
			FileType:   fileType,
			FileExt:    sql.NullString{Valid: false},
			FileMethod: LINK_STORAGE,
		},
		Link: link,
	}
	if err := repo.db.Create(&added).Error; err != nil {
		return LinkFile{}, err
	}
	return added, nil
}

func (repo FileRepository) GetOneLocalStorageByID(idFile string) (LocalStorageFile, error) {
	var result LocalStorageFile
	if err := repo.db.Model(&LocalStorageFile{}).Preload("FileBase").Where("file_base_id = ?", idFile).First(&result).Error; err != nil {
		return LocalStorageFile{}, err
	}
	return result, nil
}

func (repo FileRepository) AddOneLocalStorage(userId uint, attachment *multipart.FileHeader, fileGroup string) (LocalStorageFile, error) {
	fileExt := filepath.Ext(attachment.Filename)

	attachedFile, buffErr := attachment.Open()
	if buffErr != nil {
		return LocalStorageFile{}, buffErr
	}

	fileType, mimeErr := mimetype.DetectReader(attachedFile)
	if mimeErr != nil {
		return LocalStorageFile{}, mimeErr
	}

	fileSize, sizeErr := attachedFile.Seek(0, 2)
	if sizeErr != nil {
		return LocalStorageFile{}, sizeErr
	}
	defer attachedFile.Close()

	fileName := GenerateFileName(fileExt)
	repo.storage.PutFileAs(fileGroup, attachment, fileName)

	added := LocalStorageFile{
		FileBase: FileBase{
			UserID:    userId,
			FileGroup: fileGroup,
			FileType:  fileType.String(),
			FileExt: sql.NullString{
				String: fileExt,
				Valid:  true,
			},
			FileMethod: LOCAL_STORAGE,
		},
		RealFilename: attachment.Filename,
		Path:         fileGroup + "/" + fileName,
		Size:         uint64(fileSize),
	}
	if err := repo.db.Create(&added).Error; err != nil {
		return LocalStorageFile{}, err
	}
	return added, nil
}

func (repo FileRepository) GetUserFiles(userId uint) ([]interface{}, error) {
	fileBases := repo.db.Model(&FileBase{}).Select("file_base_id").Where("user_id = ?", userId)
	var files1 []LocalStorageFile
	if err := repo.db.Model(&LocalStorageFile{}).Preload("FileBase").Where("file_base_id IN (?)", fileBases).Find(&files1).Error; err != nil {
		return nil, err
	}
	var files2 []LinkFile
	if err := repo.db.Model(&LinkFile{}).Preload("FileBase").Where("file_base_id IN (?)", fileBases).Find(&files2).Error; err != nil {
		return nil, err
	}
	var result []interface{}
	result = append(result, files1, files2)
	return result, nil
}

func (repo FileRepository) GetFileBase(idFile string) (FileBase, error) {
	var result FileBase
	if err := repo.db.Model(&FileBase{}).Where("file_base_id = ?", idFile).First(&result).Error; err != nil {
		return FileBase{}, err
	}
	return result, nil
}

func (repo FileRepository) GetFileIDByGroup(fileGroup string) ([]string, error) {
	var result []string
	if err := repo.db.Model(&FileBase{}).Select("file_base_id").Where("file_group = ?", fileGroup).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
