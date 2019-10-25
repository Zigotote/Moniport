package main

import (
	"Moniport/internal/helpers/errorHandler"
	"Moniport/internal/helpers/mqtt"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func main() {
	//Lecture de l'adresse du fichier de config donnée en argument => lancer "go run <adresse-main.go> -config <adresse-config.json>
	configFilename := getArgConfig()
	fmt.Println("configFilename : ", configFilename)

	//Lecture du fichier de configuration
	config := readConfiguration(configFilename)
	fmt.Println("config : ", config)

	s1 := generateSensorFromConfig(config)

	mqttAdress := "tcp://" + s1.MqttAdress + ":" + strconv.Itoa(s1.MqttPort)
	c1 := mqtt.Connect(mqttAdress, s1.IdAirport+":"+strconv.Itoa(s1.Id))

	for range time.Tick(10 * time.Second) {
		fmt.Printf("Envoi message...")
		out, err := json.Marshal(s1.GenerateMessage(time.Now()))
		errorHandler.CheckError(err)
		c1.Publish("airport_measures", s1.MqttQos, false, out)
	}
}
