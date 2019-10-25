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

	mqttAdress := "tcp://" + s1.mqttAdress + ":" + strconv.Itoa(s1.mqttPort)
	c1 := mqtt.Connect(mqttAdress, s1.idAirport+":"+strconv.Itoa(s1.id))

	for range time.Tick(10 * time.Second) {
		fmt.Printf("Envoi message...")
		out, err := json.Marshal(s1.GenerateMessage(time.Now()))
		errorHandler.CheckError(err)
		c1.Publish("airport_measures", s1.mqttQos, false, out)
	}
}
