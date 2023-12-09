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

	var turns_made_by_location []int = make([]int, len(current_locations))

	for i, cur_loc := range current_locations {
		turns_made := 0
		for directions_index := 0; cur_loc[2] != 'Z'; directions_index++ {
			if directions_index == len(directions) {
				directions_index = 0
			}
			cur_dir := directions[directions_index]
			direction_struct := network[cur_loc]
			if cur_dir == 'R' {
				cur_loc = direction_struct.Right
			} else if cur_dir == 'L' {
				cur_loc = direction_struct.Left
			} else {
				panic("Invalid direction")
			}
			turns_made++
		}
		turns_made_by_location[i] = turns_made
	}

	println("Turns Made", LCM(turns_made_by_location...))
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for a%b != 0 {
		t := b
		b = a % b
		a = t
	}
	return b
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {
	lcm := 1
	for i := 0; i < len(integers); i++ {
		lcm = (lcm * integers[i]) / GCD(lcm, integers[i])
	}
	return lcm
}
