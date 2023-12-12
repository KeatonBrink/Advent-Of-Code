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

	for row_ind, row_parts := range damaged_condition_records {
		row_arrangement := partial_arrangement_groups[row_ind]
		for start_ind := 0; start_ind < len(row_parts)
	}
	fmt.Printf("Final arrangement count: %d", possible_count)
}
