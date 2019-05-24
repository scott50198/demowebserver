package dbhelper

import (
	"database/sql"
	"demowebserver/config"
	"demowebserver/model"
	"errors"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func OpenDB() error {
	db, err := sql.Open("mysql", config.SqlRootAccount+":"+config.SqlRootPassword+"@/"+config.SqlDBName)
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func Register(user model.UserInfo) error {
	_, err := DB.Exec(
		INSERT_USERINFO,
		user.Account, user.Password, user.Name, user.Email,
	)

	if err != nil {
		return mysqlErrorTranslater(err)
	}

	return nil

}

func mysqlErrorTranslater(err error) error {
	if mysqlError, ok := err.(*mysql.MySQLError); ok {

		switch mysqlError.Number {
		case 1062:
			return errors.New("帳號重複，請重新輸入")
		default:
			return errors.New("未知錯誤")
		}

	} else {
		return errors.New("未知錯誤")
	}

}
