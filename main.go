package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
)

const (
	defaultFilename = "modules.json"
)

func parseFile(file string) ([]int, error) {
	var modules []int
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &modules)
	return modules, err
}

func calculateFuel(mass float64) float64 {
	return math.Floor(mass/3) - 2
}

func main() {
	modules, err := parseFile(defaultFilename)
	if err != nil {
		fmt.Println(err)
	}

	var totalFuelNeeded float64

	for _, i := range modules {
		totalFuelNeeded += calculateFuel(float64(i))
	}

	fmt.Printf("Total fuel needed: %f\n", totalFuelNeeded)
}
