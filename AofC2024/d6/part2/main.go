package main

// Go routines, probably a bad idea, but I need to try

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	up    = iota
	right = iota
	down  = iota
	left  = iota
)

type Person struct {
	row       int
	col       int
	direction int
}

type VisitedSpot struct {
	isVisited bool
	direction int
}

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	rows := len(input)
	cols := len(input[0])

	is_visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		is_visited[i] = make([]bool, cols)
	}

	guard := Person{direction: up}

	// Find the player
	player_found := false
	for ri, row := range input {
		for ci, col := range row {
			if col == '^' {
				guard.col = ci
				guard.row = ri
				is_visited[ri][ci] = true
				input[ri] = strings.Replace(input[ri], "^", ".", 1)
				player_found = true
				break
			}
		}
		if player_found {
			break
		}
	}

	lab_layout := input

	run_count := 0
	for !guard.isEnd(rows, cols) {
		run_count++
		if guard.direction == up {
			if lab_layout[guard.row-1][guard.col] == '.' {
				guard.row--
			} else {
				guard.direction = right
			}
		} else if guard.direction == right {
			if lab_layout[guard.row][guard.col+1] == '.' {
				guard.col++
			} else {
				guard.direction = down
			}
		} else if guard.direction == down {
			if lab_layout[guard.row+1][guard.col] == '.' {
				guard.row++
			} else {
				guard.direction = left
			}
		} else if guard.direction == left {
			if lab_layout[guard.row][guard.col-1] == '.' {
				guard.col--
			} else {
				guard.direction = up
			}
		}
		is_visited[guard.row][guard.col] = true
	}

	visit_count := 0

	// Count visits
	for _, row := range is_visited {
		for _, elem := range row {
			if elem {
				visit_count++
			}
		}
	}

	fmt.Println("Rows, Cols: ", rows, cols)
	fmt.Println("Final: ", visit_count, " Run Count: ", run_count)
}

// func testWorker

func (p Person) isEnd(rows int, cols int) bool {
	return (p.col == 0 && p.direction == left) || (p.row == 0 && p.direction == up) || (p.row == rows-1 && p.direction == down) || (p.col == cols-1 && p.direction == right)
}

func getInputAsLines() ([]string, error) {
	// Read in files
	f, err := os.Open("input.txt")
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
