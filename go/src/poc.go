package main

import (
	"log"
	"net/http"

	"github.com/getsalty/Webservice_POC/go/src/continent"
)

func main() {
	http.HandleFunc("/continent", continent.ListContinents)
	http.HandleFunc("/continent/", continent.ListContinents)
	http.HandleFunc("/continent/image/", continent.GetImage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
