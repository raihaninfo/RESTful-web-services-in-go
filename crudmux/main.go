package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	{3, "Honda", "Civic", 15000},
}

func allCars(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicle)
}

func getCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	for _, car := range vehicle {
		if car.Id == carId {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(car)
		}
	}
}

func updateCars(w http.ResponseWriter, r *http.Request) {

}

func addCars(w http.ResponseWriter, r *http.Request) {
	var newCar Vehicle
	json.NewDecoder(r.Body).Decode(&newCar)
	vehicle = append(vehicle, newCar)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(vehicle)
}
func deleteCars(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range vehicle {
		if v.Id == carId {
			vehicle = append(vehicle[:k], vehicle[k+1:]...)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicle)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/cars", allCars).Methods("GET")
	router.HandleFunc("/cars/{id}", getCarById).Methods("GET")
	router.HandleFunc("/cars/{id}", updateCars).Methods("PUT")
	router.HandleFunc("/cars", addCars).Methods("POST")
	router.HandleFunc("/cars/{id}", deleteCars).Methods("DELETE")
	http.ListenAndServe(":8081", router)

}
