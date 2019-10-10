package main

import (
	"Moniport/cmd/mqtt"
	"fmt"
	"strconv"
	"time"
)

func main() {
	s1 := Sensor{
		id:         001,
		idAirport:  "NTE",
		measure:    TEMPERATURE,
		mqttAdress: "localhost",
		mqttPort:   1883,
		mqttQos:    0,
	}
	mqttAdress := "tcp://" + s1.mqttAdress + ":" + strconv.Itoa(s1.mqttPort)
	c1 := mqtt.Connect(mqttAdress, strconv.Itoa(s1.id))

	for range time.Tick(10 * time.Second) {
		//fmt.Printf("Envoi message...")
		c1.Publish("topic", s1.mqttQos, false, s1.GenerateMessage(time.Now()))
	}

}
