package main

import (
	"encoding/json"
	"github.com/escatran/golang-repo/crud-rest-api/model"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	model.ConnectDatabase()
	defer model.DisconnectDatabase()
	router := mux.NewRouter()
	router.HandleFunc("/record", createReport).Methods("POST")
	//	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":9191", router)
}

func createReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var record model.Record
	_ = json.NewDecoder(r.Body).Decode(&record)
	model.CreateRecord(&record)
	json.NewEncoder(w).Encode(&record)
}
