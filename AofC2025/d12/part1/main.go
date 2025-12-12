package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeSpace struct {
	size           int
	presents_count []int
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
	// fmt.Println(input)
	cur_present_size := 0
	is_presents := true
	var presents []int
	var tree_problems []TreeSpace
	for _, line := range input {
		if strings.Contains(line, "x") && is_presents {
			is_presents = false
		}
		if is_presents {
			if strings.Contains(line, "#") {
				cur_present_size += strings.Count(line, "#")
			} else if len(line) == 0 {
				presents = append(presents, cur_present_size)
				cur_present_size = 0
			}
		} else {
			fields := strings.Fields(line)
			size_dimensions := strings.Split(fields[0][:len(fields[0])-1], "x")
			size := 1
			for _, side := range size_dimensions {
				side_int, err := strconv.Atoi(side)
				if err != nil {
					fmt.Println(err)
					return
				}
				size *= side_int
			}

			cur_tree_space := TreeSpace{size, make([]int, 0)}

			for _, present_count_str := range fields[1:] {
				present_count, err := strconv.Atoi(present_count_str)
				if err != nil {
					fmt.Println(err)
					return
				}
				cur_tree_space.presents_count = append(cur_tree_space.presents_count, present_count)
			}

			tree_problems = append(tree_problems, cur_tree_space)
		}
	}

	total_viable := 0
	for _, tree := range tree_problems {
		present_space := 0
		for cur_ind, cur_present := range tree.presents_count {
			present_space += presents[cur_ind] * cur_present
		}

		if present_space <= tree.size {
			total_viable += 1
		}
	}

	fmt.Println(total_viable)
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
