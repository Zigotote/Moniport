package main

import (
	"fmt"
	redis "moniport/cmd/recepteur/redis"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type measure struct {
	idSensor    string
	idAirport   string
	measureType string
	value       float64
	date        string
}

func main() {

	var sampleMeasure measure = measure{"1", "NON", "wind", 50, "2019-12-10-15-10-25"}

	redis.Connect()
	sampleMeasure.sendMeasure()
	defer redis.CloseConnection()
}

var callbackFunction mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

}

func (m measure) sendMeasure() {
	setKey := m.idAirport + ":" + m.measureType

	redis.AddToSet("airports", m.idAirport)

	setValue := fmt.Sprintf("%d_%.2f", getNewIdMeasure(), m.value)

	setTimestamp := getTimestampFromDate(m.date)

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
