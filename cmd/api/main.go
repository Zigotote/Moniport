package main

import (
	"Moniport/internal/data"
	"Moniport/internal/helpers/date"
	"Moniport/internal/helpers/redis"
	"Moniport/internal/measuresdata"
	"Moniport/internal/measurestreatment"
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
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Vous devez renseigner les paramètres start et end pour effectuer la requête. Ils doivent être au format YYYY-MM-DD-hh-mm-ss.")
		return
	}
	resp := measuresdata.GetMeasuresInRange(airport, measure, start, end)
	writeJSON(w, resp)
}

func avgMeasureHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	airport := vars["airport"]
	start := date.ParseDate(vars["date"] + "-00-00-00")
	end := date.ParseDate(vars["date"] + "-23-59-59")
	if start.Year() == 1 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Vous devez renseigner le paramètre date pour effectuer la requête. Il doit être au format YYYY-MM-DD")
		return
	}
	measureTypes := []data.MeasureType{data.PRESSURE, data.TEMPERATURE, data.WIND}
	var resp [3]AvgMesure

	for i, m := range measureTypes {
		values := measuresdata.GetMeasuresInRange(airport, m.String(), start, end)
		if len(values) != 0 {
			avg := AvgMesure{
				MeasureType: m,
				Value:       measuretreatement.GetAverageFromMeasures(values),
			}
			resp[i] = avg
		}
	}
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

type AvgMesure struct {
	MeasureType data.MeasureType
	Value       float64
}
