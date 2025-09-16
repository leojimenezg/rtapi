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
	defer db.Close()
	fmt.Println("Connected to Database!")

	handler := handlers.New(db)
	router := gin.Default()

	// Categories endpoints
	router.GET("/categories", handler.GetAllCategories)
	router.GET("/categories/:id", handler.GetCategoryById)
	router.GET("/categories/name/:name", handler.GetCategoryByName)

	// Difficulties endpoints  
	router.GET("/difficulties", handler.GetAllDifficulties)
	router.GET("/difficulties/:id", handler.GetDifficultyById)
	router.GET("/difficulties/name/:name", handler.GetDifficultyByName)

	// Resource_types endpoints
	router.GET("/resource_types", handler.GetAllResourceTypes)
	router.GET("/resource_types/:id", handler.GetResourceTypeById)
	router.GET("/resource_types/name/:name", handler.GetResourceTypeByName)

	// Topics endpoints
	router.GET("/topics", handler.GetAllTopics)
	router.GET("/topics/:id", handler.GetTopicById)
	router.GET("/topics/name/:name", handler.GetTopicByName)
	router.GET("/topics/details", handler.GetAllTopicsWithDetails)
	router.GET("/topics/details/:id", handler.GetTopicWithDetailsById)
	router.GET("/topics/details/name/:name", handler.GetTopicWithDetailsByName)

	// Random topic using query params
	router.GET("/topics/details/random", handler.GetRandomTopicWithDetails)

	router.Run()
}
