package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mvc_example/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type UserController struct {
	session *mongo.Collection
}

func NewUserController(s *mongo.Collection) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "Male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}

	log.Fatal(json.NewDecoder(r.Body).Decode(&u))

	//should probably generate uuid but examples
	err := uc.session.Database().Client().Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatalln()
	}

	_, err = uc.session.InsertOne(context.Background(), u)

	if err != nil {
		log.Fatalln(err)
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Write code to delete user\n")
}
