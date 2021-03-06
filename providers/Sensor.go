package providers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"server/models"
)

const providerBaseAddress = "https://temperature-sensor-service.herokuapp.com/"

func GetSensorTemperature(sensor models.Sensor) int8 {
	response, err := http.Get(providerBaseAddress + "sensor/" + sensor.Id)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(body, &sensor); err != nil {
		panic(err)
	}

	return sensor.Temperature
}
