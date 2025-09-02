package handlers

import (
	"library/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BorrowingHandler struct {
	DB *gorm.DB
}

func NewBorrowingHandler(db *gorm.DB) *BorrowingHandler {
	return &BorrowingHandler{DB: db}
}

func (h *BorrowingHandler) BorrowBook(c *gin.Context) {
	var req BorrowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Bad Request", Message: err.Error()})
		return
	}

	bookID := c.Param("id")
	var book models.Book
	if err := h.DB.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Not Found", Message: "Book not found"})
		return
	}

	if !book.Available {
		c.JSON(http.StatusConflict, ErrorResponse{Error: "Conflict", Message: "Book is not available"})
		return
	}

	borrowing := models.Borrowing{
		BookID:       book.ID,
		BorrowerName: req.BorrowerName,
		BorrowDate:   time.Now(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := h.DB.Create(&borrowing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Internal Server Error", Message: err.Error()})
		return
	}

	book.Available = false
	h.DB.Save(&book)

	c.JSON(http.StatusCreated, borrowing)
}

func (h *BorrowingHandler) ReturnBook(c *gin.Context) {
	bookID := c.Param("id")
	var borrowing models.Borrowing
	if err := h.DB.Where("book_id = ? AND return_date IS NULL", bookID).First(&borrowing).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Not Found", Message: "Borrowing record not found"})
		return
	}

	returnDate := time.Now()
	borrowing.ReturnDate = &returnDate
	borrowing.UpdatedAt = time.Now()

	if err := h.DB.Save(&borrowing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Internal Server Error", Message: err.Error()})
		return
	}

	var book models.Book
	h.DB.First(&book, borrowing.BookID)
	book.Available = true
	h.DB.Save(&book)

	c.JSON(http.StatusOK, borrowing)
}

func (h *BorrowingHandler) ListActiveBorrowings(c *gin.Context) {
	var borrowings []models.Borrowing
	if err := h.DB.Where("return_date IS NULL").Find(&borrowings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Internal Server Error", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, borrowings)
}
