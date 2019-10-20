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

type Person struct {
	UserName string
}

func templateHandeler(w http.ResponseWriter, r *http.Request) {
	p := Person{UserName: "toto"}
	t := template.Must(template.ParseFiles("tmpl/test.html"))
	fmt.Println(t.Execute(w, p))
}

func main() {
	http.HandleFunc("/", templateHandeler)
	http.ListenAndServe(":8080", nil)
}
