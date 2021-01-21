package product

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Continent struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Desc        string `json:"description"`
	Image       string `json:"image"`
}

var Continents = []Continent{
	{Name: "1", DisplayName: "Hello", Desc: "Article Description", Image: "Article Content"},
	{Name: "2", DisplayName: "Hello 2", Desc: "Article Description", Image: "Article Content"},
}

func ListContinents(w http.ResponseWriter, r *http.Request) {
	stringId := strings.TrimPrefix(r.URL.Path, "/product/")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		json.NewEncoder(w).Encode(Continents)
		return
	}

	index := findIndexByContinentId(id)

	if index == -1 {
		return
	}

	json.NewEncoder(w).Encode(Continents[index])
}

func findIndexByContinentId(id int) int {
	for i, p := range Continents {
		if p.Id == id {
			return i
		}
	}

	return -1
}
