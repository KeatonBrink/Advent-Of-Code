package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Path struct {
	path []string
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
	connection_dict := make(map[string][]string)

	for _, line := range input {
		line_fields := strings.Fields(line)
		home := line_fields[0]
		home = home[:len(home)-1]
		outgoing := line_fields[1:]
		connection_dict[home] = outgoing
	}

	total_paths := 0
	var paths []Path
	paths = append(paths, Path{[]string{"you"}})
	for len(paths) > 0 {
		var popped_path Path
		popped_path, paths = paths[0], paths[1:]
		cur_path := popped_path.path
		cur_loc := cur_path[len(cur_path)-1]
		potential_next_paths := connection_dict[cur_loc]
		for _, potential_next_path_add := range potential_next_paths {
			// Avoid circles
			if slices.Contains(cur_path, potential_next_path_add) {
				continue
			} else if potential_next_path_add == "out" {
				total_paths++
				continue
			}
			new_path := make([]string, len(cur_path)+1)
			copy(new_path, cur_path)
			new_path[len(new_path)-1] = potential_next_path_add
			paths = append(paths, Path{new_path})
		}
	}

	fmt.Println(total_paths)
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
