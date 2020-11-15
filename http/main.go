package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Some error ocurred")
		os.Exit(1)
	}

	io.Copy(logWriter{}, response.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
