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

	gardens_counted := make([][]bool, len(input))
	for ri, line := range input {
		gardens_counted[ri] = make([]bool, len(input))
		for _, elem := range line {
			_, ok := plant_gardens[elem]
			if !ok {
				plant_gardens[elem] = Plant{perimeter: 0, area: 0}
			}
		}
	}

	total_price := 0
	for ri, line := range input {
		for ci, elem := range line {
			if gardens_counted[ri][ci] {
				continue
			}
			temp_perim, temp_area := recurseGarden(ri, ci, elem, gardens_counted, input)

			total_price += temp_perim * temp_area

			// temp_plant := plant_gardens[elem]
			// if ri == 0 || (rune(input[ri-1][ci]) != elem) {
			// 	temp_plant.perimeter++
			// }
			// if ci == len(line)-1 || (rune(input[ri][ci+1]) != elem) {
			// 	temp_plant.perimeter++
			// }
			// if ri == len(input)-1 || (rune(input[ri+1][ci]) != elem) {
			// 	temp_plant.perimeter++
			// }
			// if ci == 0 || (rune(input[ri][ci-1]) != elem) {
			// 	temp_plant.perimeter++
			// }
			// temp_plant.area++
			// plant_gardens[elem] = temp_plant
		}
	}

	// total_price := 0

	// for plant_name, plant_stats := range plant_gardens {
	// 	plant_price := plant_stats.area * plant_stats.perimeter
	// 	fmt.Println("Plant, perimeter, area, cost ", plant_name, plant_stats.perimeter, plant_stats.area, plant_price)
	// 	total_price += plant_price
	// }

	fmt.Println("Final: ", total_price)
}

func recurseGarden(ri, ci int, elem rune, is_visited_gardens [][]bool, input []string) (int, int) {
	is_visited_gardens[ri][ci] = true
	cur_perim := 0
	cur_area := 1
	if ri == 0 || (rune(input[ri-1][ci]) != elem) {
		cur_perim++
	} else if ri > 0 && !is_visited_gardens[ri-1][ci] {
		temp_perim, temp_area := recurseGarden(ri-1, ci, elem, is_visited_gardens, input)
		cur_perim += temp_perim
		cur_area += temp_area

	}
	if ci == len(input[ri])-1 || (rune(input[ri][ci+1]) != elem) {
		cur_perim++
	} else if ci < len(input[ri])-1 && !is_visited_gardens[ri][ci+1] {
		temp_perim, temp_area := recurseGarden(ri, ci+1, elem, is_visited_gardens, input)
		cur_perim += temp_perim
		cur_area += temp_area
	}
	if ri == len(input)-1 || (rune(input[ri+1][ci]) != elem) {
		cur_perim++
	} else if ri < len(input)-1 && !is_visited_gardens[ri+1][ci] {
		temp_perim, temp_area := recurseGarden(ri+1, ci, elem, is_visited_gardens, input)
		cur_perim += temp_perim
		cur_area += temp_area
	}
	if ci == 0 || (rune(input[ri][ci-1]) != elem) {
		cur_perim++
	} else if ci > 0 && !is_visited_gardens[ri][ci-1] {
		temp_perim, temp_area := recurseGarden(ri, ci-1, elem, is_visited_gardens, input)
		cur_perim += temp_perim
		cur_area += temp_area
	}
	return cur_perim, cur_area
}

func getInputAsLines() ([]string, error) {
	// Read in files
	// f, err := os.Open("test2.txt")
	// f, err := os.Open("test.txt")
	f, err := os.Open("input.txt")
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
