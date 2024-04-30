// package main

// import (
// 	"context"
// 	"net/http"

// 	"github.com/julienschmidt/httprouter"
// 	"github.com/poornatejav/Mongo_REST_API/controllers"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func main() {

// 	r := httprouter.New()
// 	uc := controllers.NewUserController(getSession())

// 	r.GET("user/id", uc.GetUser)
// 	r.POST("user", uc.CreateUser)
// 	r.DELETE("user/:id", uc.DeleteUser)

// 	http.ListenAndServe("localhost:8080", r)

// }

// func getSession() *mongo.Client {
// 	clientOptions := options.Client().ApplyURI("mongodb+srv://poornateja340:MYd2chTy972QtSCm@mycluster.8a6xt10.mongodb.net/?retryWrites=true&w=majority")
// 	client, err := mongo.Connect(context.Background(), clientOptions)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return client
// }

package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/poornatejav/Mongo_REST_API/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	r := httprouter.New()
	client := getSession()
	uc := controllers.NewUserController(client)

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8080", r)

}

func getSession() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://poornateja340:MYd2chTy972QtSCm@mycluster.8a6xt10.mongodb.net/API_test?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	return client
}
