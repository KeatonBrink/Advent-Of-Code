package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Up = iota
	Down
	Left
	Right
)

type EndOfLight struct {
	Row, Col, Direction int
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

	mirror_layout := text

	var current_lights []EndOfLight

	var visited_positions := 

	start_pos := mirror_layout[0][0]

	start_EOL := EndOfLight{Row: 0, Col: 0, Direction: Right}

	if start_pos == '.' || start_pos == '-' {
		current_lights = append(current_lights, start_EOL)
	} else {
		start_EOL.Direction = Down
		current_lights = append(current_lights, start_EOL)
	}

	ret_val := 0

	fmt.Printf("Final lens hash: %d\n", ret_val)
}

struct EOLToString(eol EndOfLight) string {
	return string(eol.Row) + "," string(eol.Col) + "," + string(Direction)
}