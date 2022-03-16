package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSensorStatusAllGood(t *testing.T) {
	sensor := new(Sensor)

	sensor.MinimumTemperature = 2
	sensor.MaximumTemperature = 5
	sensor.Temperature = 3

	sensor.CheckStatus()

	assert.Equal(t, TemperatureStatuses.AllGood, string(sensor.Status))
}

func TestSensorStatusTooLow(t *testing.T) {
	sensor := new(Sensor)

	sensor.MinimumTemperature = -1
	sensor.MaximumTemperature = 5
	sensor.Temperature = -2

	sensor.CheckStatus()

	assert.Equal(t, TemperatureStatuses.TooLow, string(sensor.Status))
}

func TestSensorStatusTooHigh(t *testing.T) {
	sensor := new(Sensor)

	sensor.MinimumTemperature = -5
	sensor.MaximumTemperature = -1
	sensor.Temperature = 0

	sensor.CheckStatus()

	assert.Equal(t, TemperatureStatuses.TooHigh, string(sensor.Status))
}
