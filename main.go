package main

import (
	"library/internal/database"
	"library/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.Migrate(database.DB)

	// Set up Gin router
	router := gin.Default()
	bookHandler := handlers.NewBookHandler(database.DB)
	borrowHandler := handlers.NewBorrowingHandler(database.DB)

	// Define API routes for books
	router.GET("/api/books", bookHandler.GetAllBooks)
	router.GET("/api/books/:id", bookHandler.GetBookByID)
	router.POST("/api/books", bookHandler.CreateBook)
	router.PUT("/api/books/:id", bookHandler.UpdateBook)
	router.DELETE("/api/books/:id", bookHandler.DeleteBook)

	// Define API routes for borrowings
	router.POST("/api/books/:id/borrow", borrowHandler.BorrowBook)
	router.POST("/api/books/:id/return", borrowHandler.ReturnBook)
	router.GET("/api/borrowings", borrowHandler.ListActiveBorrowings)

	// Start the server
	router.Run(":8080")
}
