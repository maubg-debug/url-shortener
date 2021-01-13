package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Respuesta struct {
	codigo string `json:"codigo"`
	estado int    `json:"estado"`
}

func main() {

	tmpl := template.Must(template.ParseFiles("web/index.html"))
	fs := http.FileServer(http.Dir("./assets"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, fs)
			return
		}

		data := r.FormValue("url")

		url := "https://maubot.maucode.com/api/redir/crear?url=" + data

		spaceClient := http.Client{
			Timeout: time.Second * 2, // Timeout after 2 seconds
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		res, getErr := spaceClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		fmt.Println(body)

		people1 := Respuesta{}
		jsonErr := json.Unmarshal(body, &people1)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		fmt.Println(people1.codigo)

		// codigof := "https://maubot.maucode.com/api/redir?codigo=" + record.codigo

		// fmt.Println(codigof)
	})

	// http.HandleFunc("/", index)
	fmt.Printf("Escuchando en el puero :8000\n")
	http.ListenAndServe(":8000", nil)
}
