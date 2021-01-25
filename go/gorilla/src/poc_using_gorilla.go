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
		var continentList []continent.Continent
		continent.ListContinents("", &continentList)

		if continentList != nil {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(continentList)
		}
	}).Methods("GET")

	router.HandleFunc("/continent/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var continentObject []continent.Continent
		continent.ListContinents(vars["name"], &continentObject)

		if continentObject == nil {
			w.WriteHeader(404)
			w.Write([]byte(`Could not find continent with desired name.`))
		} else {
			w.Header().Add("Content-Type", "application/json")
			if len(continentObject) == 1 {
				json.NewEncoder(w).Encode(continentObject[0])
			} else {
				json.NewEncoder(w).Encode(continentObject)
			}
		}
	}).Methods("GET")

	router.HandleFunc("/continent/image/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		image := continent.GetImage(vars["name"])

		if image == nil {
			w.WriteHeader(404)
			w.Write([]byte(`Could not find image with desired name.`))
		} else {
			w.Header().Add("Content-Type", "image/svg+xml")
			w.Write(image)
		}
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}
