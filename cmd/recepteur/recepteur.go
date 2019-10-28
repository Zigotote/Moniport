package main

import (
	data "Moniport/internal/data"
	mqtt "Moniport/internal/helpers/mqtt"
	"Moniport/internal/helpers/readConfig"
	"encoding/json"
	"fmt"
	"log"
	"moniport/internal/measuresdata"

	mymqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	//Lecture de l'IdAirport donnée en argument => lancer "go run <adresse-recepteur.exe> -config <IdAirport>
	topic := readConfig.GetArgConfig()

	measuresdata.Connect()
	// var sampleMeasure data.Measure = data.Measure{"1", "NON", "wind", 50, "2019-12-10-15-10-25"}
	client := mqtt.Connect("tcp://localhost:1883", "my-subscriber")
	for true {
		client.Subscribe(topic, 0, callbackFunction)
	}
	defer measuresdata.Disconnect()
}

var callbackFunction mymqtt.MessageHandler = func(client mymqtt.Client, msg mymqtt.Message) {

	newMeasure := data.Measure{}
	err := json.Unmarshal(msg.Payload(), &newMeasure)
	fmt.Println(newMeasure)
	if err != nil {
		log.Fatal(err)
	}
	measuresdata.SendMeasure(newMeasure)

}
