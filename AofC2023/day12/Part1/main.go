package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	Row, Col int
}

func main() {
	// input_file_name := "input.txt"
	input_file_name := "test_input.txt"

	read_file, err := os.Open(input_file_name)
	if err != nil {
		panic(err)
	}

	file_scanner := bufio.NewScanner(read_file)
	file_scanner.Split(bufio.ScanLines)

	var text []string

	for file_scanner.Scan() {
		text = append(text, file_scanner.Text())
	}

	// Grab the records and groups
	var damaged_condition_records [][]byte
	var partial_arrangement_groups [][]int
	for _, line := range text {
		split_row := strings.Split(line, " ")
		if len(split_row) != 2 {
			panic("Split row split weird")
		}
		damaged_condition_records = append(damaged_condition_records, []byte(split_row[0]))
		split_condition_records := strings.Split(split_row[1], ",")
		var row_partial_arrangement []int
		for _, num_as_string := range split_condition_records {
			temp_condition, err := strconv.Atoi(num_as_string)
			if err != nil {
				panic(err)
			}
			println(temp_condition)
			row_partial_arrangement = append(row_partial_arrangement, temp_condition)
		}
		partial_arrangement_groups = append(partial_arrangement_groups, row_partial_arrangement)
	}

	println(string(damaged_condition_records[0]))

	println()
	possible_count := 0

	for row_ind := 0; row_ind < len(damaged_condition_records); row_ind++ {
		println(string(damaged_condition_records[row_ind]))
		possible_count += Recursive_Row(damaged_condition_records[row_ind], partial_arrangement_groups[row_ind])
	}
	fmt.Printf("Final arrangement count: %d", possible_count)
}

func Recursive_Row(parts_row []byte, valid_config_row []int) int {
	// Both slices are empty, then it worked
	if len(parts_row) == 0 && len(valid_config_row) == 0 {
		println("yello")
		return 1
	}
	// If only the config input is empty
	if len(valid_config_row) == 0 {
		// Ensure that there are no more confirmed parts left in byte slice input
		for i := 0; i < len(parts_row); i++ {
			if parts_row[i] == '#' {
				println("# found")
				return 0
			}
		}
		println("ello")
		return 1
	}
	// If there are no more parts, but are still configs, then return 0
	if len(parts_row) == 0 {
		println("Parts still left")
		return 0
	}
	count := 0

	var cur_parts_row []byte

	copy(cur_parts_row, parts_row)

	println(cur_parts_row)

	for part_ind := 0; part_ind < len(cur_parts_row); part_ind++ {
		if len(cur_parts_row) < valid_config_row[0] {
			break
		}
		is_valid_config := true
		// Cycle through
		for temp_part_ind := part_ind; temp_part_ind < part_ind+valid_config_row[0]; temp_part_ind++ {
			if parts_row[temp_part_ind] == '.' {
				is_valid_config = false
				break
			}
		}
		if !is_valid_config {
			continue
		}
		r_parts_row := cur_parts_row[part_ind+valid_config_row[0]:]
		if len(r_parts_row) == 1 && r_parts_row[0] == '#' {
			continue
		}

		r_valid_config_row := valid_config_row[1:]

		count += Recursive_Row(r_parts_row, r_valid_config_row)

	}

	fmt.Println("Returning recursive")

	return count
}
