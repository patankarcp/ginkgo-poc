package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	dbUser     = "postgres"
	dbPassword = "password"
	dbName     = "usersdb"
)

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", deleteUser).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Postgres Microservice listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func dbConn() (db *sql.DB) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Implement CRUD handlers here
func getUsers(w http.ResponseWriter, r *http.Request)   { /* ... */ }
func getUser(w http.ResponseWriter, r *http.Request)    { /* ... */ }
func createUser(w http.ResponseWriter, r *http.Request) { /* ... */ }
func updateUser(w http.ResponseWriter, r *http.Request) { /* ... */ }
func deleteUser(w http.ResponseWriter, r *http.Request) { /* ... */ }
