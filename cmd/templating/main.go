package main

import (
	data "Moniport/internal/data"
	redis "Moniport/internal/helpers/redis"
	measuresrtrv "Moniport/internal/measuresrtrv"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func templateHandeler(w http.ResponseWriter, r *http.Request) {
	/*datas := data.Datas{Aeroports: []data.Aeroport{
		{Name: "Nantes", Id: "NAN"},
		{Name: "Londres", Id: "LND"},
	}}*/
	datas := data.Datas{Aeroports: []data.Aeroport{}}
	airportIds := measuresrtrv.GetAirports()
	for _, element := range airportIds {
		datas.Aeroports = append(datas.Aeroports, data.Aeroport{Name: element, Id: element})
	}
	t := template.Must(template.ParseFiles("tmpl/pageAcceuil.html"))
	fmt.Println(t.Execute(w, datas))
}

func airportHandeler(w http.ResponseWriter, r *http.Request) {
	airport := strings.Replace(r.URL.Path, "/airport/", "", 1)
	datas := data.AirportData{Airportname: airport}
	t := template.Must(template.ParseFiles("tmpl/aeroportDetails.html"))

	fmt.Println(t.Execute(w, datas))
}

func main() {
	redis.Connect()

	http.Handle("/tmpl/", http.StripPrefix("/tmpl/", http.FileServer(http.Dir("tmpl"))))
	http.HandleFunc("/", templateHandeler)
	http.HandleFunc("/airport/", airportHandeler)
	http.ListenAndServe(":8085", nil)
	defer redis.CloseConnection()
}
