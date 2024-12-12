package main

import (
	"bufio"
	"fmt"
	"os"
)

type Plant struct {
	perimeter int
	area      int
}

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)

	plant_gardens := make(map[rune]Plant)

	for _, line := range input {
		for _, elem := range line {
			_, ok := plant_gardens[elem]
			if !ok {
				plant_gardens[elem] = Plant{perimeter: 0, area: 0}
			}
		}
	}

	for ri, line := range input {
		for ci, elem := range line {
			temp_plant := plant_gardens[elem]
			if ri == 0 || (rune(input[ri-1][ci]) != elem) {
				temp_plant.perimeter++
			}
			if ci == len(line)-1 || (rune(input[ri][ci+1]) != elem) {
				temp_plant.perimeter++
			}
			if ri == len(input)-1 || (rune(input[ri+1][ci]) != elem) {
				temp_plant.perimeter++
			}
			if ci == 0 || (rune(input[ri][ci-1]) != elem) {
				temp_plant.perimeter++
			}
			temp_plant.area++
			plant_gardens[elem] = temp_plant
		}
	}

	total_price := 0

	for plant_name, plant_stats := range plant_gardens {
		plant_price := plant_stats.area * plant_stats.perimeter
		fmt.Println("Plant, perimeter, area, cost ", plant_name, plant_stats.perimeter, plant_stats.area, plant_price)
		total_price += plant_price
	}

	fmt.Println("Final: ", total_price)
}

func getInputAsLines() ([]string, error) {
	// Read in files
	f, err := os.Open("test.txt")
	// f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Create scanner
	file_scanner := bufio.NewScanner(f)
	file_scanner.Split(bufio.ScanLines)

	// Get lines
	var text []string
	for file_scanner.Scan() {
		text = append(text, file_scanner.Text())
	}
	return text, nil
}
