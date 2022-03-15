package controllers

import (
	"encoding/json"
	"net/http"
	"server/data"
	"server/providers"

	"github.com/gorilla/mux"
)

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
