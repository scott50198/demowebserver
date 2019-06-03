package dbhelper

import (
	"database/sql"
	"demowebserver/config"
	"demowebserver/model"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func OpenDB() error {
	db, err := sql.Open("mysql", config.SqlRootAccount+":"+config.SqlRootPassword+"@tcp(35.243.209.255)/"+config.SqlDBName)
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

func CheckAccountAndPasswordValidate(account string, password string) bool {
	row, err := DB.Query(CHECK_ACCOUNT_AND_PASSWORD_VALIDATE, account, password)
	defer row.Close()

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return row.Next()
}

func CheckAccountExist(accout string) bool {
	row, err := DB.Query(CHECK_ACCOUNT_EXIST, accout)
	defer row.Close()

	if err != nil {
		fmt.Println(err.Error())
		return true
	}

	return row.Next()
}

func CheckEmailExist(email string) bool {
	row, err := DB.Query(CHECK_EMAIL_EXIST, email)
	defer row.Close()

	if err != nil {
		fmt.Println(err.Error())
		return true
	}

	return row.Next()
}

func GetUserIdFromAccount(account string) (int, error) {
	row, err := DB.Query(GET_USER_ID_FROM_ACCOUNT, account)
	defer row.Close()

	if err != nil {
		fmt.Println(err.Error())
		return -1, mysqlErrorTranslater(err)
	}

	var id int
	row.Scan(&id)

	return id, nil
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
