package dbhelper

import (
	"crypto/sha256"
	"database/sql"
	"demowebserver/config"
	"demowebserver/model"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func OpenDB() error {
	db, err := sql.Open("mysql", config.SqlRootAccount+":"+config.SqlRootPassword+"@tcp("+config.SqlDBHost+")/"+config.SqlDBName)
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func Register(user model.UserInfo) error {
	_, err := DB.Exec(
		INSERT_USERINFO_2,
		user.Account, PasswordEncrypt(user.Password), user.Name, user.Email,
	)
	if err != nil {
		return mysqlErrorTranslater(err)
	}
	return nil
}

func CheckAccountAndPasswordValidate(account string, password string) bool {
	row, err := DB.Query(CHECK_ACCOUNT_AND_PASSWORD_VALIDATE, account, PasswordEncrypt(password))
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

func GetUserInfoFromAccount(account string) (model.UserInfo, error) {
	row, err := DB.Query(GET_USER_INFO_FROM_ACCOUNT, account)
	defer row.Close()

	info := model.UserInfo{}
	if err != nil {
		fmt.Println(err.Error())
		return info, err
	}

	for row.Next() {
		row.Scan(&info.Id, &info.Account, &info.Name, &info.Email)
	}

	return info, nil
}

func mysqlErrorTranslater(err error) error {
	if mysqlError, ok := err.(*mysql.MySQLError); ok {

		switch mysqlError.Number {
		case 1062:
			return errors.New("帳號重複，請重新輸入")
		default:
			fmt.Println(err.Error())
			return errors.New("未知錯誤")
		}

	} else {
		fmt.Println(err.Error())
		return errors.New("未知錯誤")
	}

}

func PasswordEncrypt(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
