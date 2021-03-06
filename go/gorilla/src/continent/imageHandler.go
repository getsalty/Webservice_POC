package continent

import (
	"io/ioutil"
)

var dirPathName = "../../common/images/"

func GetImage(name string) (result []byte) {
	imagePath := (dirPathName + name + ".svg")

	byteValue, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return result
	}

	return byteValue
}
