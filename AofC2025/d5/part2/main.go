package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file_name := "input.txt"
	args := os.Args[1:]
	if len(args) >= 1 {
		file_name = args[0]
	}
	input, err := getInputAsLines(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	var fresh_ranges [][]int
	cur_ind := 0
	cur_line := input[cur_ind]
	for len(cur_line) > 0 {
		result_slice := strings.Split(cur_line, "-")
		start_val, err := strconv.Atoi(result_slice[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		end_val, err := strconv.Atoi(result_slice[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		fresh_range := []int{start_val, end_val}

		if len(fresh_ranges) == 0 {
			fresh_ranges = append(fresh_ranges, fresh_range)
		} else {
			// Sorting along the ways
			inserted := false
			for ind, temp_range := range fresh_ranges {
				if temp_range[0] > fresh_range[0] {
					fresh_ranges = slices.Insert(fresh_ranges, ind, fresh_range)
					inserted = true
					break
				}
			}
			if !inserted {
				fresh_ranges = append(fresh_ranges, fresh_range)
			}
		}
		cur_ind += 1
		cur_line = input[cur_ind]
	}
	// fmt.Println(fresh_ranges)

	max_value := -1
	total := 0
	// Options for overlapping ranges:
	// Ranges do not overlap
	// 	[ rangeA ] [ rangeB ]
	// Ranges overlap in some part
	// 	This includes if rangeA == rangeB
	// 	[ rangeA	[ overlapAB ]	rangeB ]
	// Ranges encapsulate each other
	// 	[ rangeA	[ rangeB ]	]
	// Only works because of the previous sorting process
	for _, cur_range := range fresh_ranges {
		// Previous range encapsulated cur_range
		if max_value >= cur_range[1] {
			continue
			// Previous range included part of cur_range
		} else if max_value >= cur_range[0] {
			total += cur_range[1] - max_value
			// Previous range ended before start of current range
		} else {
			total += cur_range[1] - cur_range[0] + 1
		}
		max_value = cur_range[1]
	}

	fmt.Println(fresh_ranges)
	// fmt.Println(results)
	fmt.Println(total)
}

func getInputAsLines(file_name string) ([]string, error) {
	// Read in files
	f, err := os.Open(file_name)
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
