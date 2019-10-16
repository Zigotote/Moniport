package main

import (
	"Moniport/internal/helpers/mqtt"
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
	PortBroker   int    `json:"portBroker"`
	LevelQos     byte   `json:"levelQos"`
	IDSensor     int    `json:"idSensor"`
	IDAirport    string `json:"idAirport"`
}

func main() {
	//Lecture de l'adresse du fichier de config => lancer "go run <adresse-main.go> -config <adresse-config.json>
	configFilename := getArgConfig()
	fmt.Println("configFilename : ", configFilename)

	//Lecture du fichier de configuration
	config := readConfiguration(configFilename)
	fmt.Println(config)

	s1 := Sensor{
		id:         config.IDSensor,
		idAirport:  config.IDAirport,
		measure:    getMeasureType(config.IDSensor),
		mqttAdress: config.AdressBroker,
		mqttPort:   config.PortBroker,
		mqttQos:    config.LevelQos,
	}
	mqttAdress := "tcp://" + s1.mqttAdress + ":" + strconv.Itoa(s1.mqttPort)
	c1 := mqtt.Connect(mqttAdress, s1.idAirport+":"+strconv.Itoa(s1.id))

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
		log.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &_configuration)

	if err != nil {
		log.Fatal(err)
	}

	//TODO error handling

	fmt.Println("file read")
	return _configuration
}

func getMeasureType(id int) Measure {
	switch id {
	case 0:
		return TEMPERATURE
	case 1:
		return PRESSURE
	case 2:
		return WIND
	default:
		return TEMPERATURE
	}
}
