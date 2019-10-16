package main

import (
	"Moniport/cmd/mqtt"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type configuration struct {
	AdressBroker string `json:"adressBroker"`
	PortBroker   string `json:"portBroker"`
	LevelQos     int    `json:"levelQos"`
	IDClient     int    `json:"idClient"`
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
		fmt.Printf("Envoi message...")
		out, err := json.Marshal(s1.GenerateMessage(time.Now()))
		if err != nil {
			log.Fatal(err)
		}
		c1.Publish("airport_measures", s1.mqttQos, false, out)
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
		fmt.Println(err)
		return _configuration
	}

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &_configuration)

	if err != nil {
		fmt.Println(err)
		return _configuration
	}

	//TODO error handling

	fmt.Println("file read")
	return _configuration
}
