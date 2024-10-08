package db

import (
	"database/sql"

	"new-token/pkg/log"

	_ "github.com/go-sql-driver/mysql"
)

// MySQL Configuration Struct
type mysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// MySQL Configuration Variable
var mysqlCfg mysqlConfig

// MySQL Variable
var MySQL *sql.DB

// MySQL Connect Function
func mysqlConnect() *sql.DB {
	// Initialize Connection
	conn, err := sql.Open("mysql", mysqlCfg.User+":"+mysqlCfg.Password+"@tcp("+mysqlCfg.Host+":"+mysqlCfg.Port+")/"+mysqlCfg.Name)
	if err != nil {
		log.Println(log.LogLevelFatal, "mysql-connect", err.Error())
	}

	// Test Connection
	err = conn.Ping()
	if err != nil {
		log.Println(log.LogLevelFatal, "mysql-connect", err.Error())
	}

	// Return Connection
	return conn
}
