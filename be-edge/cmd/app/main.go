package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"be-edge/internal/handler"

	"be-edge/internal/repository"
	"be-edge/internal/service"

	_ "github.com/mattn/go-sqlite3"
)

func enableCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}


// @title News API
// @version 1.0
// @description API for managing news posts
// @host localhost:8082
// @BasePath /
func main() {
	db := repository.NewDB("./database/news.db")
	defer db.Close()

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	router := mux.NewRouter()

	router.HandleFunc("/posts", postHandler.CreatePost).Methods("POST")
	router.HandleFunc("/posts", postHandler.GetPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", postHandler.GetPostByID).Methods("GET")
	router.HandleFunc("/posts/{id}", postHandler.DeletePost).Methods("DELETE")
	router.HandleFunc("/posts/{id}", postHandler.UpdatePost).Methods("PUT")


	fmt.Println("Server started at :8082")
	log.Fatal(http.ListenAndServe(":8082", enableCorsMiddleware(router)))
}
