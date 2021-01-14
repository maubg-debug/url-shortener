package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
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

		data := r.FormValue("url")

		url := "https://maubot.maucode.com/api/redir/crear?url=" + data

		response, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
			// jsonValue, _ := json.Marshal(data)

			k := make([]string, len(data))
			i := 0
			for s := range data {
				k[i] = fmt.Sprint(s)
				i++
			}
			fmt.Println(k)
		}

		// codigof := "https://maubot.maucode.com/api/redir?codigo=" + record.codigo
	})

	fmt.Printf("Escuchando en el puero :5000\n")
	http.ListenAndServe(":5000", nil)
}
