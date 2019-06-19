package fshare

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func PointerString(input string) *string {
	return &input
}

func logBody(data *io.ReadCloser) {
	dataBytes, err := ioutil.ReadAll(*data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataBytes))
	*data = ioutil.NopCloser(bytes.NewBuffer(dataBytes))
}
