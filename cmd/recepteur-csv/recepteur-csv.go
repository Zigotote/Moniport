package main

import (
	data "Moniport/internal/data"
	errorHandler "Moniport/internal/helpers/errorHandler"
	mqtt "Moniport/internal/helpers/mqtt"
	"Moniport/internal/helpers/readConfig"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	mymqtt "github.com/eclipse/paho.mqtt.golang"
)

var path_csv_dir = filepath.Join("src", "Moniport", "ressources", "csv-files") + string(os.PathSeparator)

func main() {
	//Lecture de l'IdAirport donnée en argument => lancer "go run <adresse-recepteur.exe> -config <IdAirport>
	topic := readConfig.GetArgConfig()

	client := mqtt.Connect("tcp://localhost:1883", "my-subscriber")
	for true {
		client.Subscribe(topic, 0, callbackFunction)
	}
}

var callbackFunction mymqtt.MessageHandler = func(client mymqtt.Client, msg mymqtt.Message) {

	//Ecriture dans le fichier voulu en fonction du type de mesure
	newMeasure := data.Measure{}
	err := json.Unmarshal(msg.Payload(), &newMeasure)
	fmt.Println(newMeasure)
	errorHandler.CheckError(err)

	writeMeasure(newMeasure)
}

func writeMeasure(m data.Measure) {

	csvDir := os.Getenv("GOPATH") + string(os.PathSeparator) + path_csv_dir

	os.MkdirAll(csvDir, 0700)

	file := openFile(csvDir + m.IDAirport + "-" + m.Date[0:10] + "-" + strings.ToUpper(m.MeasureType) + ".csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	line := []string{strings.Replace(m.Date[11:], "-", ":", 2), fmt.Sprintf("%f", m.Value)}

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

		file, err := os.OpenFile(file_path, os.O_APPEND|os.O_WRONLY, 0666)
		errorHandler.CheckError(err)

		writer := csv.NewWriter(file)
		defer writer.Flush()

		line := []string{"Date", "Value"}

		err = writer.Write(line)
		errorHandler.CheckError(err)
	}

	file, err := os.OpenFile(file_path, os.O_APPEND|os.O_WRONLY, 0666)
	errorHandler.CheckError(err)

	return file
}
