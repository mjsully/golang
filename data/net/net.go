package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Message string
}

func main() {

	db, err := setupDatabase()
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Message{})

	fmt.Println("Starting REST API on :8080")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /test", testFunctionGet)
	mux.HandleFunc("GET /test/{id}", testFunctionGetId)
	mux.HandleFunc("POST /test", testFunctionPost)

	log.Fatal(http.ListenAndServe(":8080", mux))

}

func setupDatabase() (*gorm.DB, error) {

	db, dbErr := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	return db, dbErr

}

func testFunctionGet(w http.ResponseWriter, r *http.Request) {

	db, dbErr := setupDatabase()
	if dbErr != nil {
		http.Error(w, dbErr.Error(), http.StatusBadRequest)
	}

	// Debugging statements
	fmt.Println("HTTP endpoint reached")

	// Handle GET request
	var messages []Message
	result := db.Find(&messages)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
	} else {
		fmt.Println(result.RowsAffected)

	}
}

func testFunctionGetId(w http.ResponseWriter, r *http.Request) {

	db, dbErr := setupDatabase()
	if dbErr != nil {
		http.Error(w, dbErr.Error(), http.StatusBadRequest)
	}

	// Debugging statements
	fmt.Println("HTTP endpoint reached")
	id := r.PathValue("id")

	// Handle GET request on ID
	var m Message
	err := db.First(&m, id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, m.Message)
	}

}

func testFunctionPost(w http.ResponseWriter, r *http.Request) {

	db, dbErr := setupDatabase()
	if dbErr != nil {
		http.Error(w, dbErr.Error(), http.StatusBadRequest)
	}

	// Debugging statements
	fmt.Println("HTTP endpoint reached")
	fmt.Fprint(w, "POST")

	// Handle POST request
	var m Message
	jsonErr := json.NewDecoder(r.Body).Decode(&m)
	db.Create(&Message{Message: m.Message})
	fmt.Println(m)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusBadRequest)
		return
	}
}
