package sql

import "fmt"

type ConnectionInfo struct {
	Address  string
	Username string
	Password string
	DbName   string
}

func (db ConnectionInfo) URI() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		db.Username, db.Password, db.Address, db.DbName)
}

func (db ConnectionInfo) PostgresURI() string {
	return fmt.Sprintf(`postgres://%s:%s@%s/%s?sslmode=disable`,
		db.Username, db.Password, db.Address, db.DbName)
}
