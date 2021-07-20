package sql

import (
	"context"
	"database/sql"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type (
	Config interface {
		// GORM perform single create, update, delete operations in transactions by default to ensure database data integrity
		// You can disable it by setting `SkipDefaultTransaction` to true
		SkipDefaultTransaction() bool
		// NamingStrategy tables, columns naming strategy
		NamingStrategy() schema.Namer
		// FullSaveAssociations full save associations
		FullSaveAssociations() bool
		// Logger
		Logger() logger.Interface
		// NowFunc the function to be used when creating a new timestamp
		NowFunc() func() time.Time
		// DryRun generate sql without execute
		DryRun() bool
		// PrepareStmt executes the given query in cached statement
		PrepareStmt() bool
		// DisableAutomaticPing
		DisableAutomaticPing() bool
		// DisableForeignKeyConstraintWhenMigrating
		DisableForeignKeyConstraintWhenMigrating() bool
		// AllowGlobalUpdate allow global update
		AllowGlobalUpdate() bool
		// ClauseBuilders clause builder
		ClauseBuilders() map[string]clause.ClauseBuilder
		// ConnPool db conn pool
		ConnPool() gorm.ConnPool
		// Plugins registered plugins
		Plugins() map[string]gorm.Plugin
	}

	ChainableAPI interface {
		Model(value interface{}) API

		// Clauses Add clauses
		Clauses(conds ...clause.Expression) API

		// Table specify the table you would like to run db operations
		Table(name string, args ...interface{}) API

		// Distinct specify distinct fields that you want querying
		Distinct(args ...interface{}) API

		// Select specify fields that you want when querying, creating, updating
		Select(query interface{}, args ...interface{}) API

		// Omit specify fields that you want to ignore when creating, updating and querying
		Omit(columns ...string) API

		// Where add conditions
		Where(query interface{}, args ...interface{}) API

		// Not add NOT conditions
		Not(query interface{}, args ...interface{}) API

		// Or add OR conditions
		Or(query interface{}, args ...interface{}) API

		// Joins specify Joins conditions
		//     db.Joins("Account").Find(&user)
		//     db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
		Joins(query string, args ...interface{}) API

		// Group specify the group method on the find
		Group(name string) API

		// Having specify HAVING conditions for GROUP BY
		Having(query interface{}, args ...interface{}) API

		// Order specify order when retrieve records from database
		//     db.Order("name DESC")
		//     db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}, Desc: true})
		Order(value interface{}) API

		// Limit specify the number of records to be retrieved
		Limit(limit int) API

		// Offset specify the number of records to skip before starting to return the records
		Offset(offset int) API

		// Scopes pass current database connection to arguments `func(DB) DB`, which could be used to add conditions dynamically
		//     func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
		//         return db.Where("amount > ?", 1000)
		//     }
		//
		//     func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
		//         return func (db *gorm.DB) *gorm.DB {
		//             return db.Scopes(AmountGreaterThan1000).Where("status in (?)", status)
		//         }
		//     }
		//
		//     db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
		Scopes(funcs ...func(db API) API) API

		// Preload preload associations with given conditions
		//    db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
		Preload(query string, args ...interface{}) API

		Attrs(attrs ...interface{}) API

		Assign(attrs ...interface{}) API

		Unscoped() API

		Raw(sql string, values ...interface{}) API
	}

	FinisherAPI interface {
		// Create insert the value into database
		Create(value interface{}) API

		// Save update value in database, if the value doesn't have primary key, will insert it
		Save(value interface{}) API

		// First find first record that match given conditions, order by primary key
		First(dest interface{}, conds ...interface{}) API

		// Take return a record that match given conditions, the order will depend on the database implementation
		Take(dest interface{}, conds ...interface{}) API

		// Last find last record that match given conditions, order by primary key
		Last(dest interface{}, conds ...interface{}) API

		// Find find records that match given conditions
		Find(dest interface{}, conds ...interface{}) API

		// FindInBatches find records in batches
		FindInBatches(dest interface{}, batchSize int, fc func(tx API, batch int) error) API

		FirstOrInit(dest interface{}, conds ...interface{}) API

		FirstOrCreate(dest interface{}, conds ...interface{}) API

		// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
		Update(column string, value interface{}) API

		// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
		Updates(values interface{}) API

		UpdateColumn(column string, value interface{}) API

		UpdateColumns(values interface{}) API

		// Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition
		Delete(value interface{}, conds ...interface{}) API

		Count(count *int64) API

		Row() *sql.Row

		Rows() (*sql.Rows, error)

		// Scan scan value to a struct
		Scan(dest interface{}) API

		// Pluck used to query single column from a model as a map
		//     var ages []int64
		//     db.Find(&users).Pluck("age", &ages)
		Pluck(column string, dest interface{}) API

		ScanRows(rows *sql.Rows, dest interface{}) error

		// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
		Transaction(fc func(API) error, opts ...*sql.TxOptions) (err error)

		// Begin begins a transaction
		Begin(opts ...*sql.TxOptions) API

		// Commit commit a transaction
		Commit() API

		// Rollback rollback a transaction
		Rollback() API

		SavePoint(name string) API

		RollbackTo(name string) API

		// Exec execute raw sql
		Exec(sql string, values ...interface{}) API
	}

	API interface {
		Config
		ChainableAPI
		FinisherAPI

		// Session create new db session
		Session(config *gorm.Session) API

		// WithContext change current instance db's context to ctx
		WithContext(ctx context.Context) API

		// Debug start debug mode
		Debug() API

		// Set store value with key into current db instance's context
		Set(key string, value interface{}) API

		// Get get value with key from current db instance's context
		Get(key string) (interface{}, bool)

		// InstanceSet store value with key into current db instance's context
		InstanceSet(key string, value interface{}) API

		// InstanceGet get value with key from current db instance's context
		InstanceGet(key string) (interface{}, bool)

		// AddError add error to db
		AddError(err error) error

		// DB returns `*sql.DB`
		DB() (*sql.DB, error)

		SetupJoinTable(model interface{}, field string, joinTable interface{}) error

		Use(plugin gorm.Plugin) (err error)
		Gorm() *gorm.DB
		Dialector() gorm.Dialector
		Migrator() gorm.Migrator

		Error() error
		RowsAffected() int64

		AutoMigrate(dst ...interface{}) error
		Association(column string) *gorm.Association
		Statement() *gorm.Statement

		Ping(ctx context.Context) error
	}
)
