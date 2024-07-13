package main

import (
	"log"
	"mvc_example/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
