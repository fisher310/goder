package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, err := os.Open("inpractice/configuration/config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var conf configuration = configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println(conf.Path)
}
