package integration

import (
	"os"
	"testing"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/leojimenezg/rtapi/api/envuse"
)

const (
	envfile = "integration.env"
	dtabase = "MYSQL_DATABASE"
	user = "MYSQL_USERNAME"
	password = "MYSQL_PASSWORD"
	network = "MYSQL_NETWORK"
	address = "MYSQL_ADDRESS"
)

var TestingDB *sql.DB

func connectToMySQLDB() *sql.DB {
	envErr := envuse.LoadEnvFile(envfile)
	if envErr != nil { panic("failed to open test env file") }
	driverConfig := mysql.NewConfig()
	driverConfig.User = envuse.GetEnv(user)
	driverConfig.Passwd = envuse.GetEnv(password)
	driverConfig.Net = envuse.GetEnv(network)
	driverConfig.Addr = envuse.GetEnv(address)
	driverConfig.DBName = envuse.GetEnv(dtabase)
	db, dbErr := sql.Open("mysql", driverConfig.FormatDSN())
	if dbErr != nil { panic("failed to open database connection") }
	if db.Ping() != nil {
		db.Close()
		panic("failed to ping to database")
	}
	return db
}

func TestMain(m *testing.M) {
	TestingDB = connectToMySQLDB()
	if TestingDB == nil { panic("could not connect to test database") }
	exitCode := m.Run()
	TestingDB.Close()
	os.Exit(exitCode)
}
