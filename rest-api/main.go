package main

import (
	"context"
	"fmt"

	"github.com/TimiBolu/live-go-examples/rest-api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// init app
	err := initApp()
	if err != nil {
		panic(err)
	}

	defer database.CloseMongoDB()

	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		// write a todo to the database
		sampleDoc := bson.M{"name": "sample todo"}
		collection := database.GetCollection("todos")
		nDoc, err := collection.InsertOne(context.TODO(), sampleDoc)

		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Error inserting todo")
		}

		return c.JSON(nDoc)
	})

	app.Listen(":4005")
}

func initApp() error {
	err := loadENV()
	if err != nil {
		return err
	}

	err = database.StartMongoDB()
	if err != nil {
		return err
	}

	return nil
}

func loadENV() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
