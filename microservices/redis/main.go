package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var ctx = context.Background()

// Redis client setup
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/cache/{id:[0-9]+}", cacheUser).Methods("GET")
	r.HandleFunc("/cache/{id:[0-9]+}", evictCache).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Redis Microservice listening on port 8001")
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func cacheUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userData := "user data for " + id // Example data
	err := rdb.Set(ctx, id, userData, 24*time.Hour).Err()
	if err != nil {
		http.Error(w, "Error caching user", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Cached user %s", id)
}

func evictCache(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := rdb.Del(ctx, id).Err()
	if err != nil {
		http.Error(w, "Error evicting cache", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Evicted user %s", id)
}
