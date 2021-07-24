package external

import (
	"github.com/ravielze/oculi/example/config"
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/persistent/sql"
	"github.com/ravielze/oculi/persistent/sql/postgre"
	sqlv2 "github.com/ravielze/oculi/persistent/sql/v2"
)

func NewPostgreSQL(config *config.Env, log logs.Logger) (sql.API, error) {
	api, err := sqlv2.NewClient(
		postgre.New(sql.ConnectionInfo{
			Address:  config.DatabaseAddress,
			Username: config.DatabaseUsername,
			Password: config.DatabasePassword,
			DbName:   config.DatabaseName,
		}), false,
		sql.WithMaxIdleConnection(config.DatabaseMaxIdleConnection),
		sql.WithMaxOpenConnection(config.DatabaseMaxOpenConnection),
		sql.WithConnMaxLifetime(config.DatabaseConnMaxLifetime),
		sql.WithLogMode(config.DatabaseLogMode),
		sql.WithLogger(log))
	if err != nil {
		return nil, err
	}

	api.AutoMigrate(dao.User{}, dao.Todo{})
	return api, nil
}
