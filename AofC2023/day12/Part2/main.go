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

var PROBLEM_MAP = make(map[string]int)
var DUPS_FOUND = 0

func main() {
	input_file_name := "input.txt"
	// input_file_name := "test_input.txt"

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
	var damaged_condition_records []string
	var partial_arrangement_groups [][]int
	for ind, line := range text {
		split_row := strings.Split(line, " ")
		if len(split_row) != 2 {
			panic("Split row split weird")
		}
		damaged_condition_records = append(damaged_condition_records, split_row[0])
		for i := 0; i < 4; i++ {
			damaged_condition_records[ind] += "?"
			damaged_condition_records[ind] += split_row[0]
		}
		split_condition_records := strings.Split(split_row[1], ",")
		var row_partial_arrangement []int
		for i := 0; i < 5; i++ {
			for _, num_as_string := range split_condition_records {
				temp_condition, err := strconv.Atoi(num_as_string)
				if err != nil {
					panic(err)
				}
				row_partial_arrangement = append(row_partial_arrangement, temp_condition)
			}
		}
		for i := 0; i < len(row_partial_arrangement); i++ {
			fmt.Printf("%d ", row_partial_arrangement[i])
		}
		println()
		partial_arrangement_groups = append(partial_arrangement_groups, row_partial_arrangement)
	}

	possible_count := 0

	for row_ind := 0; row_ind < len(damaged_condition_records); row_ind++ {
		new_count := Recursive_Row(damaged_condition_records[row_ind], partial_arrangement_groups[row_ind])
		fmt.Printf("\n\n\n\nLine %s, Count: %d\n\n\n\n\n", damaged_condition_records[row_ind], new_count)
		possible_count += new_count
	}
	fmt.Printf("Length of Memo: %d\n", len(PROBLEM_MAP))
	fmt.Printf("Total dups found: %d\n", DUPS_FOUND)
	fmt.Printf("Final arrangement count: %d\n", possible_count)
}

func Recursive_Row(parts_row string, valid_config_row []int) int {
	problem_as_string := ProblemPartsToString(parts_row, valid_config_row)
	val, ok := PROBLEM_MAP[problem_as_string]
	if ok {
		DUPS_FOUND++
		return val
	}
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

	if parts_row[0] == '.' || parts_row[0] == '?' {
		count += Recursive_Row(parts_row[1:], valid_config_row)
	}

	if parts_row[0] == '#' || parts_row[0] == '?' {
		does_not_contain_dot := true
		for i := 0; i < valid_config_row[0] && i < len(parts_row); i++ {
			if parts_row[i] == '.' {
				does_not_contain_dot = false
			}
		}
		if valid_config_row[0] <= len(parts_row) && does_not_contain_dot && (valid_config_row[0] == len(parts_row) || parts_row[valid_config_row[0]] != '#') {
			if len(parts_row) == valid_config_row[0] {
				count += Recursive_Row(parts_row[valid_config_row[0]:], valid_config_row[1:])
			} else {

				count += Recursive_Row(parts_row[valid_config_row[0]+1:], valid_config_row[1:])
			}
		}
	}

	PROBLEM_MAP[problem_as_string] = count

	//fmt.Println("Returning recursive")

	return count
}

func ProblemPartsToString(parts_row string, valid_config_row []int) string {
	return parts_row + " " + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(valid_config_row)), ","), "[]")
}
