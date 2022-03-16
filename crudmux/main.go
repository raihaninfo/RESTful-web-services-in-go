package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	Id    int
	Make  string
	Model string
	Price int
}

var vehicle = []Vehicle{
	{1, "Toyota", "Corolla", 10000},
	{2, "Toyota", "Camry", 20000},
	{1, "Honda", "Civic", 15000},
}

func allCars(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicle)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/cars", allCars).Methods("GET")
	http.ListenAndServe(":8081", router)

}
