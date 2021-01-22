package reader

import (
	"io/ioutil"
)

func ReadFile(pathname string) ([]byte, error) {
	byteValue, err := ioutil.ReadFile(pathname)
	return byteValue, err
}
