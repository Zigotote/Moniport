package main

import (
	"fmt"
	redis "moniport/cmd/recepteur/redis"
	"time"
)

type measure struct {
	idSensor    string
	idAirport   string
	measureType string
	value       float64
	date        string
}

func main() {
	redis.Connect()
	redis.SendData("test", "yes")
	for i := 0; i <= 10; i++ {
		fmt.Println(getNewIdMeasure())
	}
	defer redis.CloseConnection()
}

func sendMeasure(m measure) {
	setKey := m.idAirport + ":" + m.measureType

	fmt.Println(getTimestampFromDate(m.date))
}

func getNewIdMeasure() int {
	if redis.KeyExists("currIdMeasure") {
		redis.IncrKey("currIdMeasure")
	} else {
		redis.SendData("currIdMeasure", "0")
	}

	return redis.GetDataInt("currIdMeasure")
}

func getTimestampFromDate(date string) int {
	layout := "2006-01-02-15-04-05"

	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
	}

	return t.Unix()
}
