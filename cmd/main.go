package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/gnanasuryateja-kurapati/simple-rest-api-golang/datastore"
	"github.com/gnanasuryateja-kurapati/simple-rest-api-golang/handlers"
)

// execution of program starts here
func main() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		os.Exit(1) // if any error, exit the program
	}
	// load sample data
	datastore.LoadSampleUsers()
	// define the router
	router := gin.Default()
	// map the handlers
	router.GET("/users", handlers.GetUsers)
	router.GET("/user/:id", handlers.GetUserById)
	router.POST("/user", handlers.CreateUser)
	router.PUT("/user", handlers.UpdateUser)
	router.DELETE("/user/:id", handlers.DeleteUser)
	// run the server
	router.Run(os.Getenv("PORT"))
}
