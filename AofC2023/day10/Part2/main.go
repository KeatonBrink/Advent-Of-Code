package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Coordinate struct {
	Row, Col, Depth int
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
		// fmt.Printf("End of for loop current cords length %d\n", len(current_coords))
	}

	count := 0

	for ind_row_in_network, row_in_network := range pipe_network {

		for ind_col_in_network, spot_in_network := range row_in_network {

			if spot_in_network == "." && next_to_loop(ind_row_in_network, ind_col_in_network, pipe_network) && !on_edge(ind_row_in_network, ind_col_in_network, pipe_network) {
				isValid := true
				temp_count := 0
				current_coords = make([]Coordinate, 0)
				current_coords = append(current_coords, Coordinate{Row: ind_row_in_network, Col: ind_col_in_network})
				var other_cord Coordinate
				var other_piece string
				for len(current_coords) > 0 {

					cur_coord := current_coords[0]
					current_coords = current_coords[1:]
					if on_edge(cur_coord.Row, cur_coord.Col, pipe_network) {
						println("On edge")
						isValid = false
					}
					if cur_coord.Row > 0 {
						other_piece = pipe_network[cur_coord.Row-1][cur_coord.Col]
						if other_piece == "." {
							other_cord = Coordinate{Row: cur_coord.Row - 1, Col: cur_coord.Col}
							current_coords = append(current_coords, other_cord)
						}
					}
					if cur_coord.Row < len(pipe_network)-1 {
						other_piece = pipe_network[cur_coord.Row+1][cur_coord.Col]
						if other_piece == "." {
							other_cord = Coordinate{Row: cur_coord.Row + 1, Col: cur_coord.Col}
							current_coords = append(current_coords, other_cord)
						}
					}
					if cur_coord.Col > 0 {
						other_piece = pipe_network[cur_coord.Row][cur_coord.Col-1]
						if other_piece == "." {
							other_cord = Coordinate{Row: cur_coord.Row, Col: cur_coord.Col - 1}
							current_coords = append(current_coords, other_cord)
						}
					}
					if cur_coord.Col < len(pipe_network[0])-1 {
						other_piece = pipe_network[cur_coord.Row][cur_coord.Col+1]
						if other_piece == "." {
							other_cord = Coordinate{Row: cur_coord.Row, Col: cur_coord.Col + 1}
							current_coords = append(current_coords, other_cord)
						}
					}
					temp_count += 1
					pipe_network[cur_coord.Row][cur_coord.Col] = "I"
				}
				if isValid {
					count += temp_count
					println(temp_count)
					PrintBoard(pipe_network)
				} else {
					// println("Bad")
					// PrintBoard(pipe_network)
				}
			}

		}
	}

	PrintBoard(pipe_network)

	fmt.Printf("Score found: %d\n", count)
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

func IsDigit(coord Coordinate, p_network [][]string) bool {
	return unicode.IsDigit(rune(p_network[coord.Row][coord.Col][len(p_network[coord.Row][coord.Col])-1]))
}

func PrintBoard(p_net [][]string) {
	for _, row := range p_net {
		for _, element := range row {
			if element != "I" && element != "." {
				element = "x"
			}
			fmt.Printf("%s ", element)
		}
		println()
	}
}

func next_to_loop(row, col int, p_network [][]string) bool {
	if row > 0 && unicode.IsDigit(rune(p_network[row-1][col][len(p_network[row-1][col])-1])) {
		return true
	}
	if col > 0 && unicode.IsDigit(rune(p_network[row][col-1][len(p_network[row][col-1])-1])) {
		return true
	}
	if row < len(p_network)-1 && unicode.IsDigit(rune(p_network[row+1][col][len(p_network[row+1][col])-1])) {
		return true
	}
	if col < len(p_network[0])-1 && unicode.IsDigit(rune(p_network[row][col+1][len(p_network[row][col+1])-1])) {
		return true
	}
	return false
}

func on_edge(row, col int, p_network [][]string) bool {
	return row == 0 || col == 0 || row == len(p_network)-1 || col == len(p_network[0])-1
}

func on_edge2(row, col int, p_network [][]string) bool {
	fmt.Printf("Row %d Col %d RMax %d CMax %d\n", row, col, len(p_network)-1, len(p_network[0])-1)
	return row == 0 || col == 0 || row == len(p_network)-1 || col == len(p_network[0])-1
}
