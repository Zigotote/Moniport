package main

import (
	"math/rand"
	"strconv"
	"time"
)

type Measure string

const (
	TEMPERATURE Measure = "temp"
	PRESSURE    Measure = "press"
	WIND        Measure = "wind"
)

type Sensor struct {
	id         int
	idAirport  string
	measure    Measure
	mqttAdress string
	mqttPort   int
	mqttQos    byte
}

type Message struct {
	IDSensor    string  `json:"idSensor"`
	IDAirport   string  `json:"idAirport"`
	MeasureType Measure `json:"measure"`
	Value       float32 `json:"value"`
	Date        string  `json:"date"`
}

func (s Sensor) GenerateMessage(date time.Time) Message {
	return Message{
		IDSensor:    strconv.Itoa(s.id),
		IDAirport:   s.idAirport,
		MeasureType: s.measure,
		Value:       s.randValue(),
		Date:        convertTimeToDate(date),
	}
}

func convertTimeToDate(date time.Time) string {
	return date.Format("2006-01-02-15-04-05")
}

func (s Sensor) randValue() float32 {
	r := rand.Float32()
	switch s.measure {
	case TEMPERATURE:
		r = r * 40
	case PRESSURE:
		r = 1013.25 + (r * 10)
	case WIND:
		r = r * 100
	}
	return r
}
