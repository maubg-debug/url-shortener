package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Respuesta struct {
	status bool   `json:"estado"`
	codigo string `json:"codigo"`
}

func fetch(url string) {

	urlN := "https://maubot.maucode.com/api/redir/crear?url=" + url

	req, err := http.NewRequest("GET", urlN, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var record Respuesta

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	return record.codigo
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

		codigo := fetch(data)

		fmt.Printf(codigo)
	})

	// http.HandleFunc("/", index)
	fmt.Printf("Escuchando en el puero :8000\n")
	http.ListenAndServe(":8000", nil)
}
