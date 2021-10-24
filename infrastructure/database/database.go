package database

import (
	"database/sql"
	"errors"
	"fmt"
	"ozon-seller-api/internal/config"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// all errros from database
var (
	ErrDatabaseConfigIsNil      = errors.New("database configuration is nil")
	ErrNotUseDatabase           = errors.New("application not using any database")
	ErrNotConnectMysql          = errors.New("can't connect database mysql")
	ErrNotConnectPostgres       = errors.New("can't connect database postgres")
	ErrDatabaseDriverNotSupport = errors.New("driver database not support")
	ErrDialectNotSupport        = errors.New("dialect not support")
)

// Database is
type Database struct {
	Conn *sql.DB
}

// NewDatabase is
func NewDatabase(dc *config.DatabaseConfig) (*Database, error) {

	// TODO singletone ????
	// if db != nil {
	// 	return db, nil
	// }

	if dc == nil {
		return nil, ErrDatabaseConfigIsNil
	}

	if dc.Driver == "" {
		return nil, ErrNotUseDatabase
	}

	var err error
	var db = &Database{}

	switch dc.Driver {
	case "mysql":
		mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dc.Username, dc.Password, dc.Host, dc.Port, dc.Name)
		db.Conn, err = sql.Open("mysql", mysqlInfo)

		if err != nil {
			return nil, ErrNotConnectMysql
		}

	case "postgres":

		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dc.Host, dc.Port, dc.Username, dc.Password, dc.Name)
		db.Conn, err = sql.Open("postgres", psqlInfo)

		if err != nil {
			return nil, ErrNotConnectPostgres
		}

	default:
		return nil, ErrDatabaseDriverNotSupport
	}

	return db, nil
}
