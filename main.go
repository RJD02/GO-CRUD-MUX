package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func apiHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello World"))
}

func main() {
	route := mux.NewRouter()

	s := route.PathPrefix("/api").Subrouter() // Base path
	s.HandleFunc("/", apiHome).Methods("GET")

	s.HandleFunc("/createProfile", createProfile).Methods("POST")
	s.HandleFunc("/getAllUsers", getAllUsers).Methods("GET")
	s.HandleFunc("/getUserProfile", getUserProfile).Methods("POST")
	s.HandleFunc("/updateProfile", updateProfile).Methods("PUT")
	s.HandleFunc("/deleteProfile/{id}", deleteProfile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", s))
}
