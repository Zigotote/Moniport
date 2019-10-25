package main

import (
	data "Moniport/internal/data"
	measuresdata "Moniport/internal/measuresdata"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

func templateHandeler(w http.ResponseWriter, r *http.Request) {
	/*datas := data.Datas{Aeroports: []data.Aeroport{
		{Name: "Nantes", Id: "NAN"},
		{Name: "Londres", Id: "LND"},
	}}*/
	datas := data.Datas{Aeroports: []data.Aeroport{}}
	airportIds := measuresdata.GetAirports()
	for _, element := range airportIds {
		datas.Aeroports = append(datas.Aeroports, data.Aeroport{Name: element, Id: element})
	}
	t := template.Must(template.ParseFiles("tmpl/pageAcceuil.html"))
	fmt.Println(t.Execute(w, datas))
}

func airportHandeler(w http.ResponseWriter, r *http.Request) {
	airport := strings.Replace(r.URL.Path, "/airport/", "", 1)
	//airportIds := measuresdata.GetMeasures(airport, "temp")

	datas := data.AirportData{Airportname: airport,
		Types:     []string{data.PRESSURE, data.TEMPERATURE, data.WIND},
		AirportId: airport,
		Startime:  MeasureDateFromTimestamp(int64(time.Now().Minute())),
	}
	t := template.Must(template.ParseFiles("tmpl/aeroportDetails.html"))

	fmt.Println(t.Execute(w, datas))
}

//a recuperer d'ailleur
func MeasureDateFromTimestamp(date int64) string {
	layout := "2006-01-02-15-04-05"

	return time.Unix(date, 0).Format(layout)
}

func main() {
	measuresdata.Connect()

	http.Handle("/tmpl/", http.StripPrefix("/tmpl/", http.FileServer(http.Dir("tmpl"))))
	http.HandleFunc("/", templateHandeler)
	http.HandleFunc("/airport/", airportHandeler)
	http.ListenAndServe(":8085", nil)
	defer measuresdata.Disconnect()
}
