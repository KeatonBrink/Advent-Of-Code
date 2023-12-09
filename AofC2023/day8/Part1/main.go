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

	current_location := "AAA"

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

		network[source] = Direction{Left: left, Right: right}
	}

	turns_made := 0
	for directions_index := 0; current_location != "ZZZ"; directions_index++ {
		if directions_index == len(directions) {
			directions_index = 0
		}
		cur_dir := directions[directions_index]
		direction_struct := network[current_location]
		if cur_dir == 'R' {
			current_location = direction_struct.Right
		} else if cur_dir == 'L' {
			current_location = direction_struct.Left
		} else {
			panic("Invalid direction")
		}
		turns_made++
	}

	println("Turns Made", turns_made)
}
