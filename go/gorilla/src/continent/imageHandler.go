package continent

import (
	"io/ioutil"
)

var dirPathName = "../../common/images/"

func GetImage(name string) (result []byte) {

	if len(name) == 0 {
		return result
	}

	imagePath := (dirPathName + name + ".svg")

	byteValue, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return result
	}

	return byteValue
}
