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

func (s Sensor) GenerateMessage(d time.Time) Measure {
	return Measure{
		IDSensor:    strconv.Itoa(s.Id),
		IDAirport:   s.IdAirport,
		MeasureType: s.Measure.String(),
		Value:       s.randValue(),
		Date:        date.GetStringFromDate(d),
	}
}

func (s Sensor) randValue() float64 {
	r := rand.Float64()
	switch s.Measure {
	case TEMPERATURE:
		r = r * 40
	case PRESSURE:
		r = 1013.25 + (r * 10)
	case WIND:
		r = r * 100
	}
	return r
}
