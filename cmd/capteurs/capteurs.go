package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type configuration struct {
	AdressBroker string `json:"adressBroker"`
	PortBroker   string `json:"portBroker"`
	LevelQos     int    `json:"levelQos"`
	IDClient     int    `json:"idClient"`
}

func main() {
	var configFilename string
	flag.StringVar(&configFilename, "config", "", "Usage")
	flag.Parse()

	fmt.Println(configFilename)

	config := readConfiguration(configFilename)

	fmt.Println(config)
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

	defer file.Close()
	return _configuration
}
