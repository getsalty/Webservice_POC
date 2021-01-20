package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/getsalty/Webservice_POC/go/src/product"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/product", product.ListProducts)
	http.HandleFunc("/product/", product.ListProducts)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
