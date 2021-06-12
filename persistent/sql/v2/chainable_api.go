package sqlv2

import (
	sqlOculi "github.com/ravielze/oculi/persistent/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (i *Impl) Model(value interface{}) sqlOculi.API {
	db := i.Database.Model(value)
	return i.copy(db)
}

// Clauses Add clauses
func (i *Impl) Clauses(conds ...clause.Expression) sqlOculi.API {
	db := i.Database.Clauses(conds...)
	return i.copy(db)
}

// Table specify the table you would like to run db operations
func (i *Impl) Table(name string, args ...interface{}) sqlOculi.API {
	db := i.Database.Table(name, args...)
	return i.copy(db)
}

// Distinct specify distinct fields that you want querying
func (i *Impl) Distinct(args ...interface{}) sqlOculi.API {
	db := i.Database.Distinct(args...)
	return i.copy(db)
}

// Select specify fields that you want when querying, creating, updating
func (i *Impl) Select(query interface{}, args ...interface{}) sqlOculi.API {
	db := i.Database.Select(query, args...)
	return i.copy(db)
}

// Omit specify fields that you want to ignore when creating, updating and querying
func (i *Impl) Omit(columns ...string) sqlOculi.API {
	db := i.Database.Omit(columns...)
	return i.copy(db)
}

// Where add conditions
func (i *Impl) Where(query interface{}, args ...interface{}) sqlOculi.API {
	db := i.Database.Where(query, args...)
	return i.copy(db)
}

// Not add NOT conditions
func (i *Impl) Not(query interface{}, args ...interface{}) sqlOculi.API {
	db := i.Database.Not(query, args...)
	return i.copy(db)
}

// Or add OR conditions
func (i *Impl) Or(query interface{}, args ...interface{}) sqlOculi.API {
	db := i.Database.Or(query, args...)
	return i.copy(db)
}

// Joins specify Joins conditions
//     db.Joins("Account").Find(&user)
//     db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
func (i *Impl) Joins(query string, args ...interface{}) sqlOculi.API {
	db := i.Database.Joins(query, args...)
	return i.copy(db)
}

// Group specify the group method on the find
func (i *Impl) Group(name string) sqlOculi.API {
	db := i.Database.Group(name)
	return i.copy(db)
}

// Having specify HAVING conditions for GROUP BY
func (i *Impl) Having(query interface{}, args ...interface{}) sqlOculi.API {
	db := i.Database.Having(query, args...)
	return i.copy(db)
}

// Order specify order when retrieve records from database
//     db.Order("name DESC")
//     db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}, Desc: true})
func (i *Impl) Order(value interface{}) sqlOculi.API {
	db := i.Database.Order(value)
	return i.copy(db)
}

// Limit specify the number of records to be retrieved
func (i *Impl) Limit(limit int) sqlOculi.API {
	db := i.Database.Limit(limit)
	return i.copy(db)
}

// Offset specify the number of records to skip before starting to return the records
func (i *Impl) Offset(offset int) sqlOculi.API {
	db := i.Database.Offset(offset)
	return i.copy(db)
}

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
func (i *Impl) Scopes(funcs ...func(db sqlOculi.API) sqlOculi.API) sqlOculi.API {
	lfuncs := len(funcs)
	nfuncs := make([]func(db *gorm.DB) *gorm.DB, lfuncs)
	for idx := 0; idx < lfuncs; idx++ {
		nfuncs[idx] = func(db *gorm.DB) *gorm.DB {
			return funcs[idx](i).Gorm()
		}
	}

	db := i.Database.Scopes(nfuncs...)
	return i.copy(db)
}

// Preload preload associations with given conditions
//    db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
func (i *Impl) Preload(query string, args ...interface{}) sqlOculi.API {
	db := i.Database.Preload(query, args...)
	return i.copy(db)
}

func (i *Impl) Attrs(attrs ...interface{}) sqlOculi.API {
	db := i.Database.Attrs(attrs)
	return i.copy(db)
}

func (i *Impl) Assign(attrs ...interface{}) sqlOculi.API {
	db := i.Database.Assign(attrs...)
	return i.copy(db)
}

func (i *Impl) Unscoped() sqlOculi.API {
	db := i.Database.Unscoped()
	return i.copy(db)
}

func (i *Impl) Raw(sql string, values ...interface{}) sqlOculi.API {
	db := i.Database.Raw(sql, values...)
	return i.copy(db)
}
