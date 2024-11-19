package main

import (
	"fmt"
	"github.com/KostyaBagr/HomeServer_webPanel/initializers"  
	"github.com/KostyaBagr/HomeServer_webPanel/models"       
	"log"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}
	fmt.Println("Migrations completed successfully!")
}
