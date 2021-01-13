package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	tmpl := template.Must(template.ParseFiles("web/index.html"))
	fs := http.FileServer(http.Dir("./assets"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, fs)
			return
		}

		type Respuesta struct {
			codigo string `json:"codigo"`
			estado int    `json:"estado"`
		}

		data := r.FormValue("url")

		urlN := "https://maubot.maucode.com/api/redir/crear?url=" + data

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
			return
		}

		codigo := ""

		if record.estado == 200 {
			codigo = "https://maubot.maucode.com/api/redir?codigo=" + record.codigo
		} else {
			codigo = ""
		}

		fmt.Printf(codigo)
	})

	// http.HandleFunc("/", index)
	fmt.Printf("Escuchando en el puero :8000\n")
	http.ListenAndServe(":8000", nil)
}
