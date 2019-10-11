package main

import (
	"Moniport/cmd/mqtt"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

type configuration struct {
	adressBroker string
	portBroker   string
	levelQos     int
	idClient     string
}

func main() {
	//Lecture de l'adresse du fichier de config => lancer "go run <adresse-main.go> -config <adresse-config.json>
	configFilename := getArgConfig()
	fmt.Println("configFilename : ", configFilename)

	//Lecture du fichier de configuration
	config := readConfiguration(configFilename)
	fmt.Println(config)

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

func getArgConfig() string {
	var configFilename string
	flag.StringVar(&configFilename, "config", "", "Usage")
	flag.Parse()
	return configFilename
}

func readConfiguration(filename string) configuration {

	var _configuration configuration
	//filename is the path to the json config file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error 1")
		return _configuration
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&_configuration)
	if err != nil {
		fmt.Println("Error 2")
		return _configuration
	}

	//TODO error handling

	fmt.Println("file read")
	fmt.Println(_configuration.adressBroker)
	return _configuration
}
