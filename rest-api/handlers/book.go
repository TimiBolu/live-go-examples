package handlers

import (
	"context"

	"github.com/TimiBolu/live-go-examples/rest-api/database"
	"github.com/TimiBolu/live-go-examples/rest-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type bookDTO struct {
	Title     string `json:"title" bson:"title"`
	Author    string `json:"author" bson:"author"`
	ISBN      string `json:"isbn" bson:"isbn"`
	LibraryID string `json:"libraryID" bson:"libraryID"`
}

func CreateBook(ctx *fiber.Ctx) error {
	nBook := new(bookDTO)

	if err := ctx.BodyParser(nBook); err != nil {
		return err
	}

	libraryCollection := database.GetCollection("libraries")
	id, _ := primitive.ObjectIDFromHex(nBook.LibraryID)
	filter := bson.M{"_id": id}
	nBookData := models.Book{
		Title:  nBook.Title,
		Author: nBook.Author,
		ISBN:   nBook.ISBN,
	}
	update := bson.M{"$push": bson.M{"books": nBookData}}

	_, err := libraryCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return nil
	}

	return ctx.SendString("Book created successfully")
}
