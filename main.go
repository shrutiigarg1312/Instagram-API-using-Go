package main

import (
	"golang-api/config"
	"golang-api/controllers"
	"log"
	"net/http"
)

func main() {

	//connectong to mongodb database
	config.ConnectMongoDB()

	//Creating mux router
	mux := http.NewServeMux()

	//setting routes
	mux.HandleFunc("/", controllers.Welcome)
	mux.HandleFunc("/users", controllers.UsersHandler)
	mux.HandleFunc("/posts", controllers.PostsHandler)
	mux.HandleFunc("/posts/users", controllers.DisplayPosts)

	// set listen port
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
