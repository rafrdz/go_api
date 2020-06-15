package repository

import (
	"context"
	"github.com/rafrdz/go_api/db"
	"github.com/rafrdz/go_api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type BookRepository interface {
	SaveNewBook(book *model.Book) (string, error)
	GetAllBooks() ([]model.Book, error)
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
	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	result, err := client.Database(repo.database).Collection(repo.collection).InsertOne(ctx, book)
	if err != nil {
		log.Print(err)
		return "", err
	}
	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	return insertedID, nil
}

func (repo *bookRepository) GetAllBooks() ([]model.Book, error) {
	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	var allBooks []model.Book
	result, err := client.Database(repo.database).Collection(repo.collection).Find(ctx, bson.D{})
	if err != nil {
		log.Print(err)
		return nil, err
	}
	for result.Next(context.TODO()) {
		var book model.Book
		err := result.Decode(&book)
		if err != nil {
			log.Print(err)
		}
		allBooks = append(allBooks, book)
	}
	result.Close(context.TODO())
	return allBooks, nil
}