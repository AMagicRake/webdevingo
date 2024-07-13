package main

import (
	"context"
	"log"
	"mvc_example/controllers"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func getSession() *mongo.Collection {
	connString := ""
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	//Connect to local host mongo
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		log.Fatalln(err)
	}

	s := client.Database("learning").Collection("golangExperimenting")
	return s
}
