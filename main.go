package main

import (
	"fmt"
	"log"
	"net/http"
	"server/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/sensors", controllers.GetAvailableSensors).Methods("GET", "OPTIONS")
	router.HandleFunc("/temperature/{id}", controllers.GetSensorTemperature).Methods("GET", "OPTIONS")

	fmt.Println("Starting server on the port 8081...")

	log.Fatal(http.ListenAndServe(":8081", router))
}
