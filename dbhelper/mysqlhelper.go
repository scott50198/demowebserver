package dbhelper

import (
	"database/sql"
	"demowebserver/config"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.SqlRootAccount+":"+config.SqlRootPassword+"@/"+config.SqlDBName)
	if err != nil {
		return nil, err
	}
	return db, nil
}
