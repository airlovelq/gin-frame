package database

import (
	"fmt"
	"scoremanager/utils"

	//	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v3"
)

type DatabaseOp struct {
	tx *sqlx.Tx
}

var db *sqlx.DB
var connectErr error

func init() {
	host := utils.GetEnvDefault("DB_HOST", "192.168.100.103")
	port := utils.GetEnvDefault("DB_PORT", "10202")
	dbName := utils.GetEnvDefault("DB_NAME", "sugar2")
	password := utils.GetEnvDefault("DB_PASSWORD", "postgres")
	user := utils.GetEnvDefault("DB_USER", "postgres")
	dbtype := utils.GetEnvDefault("DB_TYPE", "postgresql")
	var strConnect string
	switch dbtype {
	case "postgresql":
		strConnect := fmt.Sprint(dbtype, "://", user, ":", password, "@", host, ":", port, "/", dbName, "?sslmode=disable")
		db, connectErr = sqlx.Connect("postgres", strConnect)
	case "mysql":
		strConnect := fmt.Sprint(dbtype, "://", user, ":", password, "@tcp(", host, ":", port, ")/", dbName)
		db, connectErr = sqlx.Connect("mysql", strConnect)
	}

	fmt.Print(strConnect)
	if connectErr != nil {
		panic(connectErr)
	}
	//var err error

}

func NewDatabaseOp() *DatabaseOp {
	dbOp := DatabaseOp{}
	return &dbOp
}

func (dbOp *DatabaseOp) Begin() error {
	if connectErr != nil {
		return connectErr
	}
	var err error
	dbOp.tx, err = db.Beginx()
	return err
}

func (dbOp *DatabaseOp) Rollback() error {
	if dbOp.tx != nil {
		return dbOp.tx.Rollback()
	}
	return nil
}

func (dbOp *DatabaseOp) Commit() error {
	if dbOp.tx != nil {
		return dbOp.tx.Commit()
	}
	return nil
}

func (dbOp *DatabaseOp) GetUserByEmail(email string) (User, error) {
	sqlresult := dbOp.tx.QueryRowx("SELECT * FROM user_platform WHERE email=$1", email)
	var u User
	err := sqlresult.StructScan(&u)
	return u, err
}

func (dbOp *DatabaseOp) GetUserByPhone(phone string) (User, error) {
	sqlresult := dbOp.tx.QueryRowx("SELECT * FROM user_platform WHERE phone=$1", phone)
	var u User
	err := sqlresult.StructScan(&u)
	return u, err
}

func (dbOp *DatabaseOp) GetUserByUserName(userName string) (User, error) {
	sqlresult := dbOp.tx.QueryRowx("SELECT * FROM user_platform WHERE user_name=$1", userName)
	var u User
	err := sqlresult.StructScan(&u)
	return u, err
}

func (dbOp *DatabaseOp) GetUserByUserID(userID string) (User, error) {
	sqlresult := dbOp.tx.QueryRowx("SELECT * FROM user_platform WHERE id=$1", userID)
	var u User
	err := sqlresult.StructScan(&u)
	return u, err
}

func (dbOp *DatabaseOp) CreateUserByEmail(email string, passwordHash string, userType int) (string, error) {
	userID := uuid.NewV4()
	_, err := dbOp.tx.Exec("INSERT INTO user_platform (id, email, password_hash, user_type, banned, create_date, operate_date, operator) values ($1,$2,$3,$4,0,now(),now(),$5)", userID, email, passwordHash, userType, userID)
	if err != nil {
		return "", err
	}
	return userID.String(), nil
}

func (dbOp *DatabaseOp) ResetPasswordByEmail(email string, passwordHash string) error {
	_, err := dbOp.tx.Exec("UPDATE user_platform SET password_hash=$1, operate_date=now() WHERE email=$2", passwordHash, email)
	return err
}

func (dbOp *DatabaseOp) ResetPasswordByID(userID string, passwordHash string) error {
	_, err := dbOp.tx.Exec("UPDATE user_platform SET password_hash=$1, operator=$2, operate_date=now() WHERE id=$3", passwordHash, userID, userID)
	return err
}

func (dbOp *DatabaseOp) ResetEmailByID(userID string, email string) error {
	_, err := dbOp.tx.Exec("UPDATE user_platform SET email=$1, operator=$2, operate_date=now() WHERE id=$3", email, userID, userID)
	return err
}

func (dbOp *DatabaseOp) UpdateUserInfo(userID string, sex null.Int, age null.Int, userName null.String, name null.String, info null.String) error {
	sqlstr := "UPDATE user_platform SET "
	count := 0
	values := make([]interface{}, 0)
	if sex.Valid {
		count++
		sqlstr += fmt.Sprintf("sex=$%d, ", count)
		values = append(values, sex)
	}
	if age.Valid {
		count++
		sqlstr += fmt.Sprintf("age=$%d, ", count)
		values = append(values, age)
	}
	if userName.Valid {
		count++
		sqlstr += fmt.Sprintf("user_name=$%d, ", count)
		values = append(values, userName)
	}
	if name.Valid {
		count++
		sqlstr += fmt.Sprintf("name=$%d, ", count)
		values = append(values, name)
	}
	if info.Valid {
		count++
		sqlstr += fmt.Sprintf("info=$%d, ", count)
		values = append(values, info)
	}
	sqlstr += fmt.Sprintf("operator=$%d WHERE id=$%d", count+1, count+2)
	values = append(values, userID)
	values = append(values, userID)
	_, err := dbOp.tx.Exec(sqlstr, values...)
	return err
}
