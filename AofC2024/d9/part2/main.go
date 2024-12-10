package main

import (
	"bufio"
	"fmt"
	"os"
)

type EmptySpace struct {
	index int
	size  int
}

type FileSpace struct {
	index int
	size  int
	id    int
}

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

	// Make a list of empty space structs
	// Low key want to just make everything into strings and separate them
	var ordered_empty_space []EmptySpace
	found_empty_space := false
	var temp_empty_space EmptySpace
	for expanded_notation_index, spot := range expanded_notation {
		if spot == -1 {
			if !found_empty_space {
				found_empty_space = true
				temp_empty_space = EmptySpace{index: expanded_notation_index, size: 1}
			} else {
				temp_empty_space.size++
			}
		} else if spot > -1 {
			if found_empty_space {
				found_empty_space = false
				ordered_empty_space = append(ordered_empty_space, temp_empty_space)
			}
		}
	}

	fmt.Println(ordered_empty_space)

	var ordered_file_space []FileSpace
	found_file := false
	temp_file := FileSpace{index: -2, size: -2, id: -2}
	for expanded_notation_index, spot := range expanded_notation {
		if spot > -1 {
			if !found_file {
				found_file = true
				temp_file = FileSpace{index: expanded_notation_index, size: 1, id: spot}
			} else {
				if temp_file.id == spot {
					temp_file.size++
				} else {
					ordered_file_space = append([]FileSpace{temp_file}, ordered_file_space...)
					temp_file = FileSpace{index: expanded_notation_index, size: 1, id: spot}
				}
			}
		} else if spot == -1 {
			if found_file {
				found_file = false
				ordered_file_space = append([]FileSpace{temp_file}, ordered_file_space...)
				temp_file = FileSpace{index: -2, size: -2, id: -2}
			}
		}
	}
	ordered_file_space = append([]FileSpace{temp_file}, ordered_file_space...)

	fmt.Println(ordered_file_space)

	// Start swapping
	// Swapping actually occurs in expanded notation
	for file_space_index, file := range ordered_file_space {
		for empty_space_index, empty_space := range ordered_empty_space {
			if empty_space.size >= file.size && empty_space.index < file.index {
				expanded_notation_index := empty_space.index
				expanded_notation_file_index := file.index
				// temp_file := file
				for file_rep := 0; file_rep < file.size; {
					expanded_notation[expanded_notation_index] = file.id
					expanded_notation[expanded_notation_file_index] = -1
					file.size--
					file.index++
					empty_space.size--
					empty_space.index++
					expanded_notation_index++
					expanded_notation_file_index++
				}
				// for file_rep := 0; file_rep < file.size; file_rep++ {
				// 	expanded_notation[temp_file.index+file_rep] = -1
				// }
				ordered_empty_space[empty_space_index] = empty_space
				ordered_file_space[file_space_index] = file
				break
			}
		}
	}

	fmt.Println(expanded_notation)

	checksum := 0
	// Start Counting
	for spot_index, spot := range expanded_notation {
		if spot == -1 {
			continue
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
