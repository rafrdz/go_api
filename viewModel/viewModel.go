package viewmodel

type NewBook struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
