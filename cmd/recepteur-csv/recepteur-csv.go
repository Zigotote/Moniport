package main

import (
	data "Moniport/internal/data"
	mqtt "Moniport/internal/helpers/mqtt"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	mymqtt "github.com/eclipse/paho.mqtt.golang"
)

var path_csv_dir = filepath.Join("Moniport", "cmd", "csv-files") + string(os.PathSeparator)

func main() {

	var sampleMeasure data.Measure = data.Measure{"1", "NON", "wind", 50, "2019-12-10-15-10-25"}

	writeMeasure(sampleMeasure)

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
}

func openFile(file_path string) *os.File {

	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		fmt.Println("Création du fichier...")
		f, err := os.Create(file_path)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("fichier créé : " + f.Name())
		f.Close()

	}

	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func writeMeasure(m data.Measure) {
	file := openFile(os.Getenv("GOPATH") + string(os.PathSeparator) + path_csv_dir + m.IDAirport + "-" + m.Date[0:10] + "-" + strings.ToUpper(m.MeasureType) + ".csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	line := []string{m.Date[10:], fmt.Sprintf("%f", m.Value)}

	err := writer.Write(line)
	if err != nil {
		log.Fatal(err)
	}
}
