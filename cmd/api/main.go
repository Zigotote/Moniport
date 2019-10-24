package main

import (
	"Moniport/internal/data"
	"Moniport/internal/helpers/date"
	"Moniport/internal/helpers/redis"
	"Moniport/internal/measuresdata"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	redis.Connect()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/measures/{airport}/{measure}/{start}/{end}", measureHandler)
	router.HandleFunc("/avg-measures/{airport}/{date}", avgMeasureHandler)
	err := http.ListenAndServe(":8081", router)
	log.Fatal(err)
	defer redis.CloseConnection()
}

func measureHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	airport := vars["airport"]
	measure := vars["measure"]
	start := date.ParseDate(vars["start"])
	end := date.ParseDate(vars["end"])
	if start.Year() == 1 || end.Year() == 1 {
		fmt.Fprint(w, "Vous devez renseigner les paramètres start et end pour effectuer la requête. Ils doivent être au format YYYY-MM-DD-hh-mm-ss.")
		return
	}
	resp := measuresdata.GetMeasuresInRange(airport, measure, start, end)
	writeJSON(w, resp)
}

func avgMeasureHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := date.ParseDate(vars["date"])
	/*if date < 0 {
		fmt.Fprint(w, "Vous devez renseigner le paramètre date pour effectuer la requête. Il doit être au format YYYY-MM-DD")
		return
	}*/
	date.AddDate(1, 1, 1)
	//appel redis getAvgMeasures(date)
	temp := data.Measure{
		Date:        "1571351477",
		MeasureType: "temp",
		Value:       12,
		IDSensor:    "1",
		IDAirport:   "NTE",
	}
	wind := data.Measure{
		Date:        "1571351477",
		MeasureType: "wind",
		Value:       80,
		IDSensor:    "2",
		IDAirport:   "NTE",
	}
	press := data.Measure{
		Date:        "1571351477",
		MeasureType: "press",
		Value:       1200,
		IDSensor:    "3",
		IDAirport:   "NTE",
	}
	resp := []data.Measure{temp, wind, press}
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
