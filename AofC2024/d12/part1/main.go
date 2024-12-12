package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)

	gardens_counted := make([][]bool, len(input))
	for ri := 0; ri < len(input); ri++ {
		gardens_counted[ri] = make([]bool, len(input))
	}

	total_price := 0
	for ri, line := range input {
		for ci, elem := range line {
			if gardens_counted[ri][ci] {
				continue
			}
			temp_perim, temp_area := recurseGarden(ri, ci, elem, gardens_counted, input)

			total_price += temp_perim * temp_area
		}
	}

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
