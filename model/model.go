package model

type Book struct {
	Title  string `json:"title",bson:"title"`
	Author string `json:"author",bson:"author"`
}
