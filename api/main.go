package main

import (
	"fmt"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

// Temporary local credentials
const (
	DATABASE = ""
	USERNAME = ""
	PASSWORD = ""
	ADDRESS = ""
)

func main() {
	connectionConfig := mysql.NewConfig()
	connectionConfig.User = USERNAME
	connectionConfig.Passwd = PASSWORD
	connectionConfig.Net = "tcp"
	connectionConfig.Addr = ADDRESS
	connectionConfig.DBName = DATABASE
	db, err := sql.Open("mysql", connectionConfig.FormatDSN())
	if err != nil { panic(err) }
	if ping := db.Ping(); ping != nil { panic(ping) }
	fmt.Printf("Connected to %s database!\n", DATABASE)
	db.Close()
}
