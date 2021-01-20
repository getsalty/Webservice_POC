package product

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Image string `json:"image"`
}

var Products = []Product{
	Product{Id: 1, Title: "Hello", Desc: "Article Description", Image: "Article Content"},
	Product{Id: 2, Title: "Hello 2", Desc: "Article Description", Image: "Article Content"},
}

func listProducts(w http.ResponseWriter, r *http.Request) {
	stringId := strings.TrimPrefix(r.URL.Path, "/product/")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		json.NewEncoder(w).Encode(Products)
		return
	}

	index := findIndexByProductId(id)

	if index == -1 {
		return
	}

	json.NewEncoder(w).Encode(Products[index])
}

func findIndexByProductId(id int) int {
	for i, p := range Products {
		if p.Id == id {
			return i
		}
	}

	return -1
}
