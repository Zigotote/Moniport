package main

import (
	"Moniport/internal/data"
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

func main() {
	//Lecture de l'adresse du fichier de config => lancer "go run <adresse-main.go> -config <adresse-config.json>
	configFilename := getArgConfig()
	fmt.Println("configFilename : ", configFilename)

	//Lecture du fichier de configuration
	config := readConfiguration(configFilename)
	fmt.Println(config)

	s1 := data.Sensor{
		Id:         config.IDSensor,
		IdAirport:  config.IDAirport,
		Measure:    getMeasureType(config.IDSensor),
		MqttAdress: config.AdressBroker,
		MqttPort:   config.PortBroker,
		MqttQos:    config.LevelQos,
	}
	mqttAdress := "tcp://" + s1.MqttAdress + ":" + strconv.Itoa(s1.MqttPort)
	c1 := mqtt.Connect(mqttAdress, s1.IdAirport+":"+strconv.Itoa(s1.Id))

	for range time.Tick(10 * time.Second) {
		fmt.Printf("Envoi message...")
		out, err := json.Marshal(s1.GenerateMessage(time.Now()))
		if err != nil {
			log.Fatal(err)
		}
		c1.Publish("airport_measures", s1.MqttQos, false, out)
	}
}

func getArgConfig() string {
	var configFilename string
	flag.StringVar(&configFilename, "config", "", "Usage")
	flag.Parse()
	return configFilename
}

func readConfiguration(filename string) data.Configuration {

	var _configuration data.Configuration
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

func getMeasureType(id int) data.MeasureType {
	switch id {
	case 0:
		return data.TEMPERATURE
	case 1:
		return data.PRESSURE
	case 2:
		return data.WIND
	default:
		return data.TEMPERATURE
	}
}
