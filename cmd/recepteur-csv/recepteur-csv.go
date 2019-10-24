package main

import (
	data "Moniport/internal/data"
	errorHandler "Moniport/internal/helpers/errorHandler"
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

var path_csv_dir = filepath.Join("src", "moniport", "cmd", "csv-files") + string(os.PathSeparator)

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
	errorHandler.CheckError(err)
}

func writeMeasure(m data.Measure) {

	csvDir := os.Getenv("GOPATH") + string(os.PathSeparator) + path_csv_dir

	os.MkdirAll(csvDir, 0700)
	file := openFile(csvDir + m.IDAirport + "-" + m.Date[0:10] + "-" + strings.ToUpper(m.MeasureType) + ".csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	line := []string{m.Date[10:], fmt.Sprintf("%f", m.Value)}

	err := writer.Write(line)
	errorHandler.CheckError(err)
	fmt.Println("Ecriture dans le fichier " + file.Name())
}

func openFile(file_path string) *os.File {

	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		fmt.Println("Création du fichier...")
		f, err := os.Create(file_path)

		errorHandler.CheckError(err)
		fmt.Println("fichier créé : " + f.Name())
		f.Close()

		writer := csv.NewWriter(f)
		defer writer.Flush()

		line := []string{"Date", "Value"}

		err = writer.Write(line)
		errorHandler.CheckError(err)
	}

	file, err := os.Open(file_path)
	errorHandler.CheckError(err)

	return file
}
