package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/getsalty/Webservice_POC/go/native/src/continent"
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
		var continentObject []continent.Continent
		continent.ListContinents(r.URL.Path, &continentObject)

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

	})

	http.HandleFunc("/continent/image/", func(w http.ResponseWriter, r *http.Request) {
		image := continent.GetImage(r.URL.Path)

		if image == nil {
			w.WriteHeader(404)
			w.Write([]byte(`Could not find image with desired name.`))
		} else {
			w.Header().Add("Content-Type", "image/svg+xml")
			w.Write(image)
		}
	})

	log.Fatal(http.ListenAndServe(":4000", nil))
}
