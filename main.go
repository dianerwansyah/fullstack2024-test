package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Client struct {
	ID           int        `json:"id" gorm:"primary_key"`
	Name         string     `json:"name"`
	Slug         string     `json:"slug"`
	IsProject    string     `json:"is_project"`
	SelfCapture  string     `json:"self_capture"`
	ClientPrefix string     `json:"client_prefix"`
	ClientLogo   string     `json:"client_logo"`
	Address      string     `json:"address"`
	PhoneNumber  string     `json:"phone_number"`
	City         string     `json:"city"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

var db *gorm.DB

func main() {
	db.InitDB()
	redis.InitRedis()

	r := mux.NewRouter()
	r.HandleFunc("/clients", handler.CreateClient).Methods("POST")
	r.HandleFunc("/clients/{id}", handler.UpdateClient).Methods("PUT")
	r.HandleFunc("/clients/{id}", handler.DeleteClient).Methods("DELETE")
	r.HandleFunc("/clients", handler.GetClients).Methods("GET")

	fmt.Println("Server running at :8080")
	fmt.Println(http.ListenAndServe(":8080", r))
}
