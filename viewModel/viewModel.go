package viewmodel

type NewBook struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type Book struct {
	ID     string `json:"id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
