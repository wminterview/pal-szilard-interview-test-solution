package handlers

type BorrowRequest struct {
	BorrowerName string `json:"borrower_name" binding:"required"`
}
