package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction struct {
	Left, Right string
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

	directions := text[0]

	text = text[2:]

	var current_locations []string

	network := make(map[string]Direction)

	for _, line := range text {
		var source, left, right string
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ")", "")
		num_found, err := fmt.Sscanf(line, "%s = %s %s", &source, &left, &right)
		if err != nil {
			print(line)
			panic(err)
		}
		if num_found != 3 {
			panic("Did not find three from scan")
		}

		if source[2] == 'A' {
			current_locations = append(current_locations, source)
		}

		network[source] = Direction{Left: left, Right: right}
	}

	fmt.Printf("Total locations: %d\n", len(current_locations))

	turns_made := 0

	for directions_index := 0; !AreAllNodesAtZ(current_locations); directions_index++ {
		if directions_index == len(directions) {
			directions_index = 0
		}
		cur_dir := directions[directions_index]
		for cur_loc_ind := range current_locations {
			direction_struct := network[current_locations[cur_loc_ind]]
			if cur_dir == 'R' {
				current_locations[cur_loc_ind] = direction_struct.Right
			} else if cur_dir == 'L' {
				current_locations[cur_loc_ind] = direction_struct.Left
			} else {
				panic("Invalid direction")
			}
		}
		turns_made++
		if turns_made%100000000 == 0 {
			println(turns_made)
		}
	}

	println("Turns Made", turns_made)
}

func AreAllNodesAtZ(current_locations []string) bool {
	for _, cur_loc := range current_locations {
		if cur_loc[2] != 'Z' {
			return false
		}
	}
	return true
}
