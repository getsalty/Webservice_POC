package continent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/getsalty/Webservice_POC/go/src/reader"
)

var pathName = "../common/data.json"

type Continent struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Desc        string `json:"description"`
	Image       string `json:"image"`
}

type Data struct {
	Continents []Continent `json:"continents"`
}

func ListContinents(w http.ResponseWriter, r *http.Request) {
	byteValue, err := reader.ReadFile(pathName)
	if err != nil {
		return
	}

	data := convertToData(byteValue)

	printData(data)

	name := strings.TrimPrefix(r.URL.Path, "/continent/")

	if len(name) == 0 || name == "/continent" {
		json.NewEncoder(w).Encode(data.Continents)
		return
	}

	index := findIndexByName(name, data.Continents)

	if index == -1 {
		return
	}

	json.NewEncoder(w).Encode(data.Continents[index])
}

func findIndexByName(name string, dataset []Continent) int {
	for i, p := range dataset {
		if p.Name == name {
			return i
		}
	}

	return -1
}

func convertToData(byteArray []byte) Data {
	var result Data
	json.Unmarshal(byteArray, &result)

	return result
}

func printData(data Data) {
	for i, p := range data.Continents {
		fmt.Println("     ", i, ":")
		fmt.Println("            Name:", p.Name)
		fmt.Println("            DisplayName:", p.DisplayName)
		fmt.Println("            Desc:", p.Desc)
		fmt.Println("            Image:", p.Image)
	}
}
