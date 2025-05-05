package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux" // Library untuk routing HTTP
)

// Struct Toy untuk merepresentasikan data mainan
type Toy struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Variabel global
var (
	toys   []Toy      // Menyimpan list mainan
	nextID = 1        // ID mainan berikutnya
	mu     sync.Mutex // Mutex untuk menghindari race condition
)

// Handler GET /toys
func getToys(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(toys)
}

// Handler POST /toys
func createToy(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var toy Toy
	if err := json.NewDecoder(r.Body).Decode(&toy); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	toy.ID = nextID          // Set ID mainan
	nextID++                 // Increment ID
	toys = append(toys, toy) // Tambahkan mainan ke list

	json.NewEncoder(w).Encode(toy)

	// Logging async pakai goroutine
	go logCreation(toy)
}

// Handler PUT /toys/{id}
func updateToy(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i := range toys {
		if toys[i].ID == id {
			if err := json.NewDecoder(r.Body).Decode(&toys[i]); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			toys[i].ID = id
			json.NewEncoder(w).Encode(toys[i])
			return
		}
	}

	http.Error(w, "Toy not found", http.StatusNotFound)
}

// Handler DELETE /toys/{id}
func deleteToy(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, toy := range toys {
		if toy.ID == id {
			toys = append(toys[:i], toys[i+1:]...)
			fmt.Fprintf(w, "Toy with ID %d deleted", id)
			return
		}
	}

	http.Error(w, "Toy not found", http.StatusNotFound)
}

// Fungsi logCreation untuk mencetak log dengan delay acak (simulasi async)
func logCreation(toy Toy) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("New toy created: %+v\n", toy)
}

// Fungsi utama untuk menjalankan server
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/toys", getToys).Methods("GET")
	r.HandleFunc("/toys", createToy).Methods("POST")
	r.HandleFunc("/toys/{id}", updateToy).Methods("PUT")
	r.HandleFunc("/toys/{id}", deleteToy).Methods("DELETE")

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
