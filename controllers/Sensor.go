package controllers

import (
	"encoding/json"
	"net/http"
	"server/data"
	"server/providers"

	"github.com/gorilla/mux"
)

func GetAvailableSensors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data := data.SensorSetups()

	keys := make([]string, 0, len(data))

	for key := range data {
		keys = append(keys, key)
	}

	json.NewEncoder(w).Encode(keys)
}

func GetSensorTemperature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	data := data.SensorSetups()

	sensor, hasSensor := data[params["id"]]
	if !hasSensor {
		w.WriteHeader(404)
		return
	}

	sensor.Temperature = providers.GetSensorTemperature(sensor)
	sensor.CheckStatus()

	json.NewEncoder(w).Encode(sensor)
}
