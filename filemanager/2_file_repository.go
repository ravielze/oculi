package filemanager

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
	storage "github.com/ramadani/go-filestorage"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FileStorage struct {
	storageRoot string
	storage     *storage.Storage
}

type Repository struct {
	db *gorm.DB
	fs *FileStorage
}

func NewRepository(db *gorm.DB, fs *FileStorage) IRepo {
	return Repository{
		db: db,
		fs: fs,
	}
}

func (repo Repository) DeleteFile(idFile string) error {
	var deletedFile File
	if err := repo.db.
		Where("file_id = ?", idFile).
		First(&deletedFile).
		Error; err != nil {
		return err
	} else {
		if errdel := repo.db.
			Where("file_id = ?", deletedFile.ID).
			Delete(&File{}).
			Error; errdel != nil {
			return errdel
		}
		os.Remove(fmt.Sprintf("./%s/%s", repo.fs.storageRoot, deletedFile.Path))
		return nil
	}
}

func (repo Repository) GetFile(idFile string) (File, error) {
	var result File
	if err := repo.db.
		Where("file_id = ?", idFile).
		First(&result).
		Error; err != nil {
		return File{}, err
	}
	return result, nil
}

func (repo Repository) GetFilesByGroup(fileGroup string) ([]File, error) {
	var result []File
	if err := repo.db.
		Model(&File{}).
		Where("file_group = ?", fileGroup).
		Find(&result).
		Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (repo Repository) AddFile(userId uint, fileGroup string, attachment *multipart.FileHeader) (File, error) {
	fileExt := filepath.Ext(attachment.Filename)

	attachedFile, err := attachment.Open()
	if err != nil {
		return File{}, err
	}
	defer attachedFile.Close()

	fileType, err2 := mimetype.DetectReader(attachedFile)
	if err2 != nil {
		return File{}, err2
	}

	fileSize, err3 := attachedFile.Seek(0, 2)
	if err3 != nil {
		return File{}, err3
	}

	fileName := GenerateFileName(fileType.Extension())
	repo.fs.storage.PutFileAs(fileGroup, attachment, fileName)
	added := File{
		FileGroup:    fileGroup,
		FileType:     fileType.String(),
		FileExt:      fileExt,
		RealFilename: attachment.Filename,
		Path:         fmt.Sprintf("%s/%s", fileGroup, fileName),
		Size:         uint64(fileSize),
		OwnerID:      userId,
	}
	if errf := repo.db.
		Create(&added).
		Preload(clause.Associations).Error; errf != nil {
		return File{}, errf
	}
	return added, nil
}
