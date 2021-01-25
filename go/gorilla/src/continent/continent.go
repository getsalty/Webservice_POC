package continent

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

var pathName = "../common/data.json"

type Continent struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Desc        string `json:"description"`
}

type Data struct {
	Continents []Continent `json:"continents"`
}

func ListContinents(path string, result *[]Continent) {
	byteValue, err := ioutil.ReadFile(pathName)
	if err != nil {
		return
	}

	var data Data
	json.Unmarshal(byteValue, &data)

	name := strings.TrimPrefix(path, "/continent/")

	if len(name) == 0 || name == "/continent" {
		*result = data.Continents
		return
	}

	index := findIndexByName(name, data.Continents)

	if index == -1 {
		return
	}

	data.Continents = []Continent{data.Continents[index]}
	*result = data.Continents
}

func findIndexByName(name string, dataset []Continent) int {
	for i, p := range dataset {
		if p.Name == name {
			return i
		}
	}

	return -1
}
