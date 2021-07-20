package sqlv2

import (
	"time"

	"github.com/pkg/errors"
	"github.com/ravielze/oculi/logs/zap"
	"github.com/ravielze/oculi/persistent/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewClient(dl gorm.Dialector, options ...sql.ConnectionOption) (sql.API, error) {
	option := sql.ConnectionOptions{
		MaxIdleConnection: 10,
		MaxOpenConnection: 200,
		ConnMaxLifetime:   time.Hour,
		LogMode:           false,
	}
	sql.WithLogger(zap.DefaultLog()).Apply(&option)
	for _, opt := range options {
		opt.Apply(&option)
	}

	logLevel := []logger.LogLevel{
		logger.Silent,
		logger.Info,
		logger.Info,
		logger.Warn,
		logger.Error,
		logger.Silent,
		logger.Silent,
		logger.Silent,
	}[option.Logger().Level()]

	if !option.LogMode {
		logLevel = logger.Silent
	}

	db, err := gorm.Open(dl, &gorm.Config{
		Logger: logger.New(option.Logger(), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      logLevel,
			Colorful:      true,
		}),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database connection!")
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database connection!")
	}

	sqlDb.SetMaxIdleConns(option.MaxIdleConnection)
	sqlDb.SetMaxOpenConns(option.MaxOpenConnection)
	sqlDb.SetConnMaxLifetime(option.ConnMaxLifetime)

	return &Impl{
		Database: db,
	}, nil
}
