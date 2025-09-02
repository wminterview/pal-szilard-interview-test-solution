package handlers

type CreateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	ISBN   string `json:"isbn" binding:"required"`
	Year   int    `json:"year" binding:"required,min=1900,max=2024"`
}
