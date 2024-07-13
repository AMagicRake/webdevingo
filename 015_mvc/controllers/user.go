package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"mvc_example/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
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
	u.Id = "007"

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Write code to delete user\n")
}
