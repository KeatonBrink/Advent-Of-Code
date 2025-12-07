package main

import (
	"bufio"
	"fmt"
	"os"
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
	total_splits := 0
	// Location of active lasers
	var cur_lasers []int
	for _, line := range input {
		var next_lasers []int
		if len(cur_lasers) == 0 {
			for ind, char := range line {
				if char == 'S' {
					next_lasers = []int{ind}
					break
				}
			}
		} else {
			for _, ind := range cur_lasers {
				if line[ind] == '^' {
					total_splits += 1
					if ind > 0 && !isPositionInSlice(ind-1, next_lasers) {
						next_lasers = append(next_lasers, ind-1)
					}
					if ind < len(line)-1 && !isPositionInSlice(ind+1, next_lasers) {
						next_lasers = append(next_lasers, ind+1)
					}
				} else if line[ind] == '.' && !isPositionInSlice(ind, next_lasers) {
					next_lasers = append(next_lasers, ind)
				}
			}
		}
		cur_lasers = next_lasers
	}
	fmt.Println(total_splits)
}

func isPositionInSlice(target_ind int, cur_slice []int) bool {
	for _, cur_ind := range cur_slice {
		if cur_ind == target_ind {
			return true
		}
	}
	return false
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
