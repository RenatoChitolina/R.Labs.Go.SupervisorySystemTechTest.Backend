package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"server/models"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/temperature/{id}", GetSensorTemperature).Methods("GET", "OPTIONS")

	fmt.Println("Starting server on the port 8081...")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func GetSensorTemperature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	response, err := http.Get("https://temperature-sensor-service.herokuapp.com/sensor/" + params["id"])
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var sensor models.Sensor

	if err := json.Unmarshal(body, &sensor); err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(sensor)
}
