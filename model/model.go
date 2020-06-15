package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID `bson:"_id, omitempty"`
	Title  string `json:"title",bson:"title"`
	Author string `json:"author",bson:"author"`
}
