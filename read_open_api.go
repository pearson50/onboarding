package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
		Paths struct {
			Pets struct {
				Get struct {
					Description string `yaml:"description"`
				}`yaml:"get"`
			}`yaml:"/pets"`
		}`yaml:"paths"`
}

func main() {
	fmt.Println("Parsing YAML file")

	var fileName string
	flag.StringVar(&fileName, "f", "petstore.yaml", "YAML file to parse.")
	flag.Parse()

	if fileName == "" {
		fmt.Println("Please provide yaml file by using -f option")
		return
	}

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}

	var yamlConfig config
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	fmt.Printf("Result: %v\n", yamlConfig)
}


