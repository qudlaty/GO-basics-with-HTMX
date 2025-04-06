package main

import (
	"log"
	"net/http"
	"text/template"
)

type ExampleInventoryInterface struct {
	ItemName    string
	ItemMovable string
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		inventory := map[string][]ExampleInventoryInterface{
			"Items": {
				{ItemName: "M40 Rifle", ItemMovable: "true"},
				{ItemName: "Barrel", ItemMovable: "true"},
				{ItemName: "Door", ItemMovable: "false"},
			},
		}
		tmpl.Execute(w, inventory)
	}

	http.HandleFunc("/", h1)

	http.ListenAndServe(":8001", nil)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
