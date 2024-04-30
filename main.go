package main

import (
	"net/http"

	"github.com/poornatejav/Mongo_REST_API/controllers"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	r.GET("user/id", uc.getUser)
	r.POST("user", uc.createUser)
	r.DELETE("user/:id", uc.deleteUser)

	http.ListenAndServe("localhost:8080", r)

}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost:27107")

	if err != nil {
		panic(err)
	}

	return s
}
