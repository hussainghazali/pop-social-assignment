package main

import (
	"fmt"
	"log"

	"pop-social-assignment/initializers"
	"pop-social-assignment/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Like{})
	fmt.Println("👍 Migration complete")
}