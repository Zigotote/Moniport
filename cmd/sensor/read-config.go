package main

import (
	"Moniport/internal/data"
	"Moniport/internal/helpers/errorHandler"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

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
	errorHandler.CheckError(err)

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &_configuration)

	errorHandler.CheckError(err)

	fmt.Println("file read")
	return _configuration
}

func generateSensorFromConfig(config data.Configuration) data.Sensor {
	return data.Sensor{
		Id:         config.IDSensor,
		IdAirport:  config.IDAirport,
		Measure:    getMeasureType(config.IDSensor),
		MqttAdress: config.AdressBroker,
		MqttPort:   config.PortBroker,
		MqttQos:    config.LevelQos,
	}
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
