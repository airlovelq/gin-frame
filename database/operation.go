package database

import (
	"fmt"
	"scoremanager/utils"

	//	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
