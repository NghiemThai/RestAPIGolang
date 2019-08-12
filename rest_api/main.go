package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	ID      string `json:id`
	Title   string `json:title`
	Content string `json:content`
}

var posts []Post

func main() {
	router := mux.NewRouter()

	posts = append(posts,
		Post{ID: "1", Title: "Trying Clean Architecture on Golang", Content: "Clean Architecture on Golang"},
		Post{ID: "2", Title: "Why I Love Golang", Content: "I love the Go programming language, or as some refer to it, Golang. It’s simple and it’s great."},
		Post{ID: "3", Title: "Build RESTful API service in golang using gin-gonic framework", Content: "gin-gonic framework"},
		Post{ID: "4", Title: "Clean Architecture using Golang", Content: "Clean Architecture using Golang"},
		Post{ID: "5", Title: "Web Service Architecture for Golang Developers", Content: "Web Service Architecture for Golang Developers"})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")
	router.HandleFunc("/posts", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", removePost).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)

	for _, post := range posts {
		if post.ID == params["id"] {
			json.NewEncoder(w).Encode(&post)
		}
	}
}

func addPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)

	posts = append(posts, post)
	json.NewEncoder(w).Encode(posts)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)

	for i, item := range posts {
		if item.ID == post.ID {
			posts[i] = post
		}
	}

	json.NewEncoder(w).Encode(posts)
}

func removePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var Post Post

	Post.ID = params["id"]

	for i, item := range posts {
		if item.ID == Post.ID {
			posts = append(posts[:i], posts[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(posts)
}
