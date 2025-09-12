package main

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/go-sql-driver/mysql"
	"github.com/leojimenezg/rtapi/api/envuse"
)

type Topic struct {
	id int
	name string
	description string
	count int
	category_id int
	difficulty_id int
	is_active bool
	created_at string
	updated_at string
}

const (
	ALL_TOPICS = `SELECT * FROM topics;`
	ONE_TOPIC = `SELECT * FROM topics WHERE id = 1;`
)

func main() {
	// Load env file and get variables
	envFile := os.Args[1]
	errFile := envuse.LoadEnvFile(envFile)
	if errFile != nil { panic(errFile) }
	// Configure database connection
	connectionConfig := mysql.NewConfig()
	connectionConfig.User = envuse.GetEnv("MYSQL_USERNAME")
	connectionConfig.Passwd = envuse.GetEnv("MYSQL_PASSWORD")
	connectionConfig.Net = envuse.GetEnv("MYSQL_NETWORK")
	connectionConfig.Addr = envuse.GetEnv("MYSQL_ADDRESS")
	connectionConfig.DBName = envuse.GetEnv("MYSQL_DATABASE")
	database, errDB := sql.Open("mysql", connectionConfig.FormatDSN())
	if errDB != nil { panic(errDB) }
	defer database.Close()
	if ping := database.Ping(); ping != nil { panic(ping) }
	fmt.Println("Connected to database!")
	// Query to database with multiple rows results
	rows, errRows := database.Query(ALL_TOPICS)
	if errRows != nil { panic(errRows) }
	defer rows.Close()
	var topics []Topic
	for rows.Next() {
		var topic Topic
		errQuery := rows.Scan(
			&topic.id, &topic.name, &topic.description, &topic.count, &topic.category_id,
			&topic.difficulty_id, &topic.is_active, &topic.created_at, &topic.updated_at,
		)
		if errQuery != nil { panic(errQuery) }
		topics = append(topics, topic)
	}
	if err := rows.Err(); err != nil { panic(err) }
	for _, t := range topics {
		fmt.Println(t)
	}
	// Query to database with single row results
	row := database.QueryRow(ONE_TOPIC)
	var topic Topic
	errRow := row.Scan(
		&topic.id, &topic.name, &topic.description, &topic.count, &topic.category_id,
		&topic.difficulty_id, &topic.is_active, &topic.created_at, &topic.updated_at,
	)
	if errRow != nil { panic(errRow) }
	fmt.Println(topic)
}
