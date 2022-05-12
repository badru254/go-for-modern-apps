package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

// Config holds configuration
type Config struct {
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.Wrap(err, "can't open configuration file")
	}

	defer file.Close()

	cfg := &Config{}

	return cfg, nil

}

func setupLogging() {

	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 064)

	if err != nil {
		return
	}

	log.SetOutput(out)
}

func main() {
	// setupLogging()
	// _, err := readConfig("/path/to/config.toml")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	// 	log.Printf("error: %+v", err)
	// 	os.Exit(1)
	// }
	arr := []int{23, 45, 67}
	arrItem, err := safeValue(arr, 1)

	fmt.Printf("Error : %v \n", err)
	fmt.Printf("val : %v \n", arrItem)

}

//
func safeValue(vals []int, index int) (n int, err error) {

	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	return vals[index], nil
}
