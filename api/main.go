package main

import (
	"fmt"
	"github.com/leojimenezg/rtapi/api/database"
	"github.com/leojimenezg/rtapi/api/services"
)

func main() {
	db, dbErr := database.ConnectToMySQLDB()
	if dbErr != nil { panic(dbErr) }
	defer db.Close()
	
	// Test Category
	category, err := services.GetCategoryById(db, 1)
	if err != nil { panic(err) }
	fmt.Println("Category:", category)
	
	// Test Difficulty
	difficulty, err := services.GetDifficultyById(db, 1)
	if err != nil { panic(err) }
	fmt.Println("Difficulty:", difficulty)
	
	// Test ResourceType
	resourceType, err := services.GetResourceTypeById(db, 1)
	if err != nil { panic(err) }
	fmt.Println("ResourceType:", resourceType)
	
	// Test Resource
	resource, err := services.GetResourceById(db, 1)
	if err != nil { panic(err) }
	fmt.Println("Resource:", resource)
	
	// Test Topic
	topic, err := services.GetTopicById(db, 1)
	if err != nil { panic(err) }
	fmt.Println("Topic:", topic)
}
