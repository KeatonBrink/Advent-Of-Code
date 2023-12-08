package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction struct {
	Left, Right string
}

func main() {
	// input_file_name := "input.txt"
	input_file_name := "test_input.txt"

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

	var current_location string = "AAA"

	network := make(map[string]Direction)

	for _, line := range text {
		var source, left, right string
		num_found, err := fmt.Sscanf(line, "%s = (%s, %s)\n", &source, &left, &right)
		if err != nil {
			print(line)
			panic(err)
		}
		if num_found != 3 {
			panic("Did not find three from scan")
		}

		// if i == 0 {
		// 	current_location = source
		// }
		network[source] = Direction{Left: left, Right: right}
	}

	turns_made := 0

	for directions_index := 0; current_location != "ZZZ"; directions_index++ {
		if directions_index == len(directions) {
			directions_index = 0
		}
		direction_struct := network[current_location]
		if directions[directions_index] == 'R' {
			current_location = direction_struct.Right
		} else if directions[directions_index] == 'L' {
			current_location = direction_struct.Left
		} else {
			panic("Invalid direction")
		}
		turns_made++
	}

	println("Valid attempts", turns_made)
}
