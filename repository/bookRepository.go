package repository

import (
	"github.com/rafrdz/go_api/db"
	"github.com/rafrdz/go_api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type BookRepository interface {
	SaveNewBook(book *model.Book) (string, error)
}

type bookRepository struct {
	database string
	collection string
}

func NewBookRepository(database string, collection string) BookRepository {
	return &bookRepository{
		database: database,
		collection: collection,
	}
}

func (repo *bookRepository) SaveNewBook(book *model.Book) (string, error) {
	client, ctx, cancel = db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	result, err := client.Database(repo.database).Collection(repo.collection).InsertOne(ctx, book)
	if err != nil {
		log.Printf("Could not create Book: %v", err)
		return "", err
	}
	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	return insertedID, nil
}