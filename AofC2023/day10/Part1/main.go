package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Coordinate struct {
	Row, Col, Depth int
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

	var current_coords []Coordinate
	var pipe_network [][]string
	for row, line := range text {
		var line_string []string
		for col, character := range line {
			if character == '7' {
				character = '?'
			}
			line_string = append(line_string, string(character))
			if character == 'S' {
				current_coords = append(current_coords, Coordinate{Row: row, Col: col, Depth: 0})
			}
		}
		pipe_network = append(pipe_network, line_string)
	}

	depth := 0

	for len(current_coords) > 0 {
		// println("Start of loop")
		// PrintCoords(current_coords)
		var cur_coord Coordinate
		cur_coord, current_coords = current_coords[0], current_coords[1:]
		cur_piece := pipe_network[cur_coord.Row][cur_coord.Col]
		// fmt.Printf("Current Coord %d %d %s \n", cur_coord.Row, cur_coord.Col, cur_piece)
		var other_piece string
		if cur_coord.Row > 0 {
			other_piece = pipe_network[cur_coord.Row-1][cur_coord.Col]
			if HasUp(cur_piece) && HasDown(other_piece) {
				// println("Appending")
				current_coords = append(current_coords, Coordinate{Row: cur_coord.Row - 1, Col: cur_coord.Col, Depth: cur_coord.Depth + 1})
				// PrintCoords(current_coords)
			}
		}
		if cur_coord.Row < len(pipe_network)-1 {
			other_piece = pipe_network[cur_coord.Row+1][cur_coord.Col]
			if HasDown(cur_piece) && HasUp(other_piece) {
				// println("Appending")
				current_coords = append(current_coords, Coordinate{Row: cur_coord.Row + 1, Col: cur_coord.Col, Depth: cur_coord.Depth + 1})
				// PrintCoords(current_coords)
			}
		}
		if cur_coord.Col > 0 {
			other_piece = pipe_network[cur_coord.Row][cur_coord.Col-1]
			if HasLeft(cur_piece) && HasRight(other_piece) {
				// println("Appending")
				current_coords = append(current_coords, Coordinate{Row: cur_coord.Row, Col: cur_coord.Col - 1, Depth: cur_coord.Depth + 1})
				// PrintCoords(current_coords)
			}
		}
		if cur_coord.Col < len(pipe_network[0])-1 {
			other_piece = pipe_network[cur_coord.Row][cur_coord.Col+1]
			if HasRight(cur_piece) && HasLeft(other_piece) {
				// println("Appending")
				current_coords = append(current_coords, Coordinate{Row: cur_coord.Row, Col: cur_coord.Col + 1, Depth: cur_coord.Depth + 1})
				// PrintCoords(current_coords)
			}
		}
		pipe_network[cur_coord.Row][cur_coord.Col] = strconv.Itoa(cur_coord.Depth)
		depth = cur_coord.Depth
		// fmt.Printf("End of for loop current cords length %d\n", len(current_coords))
	}

	fmt.Printf("Score found: %d\n", depth)
}

func HasUp(piece string) bool {
	return piece == "S" || piece == "|" || piece == "L" || piece == "J"
}

func HasDown(piece string) bool {
	return piece == "S" || piece == "|" || piece == "?" || piece == "F"
}

func HasLeft(piece string) bool {
	return piece == "S" || piece == "-" || piece == "?" || piece == "J"
}

func HasRight(piece string) bool {
	return piece == "S" || piece == "-" || piece == "L" || piece == "F"
}

func PrintCoords(coords []Coordinate) {
	println("Starting printCoords")
	for _, c := range coords {
		fmt.Printf("Row %d Col %d Depth %d\n", c.Row, c.Col, c.Depth)
	}
	println("Ending printCoords")
}
