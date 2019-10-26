package main

import (
	data "Moniport/internal/data"
	"Moniport/internal/helpers/date"
	measuresdata "Moniport/internal/measuresdata"
	measuretreatement "Moniport/internal/measurestreatment"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

func templateHandeler(w http.ResponseWriter, r *http.Request) {

	datas := data.Datas{Aeroports: []data.Aeroport{}}
	airportIds := measuresdata.GetAirports()
	for _, element := range airportIds {
		datas.Aeroports = append(datas.Aeroports, data.Aeroport{Name: element, Id: element})
	}
	t := template.Must(template.ParseFiles("tmpl/pageAcceuil.html"))
	fmt.Println(t.Execute(w, datas))
}

func getAverages(selectedDate string, airport string) [3]int {
	start := date.ParseDate(selectedDate + "-00-00-00")
	end := date.ParseDate(selectedDate + "-23-59-59")
	measureTypes := []data.MeasureType{data.PRESSURE, data.TEMPERATURE, data.WIND}
	var resp [3]int

	for i, m := range measureTypes {
		values := measuresdata.GetMeasuresInRange(airport, m.String(), start, end)
		if len(values) != 0 {
			avg := measuretreatement.GetAverageFromMeasures(values)
			resp[i] = (int)(avg)
		}
	}

	return resp

}

func getValuesfromArgs(args []string) (string, string, time.Time, time.Time, string, string, string) {
	airport := args[0]
	mesureType := ""
	startTimeStamp := time.Now()
	endTimeStamp := time.Now()
	startDate := "2019-10-16T07:30"
	endDate := "2019-10-17T19:30"
	selectedDate := "2019-10-16"

	if len(args) > 2 {
		mesureType = args[1]
	}
	if len(args) > 3 {
		startTimeStamp = date.ParseHTMLDate(args[2])
		startDate = args[2]
	}
	if len(args) > 4 {
		endTimeStamp = date.ParseHTMLDate(args[3])
		endDate = args[3]
	}
	if len(args) > 5 {
		selectedDate = args[4]
	}

	return airport, mesureType, startTimeStamp, endTimeStamp, startDate, endDate, selectedDate
}

func getLineGraphData(airport string, mesureType string, startTimeStamp, endTimeStamp time.Time) ([]int, []string) {
	tempDatas := measuresdata.GetMeasuresInRange(airport, mesureType, startTimeStamp, endTimeStamp)

	var graphDatas []int
	var graphDates []string
	for _, element := range tempDatas {
		graphDatas = append(graphDatas, (int)(element.Value))
		graphDates = append(graphDates, element.Date)
	}
	return graphDatas, graphDates
}

func airportHandeler(w http.ResponseWriter, r *http.Request) {
	sufixe := strings.Replace(r.URL.Path, "/airport/", "", 1)
	args := strings.Split(sufixe, "/")

	airport, mesureType, startTimeStamp, endTimeStamp, startDate, endDate, selectedDate := getValuesfromArgs(args)

	graphDatas, graphDates := getLineGraphData(airport, mesureType, startTimeStamp, endTimeStamp)

	types := []data.MeasureType{data.PRESSURE, data.TEMPERATURE, data.WIND}

	moyennes := getAverages(selectedDate, airport)

	datas := data.AirportData{Airportname: airport,
		Types:        types,
		AirportId:    airport,
		Startime:     startDate,
		EndTime:      endDate,
		GraphData:    graphDatas,
		SelectedType: mesureType,
		GraphDates:   graphDates,
		Moyennes:     moyennes,
		Date:         selectedDate,
		Graph1Error:  len(graphDatas) == 0,
		Graph2Error:  len(moyennes) == 0,
	}

	t := template.Must(template.ParseFiles("tmpl/aeroportDetails.html"))

	fmt.Println(t.Execute(w, datas))
}

func main() {
	measuresdata.Connect()

	http.Handle("/tmpl/", http.StripPrefix("/tmpl/", http.FileServer(http.Dir("tmpl"))))
	http.HandleFunc("/", templateHandeler)
	http.HandleFunc("/airport/", airportHandeler)
	http.ListenAndServe(":8085", nil)
	defer measuresdata.Disconnect()
}
