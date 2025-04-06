package main

import (
	"log"
	"net/http"
	"text/template"
)

type Objects_entities struct {
	ItemName    string
	ItemMovable bool
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		inventory := map[string][]Objects_entities{
			"Items": {
				{ItemName: "M40 Rifle", ItemMovable: true},
				{ItemName: "Barrel", ItemMovable: true},
				{ItemName: "Door", ItemMovable: false},
			},
		}
		tmpl.Execute(w, inventory)
	}

	http.HandleFunc("/", h1)

	http.ListenAndServe(":8001", nil)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
