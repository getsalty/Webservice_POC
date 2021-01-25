package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/getsalty/Webservice_POC/go/gorilla/src/continent"
)

func main() {
	http.HandleFunc("/continent", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		var continentList []continent.Continent
		continent.ListContinents(r.URL.Path, &continentList)
		if continentList != nil {
			json.NewEncoder(w).Encode(continentList)
		}
	})
	http.HandleFunc("/continent/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		var continentObject []continent.Continent
		continent.ListContinents(r.URL.Path, &continentObject)
		if continentObject != nil {
			json.NewEncoder(w).Encode(continentObject)
		}
	})
	http.HandleFunc("/continent/image/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		w.Write(continent.GetImage(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":10000", nil))
}
