package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type entry struct {
	Name string
	Done bool
}

type person struct {
	UserName string
}

type Aeroport struct {
	Name string
	id   string
}
type datas struct {
	Aeroports []Aeroport
}

func templateHandeler(w http.ResponseWriter, r *http.Request) {
	data := datas{Aeroports: []Aeroport{
		{Name: "Nantes", id: "NAN"},
		{Name: "Londres", id: "LND"},
	}}
	t := template.Must(template.ParseFiles("tmpl/pageAcceuil.html", "tmpl/css/theme.css"))
	fmt.Println(t.Execute(w, data))
}

func main() {
	http.HandleFunc("/", templateHandeler)
	http.ListenAndServe(":8080", nil)
}
