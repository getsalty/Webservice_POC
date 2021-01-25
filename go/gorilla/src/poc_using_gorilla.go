package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/getsalty/Webservice_POC/go/gorilla/src/continent"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/continent", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		var continentList []continent.Continent
		continent.ListContinents("", &continentList)
		if continentList != nil {
			json.NewEncoder(w).Encode(continentList)
		}
	})

	router.HandleFunc("/continent/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		vars := mux.Vars(r)

		var continentObject []continent.Continent
		continent.ListContinents(vars["name"], &continentObject)
		if continentObject != nil {
			json.NewEncoder(w).Encode(continentObject)
		}
	})

	router.HandleFunc("/continent/image/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		vars := mux.Vars(r)

		w.Write(continent.GetImage(vars["name"]))
	})

	log.Fatal(http.ListenAndServe(":5000", router))
}
