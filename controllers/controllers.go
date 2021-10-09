package controllers

import (
	"context"
	"fmt"
	"golang-api/config"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//Handling '/' route
func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Welcome! This is Instagram Backend API made using Golang.</h1>")
}

//Handling '/users' route
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	usersCollection := config.Mongo.DB.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "The id query parameter is missing!", http.StatusBadRequest)
			return
		}
		filterCursor, err := usersCollection.Find(ctx, bson.M{"_id": id})
		if err != nil {
			log.Fatal(err)
		}
		var userFiltered []bson.M
		if err = filterCursor.All(ctx, &userFiltered); err != nil {
			log.Fatal(err)
		}
		fmt.Print(userFiltered)
	}
	if r.Method == "POST" {
		userResult, err := usersCollection.InsertOne(ctx, bson.D{
			{Key: "name", Value: "Shruti Garg"},
			{Key: "email", Value: "shruti.garg2019@vitstudent.ac.in"},
			{Key: "password", Value: "shruti0994"},
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Created a user!\n")
		fmt.Println(userResult)
	}
}

//Handling '/posts' route
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	postsCollection := config.Mongo.DB.Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//for GET request
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "The id query parameter is missing!", http.StatusBadRequest)
			return
		}
		filterCursor, err := postsCollection.Find(ctx, bson.M{"_id": id})
		if err != nil {
			log.Fatal(err)
		}
		var postFiltered []bson.M
		if err = filterCursor.All(ctx, &postFiltered); err != nil {
			log.Fatal(err)
		}
		fmt.Print(postFiltered)
	}

	//for POST request
	if r.Method == "POST" {
		postResult, err := postsCollection.InsertOne(ctx, bson.D{
			{Key: "image", Value: "https://www.bloorresearch.com/wp-content/uploads/2013/03/MONGO-DB-logo-300x470--x.png"},
			{Key: "caption", Value: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."},
			{Key: "timestamp", Value: time.Now().String()},
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Created a post!\n")
		fmt.Println(postResult)
	}
}

//Handling '/posts/users/' route
func DisplayPosts(w http.ResponseWriter, r *http.Request) {
	postsCollection := config.Mongo.DB.Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//for GET request
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "The id query parameter is missing!", http.StatusBadRequest)
			return
		}
		filterCursor, err := postsCollection.Find(ctx, bson.M{"_id": id})
		if err != nil {
			log.Fatal(err)
		}
		var postFiltered []bson.M
		if err = filterCursor.All(ctx, &postFiltered); err != nil {
			log.Fatal(err)
		}
		fmt.Print(postFiltered)
	}
}
