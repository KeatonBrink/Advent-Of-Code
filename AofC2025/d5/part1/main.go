package main

import (
	"bufio"
	"fmt"
	"os"
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
		fresh_ranges = append(fresh_ranges, fresh_range)
		cur_ind += 1
		cur_line = input[cur_ind]
	}
	cur_ind += 1
	cur_line = input[cur_ind]
	total := 0
	for true {
		cur_ingredient_id, err := strconv.Atoi(cur_line)
		if err != nil {
			fmt.Println(err)
			return
		}
		is_fresh := false
		for _, fresh_range := range fresh_ranges {
			if cur_ingredient_id >= fresh_range[0] && cur_ingredient_id <= fresh_range[1] {
				is_fresh = true
				break
			}
		}
		if is_fresh {
			total += 1
		}
		cur_ind += 1
		if cur_ind >= len(input) {
			break
		}
		cur_line = input[cur_ind]
	}
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
