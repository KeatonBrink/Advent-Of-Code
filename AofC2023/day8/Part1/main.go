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

	network := make(map[string]Direction)

	for _, line := range text {
		var source, left, right string
		num_found, err := fmt.Sscanf(line, "%s = (%s, %s)", &source, &left, &right)
		if err != nil {
			panic(err)
		}
		if num_found != 3 {
			panic("Did not find three from scan")
		}

		network[source] = Direction{Left: left, Right: right}
	}

	turns_made := 1

	println("Valid attempts", return_val)
}
