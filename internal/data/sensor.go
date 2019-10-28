package data

import (
	"Moniport/internal/helpers/date"
	"math/rand"
	"strconv"
	"time"
)

type Sensor struct {
	Id         int
	IdAirport  string
	Measure    MeasureType
	MqttAdress string
	MqttPort   int
	MqttQos    byte
}

func (s Sensor) GenerateMessage(d time.Time, previous_measure float64) Measure {
	return Measure{
		IDSensor:    strconv.Itoa(s.Id),
		IDAirport:   s.IdAirport,
		MeasureType: s.Measure.String(),
		Value:       s.randValue(previous_measure),
		Date:        date.GetStringFromDate(d),
	}
}

func (s Sensor) randValue(previous_measure float64) float64 {
	r := rand.Float64()

	if previous_measure == -1 {
		switch s.Measure {
		case TEMPERATURE:
			r = r * 40
		case PRESSURE:
			r = 1013.25 + (r * 10)
		case WIND:
			r = r * 100
		}
	} else {
		switch s.Measure {
		case TEMPERATURE:
			r = previous_measure + r*1 - 0.5
		case PRESSURE:
			r = previous_measure + r*4 - 2
		case WIND:
			r = previous_measure + r*10 - 5
		}
	}

	return r
}
