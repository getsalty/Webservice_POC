package continent

import (
	"net/http"
	"strings"

	"github.com/getsalty/Webservice_POC/go/src/reader"
)

var dirPathName = "../common/images/"

func GetImage(w http.ResponseWriter, r *http.Request) {

	name := strings.TrimPrefix(r.URL.Path, "/continent/image/")

	if len(name) == 0 {
		return
	}

	imagePath := (dirPathName + name + ".svg")

	byteValue, err := reader.ReadFile(imagePath)
	if err != nil {
		return
	}

	w.Write(byteValue)
}
