package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type configuration struct {
	adressBroker string
	portBroker   string
	levelQos     int
	idClient     string
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
