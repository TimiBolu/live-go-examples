package main

import (
	"os"

	"github.com/TimiBolu/live-go-examples/rest-api/database"
	"github.com/joho/godotenv"
)

func main() {
	// init app
	err := initApp()
	if err != nil {
		panic(err)
	}

	defer database.CloseMongoDB()

	app := generateApp()

	// app.Post("/", func(c *fiber.Ctx) error {
	// 	// write a todo to the database
	// 	sampleDoc := bson.M{"name": "sample todo"}
	// 	collection := database.GetCollection("todos")
	// 	nDoc, err := collection.InsertOne(context.TODO(), sampleDoc)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return c.Status(fiber.StatusInternalServerError).SendString("Error inserting todo")
	// 	}

	// 	return c.JSON(nDoc)
	// })

	// get the port from the env
	port := os.Getenv("PORT")
	app.Listen(":" + port)
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
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
