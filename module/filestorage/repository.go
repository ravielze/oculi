package filestorage

import "gorm.io/gorm"

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) IFileRepo {
	return FileRepository{db: db}
}

func (f FileRepository) GetOneByID(idFile string) (File, error) {
	var result File
	if err := f.db.Model(&File{}).Where("file_id = ?", idFile).First(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

func (f FileRepository) AddOne(userId uint, realFilename, link, fileType string, size uint64) (File, error) {
	addedFile := File{
		RealFilename: realFilename,
		Link:         link,
		FileType:     fileType,
		Size:         size,
		UserID:       userId,
	}
	if err := f.db.Create(&addedFile).Error; err != nil {
		return File{}, err
	}
	return addedFile, nil
}
