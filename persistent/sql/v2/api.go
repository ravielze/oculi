package sqlv2

import (
	"context"
	"database/sql"

	"github.com/ravielze/oculi/errors"
	sqlOculi "github.com/ravielze/oculi/persistent/sql"

	"gorm.io/gorm"
)

func (i *Impl) Ping(ctx context.Context) error {
	db, err := i.Database.WithContext(ctx).DB()
	if err != nil {
		return err
	}
	return db.Ping()
}

func (i *Impl) Gorm() *gorm.DB {
	return i.Database
}

func (i *Impl) Dialector() gorm.Dialector {
	return i.Database.Dialector
}

func (i *Impl) Migrator() gorm.Migrator {
	return i.Database.Migrator()
}

func (i *Impl) Error() error {
	return errors.Convert(i.Database.Error)
}
func (i *Impl) RowsAffected() int64 {
	return i.Database.RowsAffected
}

func (i *Impl) AutoMigrate(dst ...interface{}) error {
	return errors.Convert(i.Database.AutoMigrate(dst...))
}

func (i *Impl) Association(column string) *gorm.Association {
	return i.Database.Association(column)
}

func (i *Impl) Statement() *gorm.Statement {
	return i.Database.Statement
}

// Session create new db session
func (i *Impl) Session(config *gorm.Session) sqlOculi.API {
	db := i.Database.Session(config)
	return i.copy(db)
}

// WithContext change current instance db's context to ctx
func (i *Impl) WithContext(ctx context.Context) sqlOculi.API {
	db := i.Database.WithContext(ctx)
	return i.copy(db)
}

// Debug start debug mode
func (i *Impl) Debug() sqlOculi.API {
	db := i.Database.Debug()
	return i.copy(db)
}

// Set store value with key into current db instance's context
func (i *Impl) Set(key string, value interface{}) sqlOculi.API {
	db := i.Database.Set(key, value)
	return i.copy(db)
}

// Get get value with key from current db instance's context
func (i *Impl) Get(key string) (interface{}, bool) {
	return i.Database.Get(key)
}

// InstanceSet store value with key into current db instance's context
func (i *Impl) InstanceSet(key string, value interface{}) sqlOculi.API {
	db := i.Database.InstanceSet(key, value)
	return i.copy(db)
}

// InstanceGet get value with key from current db instance's context
func (i *Impl) InstanceGet(key string) (interface{}, bool) {
	return i.Database.InstanceGet(key)
}

// AddError add error to db
func (i *Impl) AddError(err error) error {
	return errors.Convert(i.Database.AddError(err))
}

// DB returns `*sql.Database`
func (i *Impl) DB() (*sql.DB, error) {
	db, err := i.Database.DB()
	return db, errors.Convert(err)
}

func (i *Impl) SetupJoinTable(model interface{}, field string, joinTable interface{}) error {
	return errors.Convert(i.Database.SetupJoinTable(model, field, joinTable))
}

func (i *Impl) Use(plugin gorm.Plugin) error {
	return errors.Convert(i.Database.Use(plugin))
}
