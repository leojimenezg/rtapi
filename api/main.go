package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leojimenezg/rtapi/api/database"
	"github.com/leojimenezg/rtapi/api/handlers"
)

func main() {
	db, dbErr := database.ConnectToMySQLDB()
	if dbErr != nil { panic(dbErr) }
	fmt.Println("Connected to Database!")
	defer db.Close()
	handler := handlers.New(db)
	router := gin.Default()
	router.GET("/topics", handler.GetAllTopics)
	router.GET("/topics/details", handler.GetAllTopicsWithDetails)
	router.Run()
}
