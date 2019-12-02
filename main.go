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

func calculateFuel(mass int) int {
	return int(math.Floor(float64(mass)/3) - 2)
}

func calculateExtraFuel(mainFuel int) int {
	remainderFuel := mainFuel
	var extraFuel int
	for remainderFuel >= 0 {
		remainderFuel = calculateFuel(remainderFuel)
		if remainderFuel > 0 {
			extraFuel += remainderFuel
		}
	}

	return extraFuel
}

func partOneTask() int {
	modules, err := parseFile(defaultFilename)
	if err != nil {
		fmt.Println(err)
	}

	var totalModuleFuelNeeded int

	for _, i := range modules {
		totalModuleFuelNeeded += calculateFuel(int(i))
	}

	return totalModuleFuelNeeded
}

func partTwoTask() int {
	modules, err := parseFile(defaultFilename)
	if err != nil {
		fmt.Println(err)
	}

	var totalModuleFuelNeeded int

	for _, i := range modules {
		totalModuleFuelNeeded += calculateExtraFuel(int(i))
	}

	return totalModuleFuelNeeded
}

func main() {
	fmt.Printf("[Day 1] Total fuel needed: %v\n", partOneTask())
	fmt.Printf("[Day 2] Total fuel needed: %v\n", partTwoTask())

	// 4847868
}
