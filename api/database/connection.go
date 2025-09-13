package database

import (
	"fmt"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/leojimenezg/rtapi/api/envuse"
)

const (
	ENVFILE = "api.env"
	DATABASE = "MYSQL_DATABASE"
	USER = "MYSQL_USERNAME"
	PASSWORD = "MYSQL_PASSWORD"
	NETWORK = "MYSQL_NETWORK"
	ADDRESS = "MYSQL_ADDRESS"
)

type EnvFileError struct {
	File string
	Err error
}
func (e EnvFileError) Error() string {
	return fmt.Sprintf("failed to open env file %s:\n%v", e.File, e.Err)
}

type DatabaseError struct {
	Err error
}
func (e DatabaseError) Error() string {
	return fmt.Sprintf("failed to create database object:\n%v", e.Err)
}

type ConnectionError struct {
	Err error
}
func (e ConnectionError) Error() string {
	return fmt.Sprintf("failed to ping to database:\n%v", e.Err)
}

func ConnectToMySQLDB() (*sql.DB, error) {
	envErr := envuse.LoadEnvFile(ENVFILE)
	if envErr != nil { return nil, EnvFileError{ File: ENVFILE, Err: envErr } }
	driverConfig := mysql.NewConfig()
	driverConfig.User = envuse.GetEnv(USER)
	driverConfig.Passwd = envuse.GetEnv(PASSWORD)
	driverConfig.Net = envuse.GetEnv(NETWORK)
	driverConfig.Addr = envuse.GetEnv(ADDRESS)
	driverConfig.DBName = envuse.GetEnv(DATABASE)
	db, dbErr := sql.Open("mysql", driverConfig.FormatDSN())
	if dbErr != nil { return nil, DatabaseError{ Err: dbErr } }
	if pingErr := db.Ping(); pingErr != nil { 
		db.Close()
		return nil, ConnectionError{ Err: pingErr }
	}
	return db, nil
}
