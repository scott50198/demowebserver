package main

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

func main() {
	// _, err := OpenDB()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("open db success")

}
