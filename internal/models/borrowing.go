package models

import "time"

type Borrowing struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	BookID       uint       `json:"book_id"`
	Book         Book       `json:"book" gorm:"foreignKey:BookID"`
	BorrowerName string     `json:"borrower_name" gorm:"not null"`
	BorrowDate   time.Time  `json:"borrow_date"`
	ReturnDate   *time.Time `json:"return_date"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
