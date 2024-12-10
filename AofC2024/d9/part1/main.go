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

	var disk_map []int

	for _, elem := range input[0] {
		elem_int := int(elem) - '0'
		disk_map = append(disk_map, elem_int)
	}

	// Expand the disk, -1 is .

	var expanded_notation []int

	for disk_map_index, elem := range disk_map {
		var id int
		if disk_map_index%2 == 0 {
			id = getIDFromIndex(disk_map_index)
		} else {
			id = -1
		}
		for cur_rep := 0; cur_rep < elem; cur_rep++ {
			expanded_notation = append(expanded_notation, id)
		}
	}

	end_index := len(expanded_notation) - 1

	// Start swapping
	for spot_index, spot := range expanded_notation {
		if spot == -1 {
			if spot_index >= end_index {
				break
			}
			expanded_notation[spot_index] = expanded_notation[end_index]
			expanded_notation[end_index] = -1
			for ; end_index > 0 && expanded_notation[end_index] == -1; end_index-- {
			}
		}
	}

	fmt.Println(expanded_notation)

	checksum := 0
	// Start Counting
	for spot_index, spot := range expanded_notation {
		if spot == -1 {
			break
		}
		checksum += (spot_index * spot)
	}

	fmt.Println("Final: ", checksum)
}

func getIDFromIndex(i int) int {
	// fmt.Println("i: ", i)
	return i / 2
}

func getInputAsLines() ([]string, error) {
	// Read in files
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
