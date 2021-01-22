package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/getsalty/Webservice_POC/go/src/continent"
)

func main() {
	http.HandleFunc("/continent", func(w http.ResponseWriter, r *http.Request) {
		continent := continent.ListContinents(r.URL.Path)
		if continent != nil {
			json.NewEncoder(w).Encode(continent)
		}
	})
	http.HandleFunc("/continent/", func(w http.ResponseWriter, r *http.Request) {
		continent := continent.ListContinents(r.URL.Path)
		if continent != nil {
			json.NewEncoder(w).Encode(continent)
		}
	})
	http.HandleFunc("/continent/image/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(continent.GetImage(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":10000", nil))
}
