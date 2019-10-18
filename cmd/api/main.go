package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/measures/{start}/{end}", measureHandler)
	router.HandleFunc("/avg-measures/{date}", avgMeasureHandler)
	err := http.ListenAndServe(":8081", router)
	log.Fatal(err)
}

func measureHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	start := getTimestampFromDate(vars["start"])
	end := getTimestampFromDate(vars["end"])
	if start < 0 || end < 0 {
		fmt.Fprint(w, "Vous devez renseigner les paramètres start et end pour effectuer la requête. Ils doivent être au format YYYY-MM-DD-hh-mm-ss.")
		return
	}
	//appel redis getMeasure(start, end)
	resp := Message{
		Date:        "1571351477",
		MeasureType: "temp",
		Value:       12,
		IDSensor:    "1",
		IDAirport:   "NTE",
	}
	//voir pour convertir les timestamp en dates
	writeJSON(w, resp)
}

func avgMeasureHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := getTimestampFromDateDay(vars["date"])
	if date < 0 {
		fmt.Fprint(w, "Vous devez renseigner le paramètre date pour effectuer la requête. Il doit être au format YYYY-MM-DD")
		return
	}
	//appel redis getAvgMeasures(date)
	temp := Message{
		Date:        "1571351477",
		MeasureType: "temp",
		Value:       12,
		IDSensor:    "1",
		IDAirport:   "NTE",
	}
	wind := Message{
		Date:        "1571351477",
		MeasureType: "wind",
		Value:       80,
		IDSensor:    "2",
		IDAirport:   "NTE",
	}
	press := Message{
		Date:        "1571351477",
		MeasureType: "press",
		Value:       1200,
		IDSensor:    "3",
		IDAirport:   "NTE",
	}
	resp := []Message{temp, wind, press}
	//voir pour convertir les timestamp en dates
	writeJSON(w, resp)
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

//Type temporaire en attendant que le refacto soit ok
type Message struct {
	IDSensor    string  `json:"idSensor"`
	IDAirport   string  `json:"idAirport"`
	MeasureType string  `json:"measure"`
	Value       float32 `json:"value"`
	Date        string  `json:"date"`
}

//fonction reprise de recepteur.go, à refactorer
func getTimestampFromDate(date string) int64 {
	layout := "2006-01-02-15-04-05"

	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
	}

	return t.Unix()
}

//à mettre dans le même fichier que la méthode getTimestampFromDate après refactoring
func getTimestampFromDateDay(date string) int64 {
	layout := "2006-01-02"

	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
	}

	return t.Unix()
}
