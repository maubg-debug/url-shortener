package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func fetch(url string) {

	urlN := "https://maubot.maucode.com/api/redir/crear?url=" + url

	req, err := http.NewRequest(http.MethodGet, urlN, nil)
	if err != nil {
		log.Fatal(err)
	}

	return req
}

func main() {

	tmpl := template.Must(template.ParseFiles("web/index.html"))
	// fs := http.FileServer(http.Dir("./assets"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		data := r.FormValue("url")

		fData := fetch(data)

		fmt.Printf(fData)
	})

	// http.HandleFunc("/", index)
	fmt.Printf("Escuchando en el puero :8000\n")
	http.ListenAndServe(":8000", nil)
}
