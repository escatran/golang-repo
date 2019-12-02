package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"math/rand"
	"net/http"
	"strconv"
)

type Record struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/record", createReport).Methods("POST")
	//	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":9191", router)
}

func createReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var record Record
	_ = json.NewDecoder(r.Body).Decode(&record)
	record.ID = strconv.Itoa(rand.Intn(1000000))
	json.NewEncoder(w).Encode(&record)

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Record{})

	// Create
	db.Create(&record)
}


