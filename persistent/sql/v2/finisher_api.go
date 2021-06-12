package sqlv2

import (
	"database/sql"

	sqlOculi "github.com/ravielze/oculi/persistent/sql"
	"gorm.io/gorm"
)

// Create insert the value into database
func (i *Impl) Create(value interface{}) sqlOculi.API {
	db := i.Database.Create(value)
	return i.copy(db)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (i *Impl) Save(value interface{}) sqlOculi.API {
	db := i.Database.Save(value)
	return i.copy(db)
}

// First find first record that match given conditions, order by primary key
func (i *Impl) First(dest interface{}, conds ...interface{}) sqlOculi.API {
	db := i.Database.First(dest, conds...)
	return i.copy(db)
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (i *Impl) Take(dest interface{}, conds ...interface{}) sqlOculi.API {
	db := i.Database.Take(dest, conds...)
	return i.copy(db)
}

// Last find last record that match given conditions, order by primary key
func (i *Impl) Last(dest interface{}, conds ...interface{}) sqlOculi.API {
	db := i.Database.Last(dest, conds...)
	return i.copy(db)
}

// Find find records that match given conditions
func (i *Impl) Find(dest interface{}, conds ...interface{}) sqlOculi.API {
	db := i.Database.Find(dest, conds...)
	return i.copy(db)
}

// FindInBatches find records in batches
func (i *Impl) FindInBatches(dest interface{}, batchSize int, fc func(tx sqlOculi.API, batch int) error) sqlOculi.API {
	nfc := func(tx *gorm.DB, batch int) error {
		return fc(i, batchSize)
	}
	db := i.Database.FindInBatches(dest, batchSize, nfc)
	return i.copy(db)
}

func (i *Impl) FirstOrInit(dest interface{}, conds ...interface{}) sqlOculi.API {
	db := i.Database.FirstOrInit(dest, conds...)
	return i.copy(db)
}

func (i *Impl) FirstOrCreate(dest interface{}, conds ...interface{}) sqlOculi.API {
	db := i.Database.FirstOrCreate(dest, conds...)
	return i.copy(db)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (i *Impl) Update(column string, value interface{}) sqlOculi.API {
	db := i.Database.Update(column, value)
	return i.copy(db)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (i *Impl) Updates(values interface{}) sqlOculi.API {
	db := i.Database.Updates(values)
	return i.copy(db)
}

func (i *Impl) UpdateColumn(column string, value interface{}) sqlOculi.API {
	db := i.Database.UpdateColumn(column, value)
	return i.copy(db)
}

func (i *Impl) UpdateColumns(values interface{}) sqlOculi.API {
	db := i.Database.UpdateColumns(values)
	return i.copy(db)
}

// Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition
func (i *Impl) Delete(value interface{}, conds ...interface{}) sqlOculi.API {
	db := i.Database.Delete(value, conds...)
	return i.copy(db)
}

func (i *Impl) Count(count *int64) sqlOculi.API {
	db := i.Database.Count(count)
	return i.copy(db)
}

func (i *Impl) Row() *sql.Row {
	return i.Database.Row()
}

func (i *Impl) Rows() (*sql.Rows, error) {
	sqlRows, err := i.Database.Rows()
	//TODO convert error
	return sqlRows, err
}

// Scan scan value to a struct
func (i *Impl) Scan(dest interface{}) sqlOculi.API {
	db := i.Database.Scan(dest)
	return i.copy(db)
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Find(&users).Pluck("age", &ages)
func (i *Impl) Pluck(column string, dest interface{}) sqlOculi.API {
	db := i.Database.Pluck(column, dest)
	return i.copy(db)
}

func (i *Impl) ScanRows(rows *sql.Rows, dest interface{}) error {
	//TODO convert error
	return i.Database.ScanRows(rows, dest)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (i *Impl) Transaction(fc func(sqlOculi.API) error, opts ...*sql.TxOptions) error {
	nfc := func(db *gorm.DB) error {
		return fc(i)
	}
	//TODO convert error
	return i.Database.Transaction(nfc, opts...)
}

// Begin begins a transaction
func (i *Impl) Begin(opts ...*sql.TxOptions) sqlOculi.API {
	db := i.Database.Begin(opts...)
	return i.copy(db)
}

// Commit commit a transaction
func (i *Impl) Commit() sqlOculi.API {
	db := i.Database.Commit()
	return i.copy(db)
}

// Rollback rollback a transaction
func (i *Impl) Rollback() sqlOculi.API {
	db := i.Database.Rollback()
	return i.copy(db)
}

func (i *Impl) SavePoint(name string) sqlOculi.API {
	db := i.Database.SavePoint(name)
	return i.copy(db)
}

func (i *Impl) RollbackTo(name string) sqlOculi.API {
	db := i.Database.RollbackTo(name)
	return i.copy(db)
}

// Exec execute raw sql
func (i *Impl) Exec(sql string, values ...interface{}) sqlOculi.API {
	db := i.Database.Exec(sql, values...)
	return i.copy(db)
}
