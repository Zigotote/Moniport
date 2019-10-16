package main

import (
	"math/rand"
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
	idSensor  int
	idAirport string
	measure   Measure
	value     float32
	date      time.Time
}

func (s Sensor) GenerateMessage(date time.Time) Message {
	return Message{
		idSensor:  s.id,
		idAirport: s.idAirport,
		measure:   s.measure,
		value:     s.randValue(),
		date:      date,
	}
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
