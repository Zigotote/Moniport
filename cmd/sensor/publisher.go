package main

import (
	"Moniport/internal/helpers/errorHandler"
	"Moniport/internal/helpers/mqtt"
	"Moniport/internal/helpers/readConfig"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func main() {
	//Lecture de l'adresse du fichier de config donnée en argument => lancer "go run <adresse-main.go> -config <adresse-config.json>
	configFilename := readConfig.GetArgConfig()
	fmt.Println("configFilename : ", configFilename)

	//Lecture du fichier de configuration
	config_data := readConfig.ReadConfigurationPublisher(configFilename)
	fmt.Println("config_data : ", config_data)

	s1 := readConfig.GenerateSensorFromConfig(config_data)

	mqttAdress := "tcp://" + s1.MqttAdress + ":" + strconv.Itoa(s1.MqttPort)
	c1 := mqtt.Connect(mqttAdress, s1.IdAirport+":"+strconv.Itoa(s1.Id))

	//Valeur de la mesure précédente, utilisé ici pour simuler des variations au lieu de valeurs totalement aléatoires
	var previous_measure_value float64 = -1

	for range time.Tick(10 * time.Second) {
		fmt.Printf("Envoi message...")
		measure := s1.GenerateMessage(time.Now(), previous_measure_value)
		out, err := json.Marshal(measure)
		errorHandler.CheckError(err)
		previous_measure_value = measure.Value
		//Le nom du topic correspond au nom de l'aéroport
		c1.Publish(s1.IdAirport, s1.MqttQos, false, out)
	}
}
