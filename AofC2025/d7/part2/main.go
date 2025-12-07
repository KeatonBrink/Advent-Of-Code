package main

import (
	"bufio"
	"fmt"
	"os"
)

type Laser struct {
	Location int
	Total    int
}

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
	total_splits := 1
	// Location of active lasers
	var cur_lasers []int
	for line_num, line := range input {
		fmt.Println(line_num, len(cur_lasers))
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
				// fmt.Println(cur_lasers)
				if line[ind] == '^' {
					total_splits += 1
					if ind > 0 {
						next_lasers = append(next_lasers, ind-1)
					}
					if ind < len(line)-1 {
						next_lasers = append(next_lasers, ind+1)
					}
				} else if line[ind] == '.' {
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
