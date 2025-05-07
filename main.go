package main

import (
	"fmt"
	"net/http"

	"fullstack2024/db"
	"fullstack2024/handler"
	"fullstack2024/redis"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()
	redis.InitRedis()

	r := mux.NewRouter()
	r.HandleFunc("/clients", handler.CreateClient).Methods("POST")
	r.HandleFunc("/clients/{id}", handler.UpdateClient).Methods("PUT")
	r.HandleFunc("/clients/{id}", handler.DeleteClient).Methods("DELETE")
	r.HandleFunc("/clients", handler.GetClients).Methods("GET")

	fmt.Println("Server running at :8080")
	http.ListenAndServe(":8080", r)
}
