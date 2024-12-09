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

	var disk_map []int

	max_sum := 0
	for i, elem := range input[0] {
		elem_int := int(elem)
		if i%2 == 0 {
			max_sum += elem_int
		}
		disk_map = append(disk_map, elem_int)
	}

	fmt.Println("Max index: ", max_sum)

	checksum := 0
	start_map_index := 0
	end_map_index := len(disk_map) - 1
	end_file_repeats := 0
	end_file_id := 0

	no_gap_index := 0

	for start_map_index <= end_map_index {
		file_id_repeats := disk_map[start_map_index]

		if start_map_index%2 == 0 {
			file_id := getIDFromIndex(start_map_index)
			for i := 0; i < file_id_repeats; i++ {
				fmt.Println("Checksum: ", checksum, " += ", no_gap_index, " * ", file_id)
				checksum += no_gap_index * file_id
				no_gap_index++
			}
			fmt.Println("End of loop")
		} else if start_map_index%2 == 1 {
			for i := 0; i < file_id_repeats; i++ {
				if end_file_repeats == 0 {
					end_file_repeats = disk_map[end_map_index]
					end_file_id = getIDFromIndex(end_map_index)
				}
				fmt.Println("Checksum: ", checksum, " += ", no_gap_index, " * ", end_file_id)
				checksum += no_gap_index * end_file_id
				no_gap_index++
				end_file_repeats--
				if end_file_repeats == 0 {
					end_map_index -= 2
				}
			}
			fmt.Println("End of loop")
		}
		start_map_index++
	}

	fmt.Println("Final: ", checksum)
}

func getIDFromIndex(i int) int {
	fmt.Println("i: ", i)
	return i / 2
}

func getInputAsLines() ([]string, error) {
	// Read in files
	f, err := os.Open("test.txt")
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
