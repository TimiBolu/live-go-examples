package handlers

import (
	"context"
	"fmt"

	"github.com/TimiBolu/live-go-examples/rest-api/database"
	"github.com/TimiBolu/live-go-examples/rest-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type libraryDTO struct {
	Name    string   `json:"name" bson:"name"`
	Address string   `json:"address" bson:"address"`
	Empty   []string `json:"no_exists" bson:"books"`
}

// GET
func GetLibrarires(ctx *fiber.Ctx) error {
	libraryCollection := database.GetCollection("libraries")
	cursor, err := libraryCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return err
	}

	var libraries []models.Library
	if err := cursor.All(context.TODO(), &libraries); err != nil {
		return err
	}

	return ctx.JSON(libraries)
}

// POST
func CreateLibrary(ctx *fiber.Ctx) error {
	nLibrary := new(libraryDTO)

	if err := ctx.BodyParser(nLibrary); err != nil {
		return err
	}

	nLibrary.Empty = make([]string, 0)

	libraryCollection := database.GetCollection("libraries")
	nDoc, err := libraryCollection.InsertOne(context.TODO(), nLibrary)

	if err != nil {
		return err
	}

	fmt.Println(nLibrary)
	return ctx.JSON(fiber.Map{"id": nDoc.InsertedID})
}
