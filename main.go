package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	link, err := Get("http://stax.io")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(link)
}

func Get(link string) (string, error) {
	res, err := http.Get(link)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
