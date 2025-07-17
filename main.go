package main

import (
	"context"
	"log"
	"task_management/db"
	"task_management/router"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func main() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://tsigemariamzewdu20:Bbirhan2121@taskgo.men60cg.mongodb.net/")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	db.TaskCollection = client.Database("taskmanager").Collection("tasks")
	app := gin.Default()
	router.SetUpRoutes(app)
	app.Run("localhost:8081")

}
