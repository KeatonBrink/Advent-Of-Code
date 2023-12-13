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
			row_partial_arrangement = append(row_partial_arrangement, temp_condition)
		}
		partial_arrangement_groups = append(partial_arrangement_groups, row_partial_arrangement)
	}

	possible_count := 0

	for row_ind := 0; row_ind < len(damaged_condition_records); row_ind++ {
		new_count := Recursive_Row(damaged_condition_records[row_ind], partial_arrangement_groups[row_ind])
		fmt.Printf("\n\n\n\nLine %s, Count: %d\n\n\n\n\n", damaged_condition_records[row_ind], new_count)
		possible_count += new_count
	}
	fmt.Printf("Final arrangement count: %d\n", possible_count)
}

func Recursive_Row(parts_row []byte, valid_config_row []int) int {
	//println("Starting recursive row")
	//println(string(parts_row))
	// Both slices are empty, then it worked
	if len(parts_row) == 0 {
		if len(valid_config_row) == 0 {
			return 1
		}
		return 0
	}

	if len(valid_config_row) == 0 {
		for i := 0; i < len(parts_row); i++ {
			if parts_row[i] == '#' {
				return 0
			}
		}
		return 1
	}

	count := 0

	for part_ind := 0; part_ind < len(parts_row); part_ind++ {
		//println("In for loop")
		if len(parts_row[part_ind:]) < valid_config_row[0] {
			//fmt.Printf("Curpart len %d, config val %d\n", len(parts_row), valid_config_row[0])
			break
		}
		//fmt.Printf("%d %s %d\n", len(parts_row), string(parts_row), valid_config_row[0])
		is_valid_config := true
		// Cycle through
		for temp_part_ind := part_ind; temp_part_ind < part_ind+valid_config_row[0]; temp_part_ind++ {
			if parts_row[temp_part_ind] == '.' {
				//println("Is not valid config")
				is_valid_config = false
				break
			}
		}
		//println("Finish cycle")
		if !is_valid_config {
			//println("Is not valid config")
			continue
		}
		r_parts_row := parts_row[part_ind+valid_config_row[0]:]
		if len(r_parts_row) > 0 {
			if r_parts_row[0] == '#' {
				//println("Found #")
				continue
			} else {
				r_parts_row = r_parts_row[1:]
			}
		}

		r_valid_config_row := valid_config_row[1:]

		count += Recursive_Row(r_parts_row, r_valid_config_row)

		if parts_row[part_ind] == '#' {
			break
		}
	}

	// if parts_row[0] == '.' || parts_row[0] == '?' {
	// 	count += Recursive_Row(parts_row[1:], valid_config_row)
	// }

	// if parts_row[0] == '#' || parts_row[0] == '?' {
	// 	does_not_contain_dot := true
	// 	for i := 0; i < valid_config_row[0]; i++ {
	// 		if parts_row[i] == '.' {
	// 			does_not_contain_dot = false
	// 		}
	// 	}
	// 	if valid_config_row[0] <= len(parts_row) && does_not_contain_dot && (valid_config_row[0] == len(parts_row) || parts_row[valid_config_row[0]] != '#') {

	// 		count += Recursive_Row(parts_row[valid_config_row[0]+1:], valid_config_row[1:])
	// 	}
	// }

	//fmt.Println("Returning recursive")

	return count
}
