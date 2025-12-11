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

type Memo struct {
	reached map[string]int
}

func (m Memo) reset() {
	for key, _ := range m.reached {
		m.reached[key] = -1
	}
	m.reached["out"] = -1
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
	total_paths := 1

	connection_dict := make(map[string][]string)
	// targets := []string{"fft", "dac", "out"}

	memo := Memo{make(map[string]int)}

	for _, line := range input {
		line_fields := strings.Fields(line)
		home := line_fields[0]
		home = home[:len(home)-1]
		outgoing := line_fields[1:]
		connection_dict[home] = outgoing
		memo.reached[home] = -1
	}

	if false {

		// problematic_racks := []string{"fft", "dac"}
		target_destination := [][]string{{"svr", "fft"}, {"fft", "dac"}, {"dac", "out"}}

		for _, td := range target_destination {
			fmt.Println(td)
			cur_paths := 0
			var paths []Path
			paths = append(paths, Path{[]string{td[0]}})
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
					} else if potential_next_path_add == td[1] {
						cur_paths++
						continue
					}
					new_path := make([]string, len(cur_path)+1)
					copy(new_path, cur_path)
					new_path[len(new_path)-1] = potential_next_path_add
					paths = append(paths, Path{new_path})
				}
			}
			total_paths *= cur_paths
			fmt.Println(total_paths)
		}
	} else {
		target_destination := [][]string{{"svr", "fft"}, {"fft", "dac"}, {"dac", "out"}}

		for _, td := range target_destination {
			fmt.Println(td)
			memo.reset()
			memo = descend(memo, td[0], td[1], connection_dict)
			total_paths *= memo.reached[td[0]]
			fmt.Println(total_paths)
		}

	}

	fmt.Println(total_paths)
}

func descend(memo Memo, cur_loc, target string, connection_dict map[string][]string) Memo {
	// fmt.Println(cur_loc, memo)
	if cur_loc == target {
		memo.reached[cur_loc] = 1
		return memo
	}
	total := 0
	for _, connection := range connection_dict[cur_loc] {
		if memo.reached[connection] < 0 {
			memo = descend(memo, connection, target, connection_dict)
		}
		total += memo.reached[connection]
	}
	memo.reached[cur_loc] = total
	return memo
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
