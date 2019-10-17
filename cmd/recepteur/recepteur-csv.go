package main

import (
	data "Moniport/internal/data"
	mqtt "Moniport/internal/helpers/mqtt"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	mymqtt "github.com/eclipse/paho.mqtt.golang"
)

type Files_writers struct {
	temp  *csv.Writer
	press *csv.Writer
	wind  *csv.Writer
}

//Variable globale contenant les writers écrivants dans les fichiers
var files_writers Files_writers

func main() {

	//Création des fichiers et des writers écrivants dedans
	//TODO récupération des noms de fichiers
	file_temp := openFile("")
	defer file_temp.Close()
	file_press := openFile("")
	defer file_press.Close()
	file_wind := openFile("")
	defer file_wind.Close()

	files_writers.temp = csv.NewWriter(file_temp)
	files_writers.press = csv.NewWriter(file_press)
	files_writers.wind = csv.NewWriter(file_wind)
	defer files_writers.temp.Flush()
	defer files_writers.press.Flush()
	defer files_writers.wind.Flush()

	//var sampleMeasure measure = measure{"1", "NON", "wind", 50, "2019-12-10-15-10-25"}
	client := mqtt.Connect("tcp://localhost:1883", "my-subscriber")
	for true {
		client.Subscribe("airport_measures", 0, callbackFunction)
	}
}

var callbackFunction mymqtt.MessageHandler = func(client mymqtt.Client, msg mymqtt.Message) {

	//Ecriture dans le fichier voulu en fonction du type de mesure
	newMeasure := data.Measure{}
	err := json.Unmarshal(msg.Payload(), &newMeasure)
	fmt.Println(newMeasure)
	if err != nil {
		log.Fatal(err)
	}

	switch newMeasure.MeasureType {
	case "temp":
		writeMeasure(newMeasure, files_writers.temp)
	case "press":
		writeMeasure(newMeasure, files_writers.press)
	case "wind":
		writeMeasure(newMeasure, files_writers.wind)
	}
}

func openFile(filneName string) *os.File {
	file, err := os.Open("")
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func writeMeasure(m data.Measure, writer *csv.Writer) {
	line := []string{m.Date, fmt.Sprintf("%f", m.Value)}

	err := writer.Write(line)
	if err != nil {
		log.Fatal(err)
	}
}
