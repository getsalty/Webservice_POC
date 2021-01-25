package continent

import (
	"io/ioutil"
	"strings"
)

var dirPathName = "../../common/images/"

func GetImage(path string) (result []byte) {
	name := strings.TrimPrefix(path, "/continent/image/")

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
