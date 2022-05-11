package main

import (
	"fmt"
	"net/http"
)

func main() {

	res, err := contentType("https://www.google.com/")
	if err == nil {
		fmt.Printf("Content Type Header = %v\n", res)
	} else {
		fmt.Printf("error occured %v\n", res)
	}

}

//returns content type header by making get request
func contentType(url string) (string, error) {

	res, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	header := res.Header.Get("Content-Type")

	if header == "" {
		fmt.Println("Could not fetch the header")
	}

	return header, nil
}
