package main

import (
	"bookstore/controllers"
	"bookstore/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to the database
	database.ConnectDatabase()

	// Define routes for the book resource
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBook)
	r.POST("/books", controllers.CreateBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// Start the server
	r.Run()
}
