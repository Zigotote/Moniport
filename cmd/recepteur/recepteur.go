package main

import (
	mqtt "Moniport/cmd/mqtt"
	redis "Moniport/cmd/recepteur/redis"
	"encoding/json"
	"fmt"
	"log"
	"time"

	mymqtt "github.com/eclipse/paho.mqtt.golang"
)

type measure struct {
	IDSensor    string  `json:"idSensor"`
	IDAirport   string  `json:"idAirport"`
	MeasureType string  `json:"measure"`
	Value       float64 `json:"value"`
	Date        string  `json:"date"`
}

func main() {
	redis.Connect()
	//var sampleMeasure measure = measure{"1", "NON", "wind", 50, "2019-12-10-15-10-25"}
	client := mqtt.Connect("tcp://localhost:1883", "my-subscriber")
	for true {
		client.Subscribe("airport_measures", 0, callbackFunction)
	}
	defer redis.CloseConnection()
}

var callbackFunction mymqtt.MessageHandler = func(client mymqtt.Client, msg mymqtt.Message) {

	newMeasure := &measure{}
	err := json.Unmarshal(msg.Payload(), newMeasure)
	fmt.Println(*newMeasure)
	if err != nil {
		log.Fatal(err)
	}
	newMeasure.sendMeasure()

}

func (m measure) sendMeasure() {
	setKey := m.IDAirport + ":" + m.MeasureType

	redis.AddToSet("airports", m.IDAirport)

	setValue := fmt.Sprintf("%d_%.2f", getNewIdMeasure(), m.Value)

	setTimestamp := getTimestampFromDate(m.Date)

	redis.AddToOrdSet(setKey, setValue, setTimestamp)
}

func getNewIdMeasure() int {
	if redis.KeyExists("currIdMeasure") {
		redis.IncrKey("currIdMeasure")
	} else {
		redis.SendData("currIdMeasure", "0")
	}

	return redis.GetDataInt("currIdMeasure")
}

func getTimestampFromDate(date string) int64 {
	layout := "2006-01-02-15-04-05"

	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
	}

	return t.Unix()
}
