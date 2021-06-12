package sqlv2

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// GORM perform single create, update, delete operations in transactions by default to ensure database data integrity
// You can disable it by setting `SkipDefaultTransaction` to true
func (i *Impl) SkipDefaultTransaction() bool {
	return i.Database.Config.SkipDefaultTransaction
}

// NamingStrategy tables, columns naming strategy
func (i *Impl) NamingStrategy() schema.Namer {
	return i.Database.Config.NamingStrategy
}

// FullSaveAssociations full save associations
func (i *Impl) FullSaveAssociations() bool {
	return i.Database.Config.FullSaveAssociations
}

// Logger
func (i *Impl) Logger() logger.Interface {
	return i.Database.Config.Logger
}

// NowFunc the function to be used when creating a new timestamp
func (i *Impl) NowFunc() func() time.Time {
	return i.Database.Config.NowFunc
}

// DryRun generate sql without execute
func (i *Impl) DryRun() bool {
	return i.Database.Config.DryRun
}

// PrepareStmt executes the given query in cached statement
func (i *Impl) PrepareStmt() bool {
	return i.Database.Config.PrepareStmt
}

// DisableAutomaticPing
func (i *Impl) DisableAutomaticPing() bool {
	return i.Database.Config.DisableAutomaticPing
}

// DisableForeignKeyConstraintWhenMigrating
func (i *Impl) DisableForeignKeyConstraintWhenMigrating() bool {
	return i.Database.Config.DisableForeignKeyConstraintWhenMigrating
}

// AllowGlobalUpdate allow global update
func (i *Impl) AllowGlobalUpdate() bool {
	return i.Database.Config.AllowGlobalUpdate
}

// ClauseBuilders clause builder
func (i *Impl) ClauseBuilders() map[string]clause.ClauseBuilder {
	return i.Database.Config.ClauseBuilders
}

// ConnPool db conn pool
func (i *Impl) ConnPool() gorm.ConnPool {
	return i.Database.Config.ConnPool
}

// Plugins registered plugins
func (i *Impl) Plugins() map[string]gorm.Plugin {
	return i.Database.Config.Plugins
}
