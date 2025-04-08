package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Object_entities struct {
	ItemName    string
	ItemMovable bool
}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		inventory := map[string][]Object_entities{
			"Items": {
				{ItemName: "M40 Rifle", ItemMovable: true},
				{ItemName: "Barrel", ItemMovable: true},
				{ItemName: "Door", ItemMovable: false},
			},
		}
		tmpl.Execute(w, inventory)
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second) // simulate network
		log.Print("HTMX reqest recived")
		log.Print(r.Header.Get("HX-Request"))

		itemName := r.PostFormValue("itemName")
		itemMovable := r.PostFormValue("itemMovable") == "on" // handle checkbox - actual bool
		fmt.Println(itemName)
		fmt.Println(itemMovable)

		tmpl := template.Must(template.ParseFiles("index.html"))

		// Create a data structure to pass to the template
		/* data := struct {
			Name    string
			Movable bool
		}{
			Name:    itemName,
			Movable: itemMovable,
		} */
		// Define the HTML template string with Go template logic
		//htmlStr := `<li><span> {{ .Name }} </span> <span> {{ if .Movable }}YES{{ else }}NO{{ end }} </span></li>`

		/* 		tmpl, err := template.New("t").Parse(htmlStr)
		   		if err != nil {
		   			http.Error(w, "Template error", http.StatusInternalServerError)
		   			return
		   		} */
		//tmpl.Execute(w, data)

		//use template block
		tmpl.ExecuteTemplate(w, "items-list-element", Object_entities{ItemName: itemName, ItemMovable: itemMovable})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-item/", h2)

	http.ListenAndServe(":8001", nil)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
