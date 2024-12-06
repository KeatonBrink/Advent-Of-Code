package main

// Go routines, probably a bad idea, but I need to try

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
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
	is_spot_visited bool
	direction_up    bool
	direction_right bool
	direction_down  bool
	direction_left  bool
}

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	rows := len(input)
	cols := len(input[0])

	is_visited := make([][]VisitedSpot, rows)
	for i := 0; i < rows; i++ {
		is_visited[i] = make([]VisitedSpot, cols)
	}

	guard := Person{direction: up}

	// Find the player
	player_found := false
	for ri, row := range input {
		for ci, col := range row {
			if col == '^' {
				guard.col = ci
				guard.row = ri
				is_visited[ri][ci] = VisitedSpot{is_spot_visited: true, direction_up: true}
				input[ri] = strings.Replace(input[ri], "^", ".", 1)
				player_found = true
				break
			}
		}
		if player_found {
			break
		}
	}

	is_loop_ch := make(chan bool)

	lab_layout := input

	var wg sync.WaitGroup

	loop_count := 0

	fmt.Println("Rows, Cols: ", rows, cols)
	fmt.Println("Final: ", loop_count)
}

func runSimulation(lab_layout_orig []string, guard_orig Person, is_visited_orig [][]VisitedSpot, obstruction_row, obstruction_col int, wg *sync.WaitGroup, is_loop_chan chan<- bool) {
	defer (*wg).Done()
	var lab_layout []string
	copy(lab_layout, lab_layout_orig)
	rows := len(lab_layout)
	cols := len(lab_layout[0])
	guard := guard_orig
	is_visited := make([][]VisitedSpot, rows)
	for _, row := range rows {
		row = make([]VisitedSpot, cols)
		for _, col :=
	}
	copy(is_visited, is_visited_orig)



	for !guard.isEndOfLab(rows, cols) {
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
		visited_spot := is_visited[guard.row][guard.col]
		visited_spot.is_spot_visited = true
		switch guard.direction {
		case up:
			visited_spot.direction_up = true
		case right:
			visited_spot.direction_right = true
		case down:
			visited_spot.direction_down = true
		case left:
			visited_spot.direction_left = true
		}
	}
}

// func testWorker

func (p Person) isEndOfLab(rows int, cols int) bool {
	return (p.col == 0 && p.direction == left) || (p.row == 0 && p.direction == up) || (p.row == rows-1 && p.direction == down) || (p.col == cols-1 && p.direction == right)
}

func getInputAsLines() ([]string, error) {
	// Read in files
	f, err := os.Open("test.txt")
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
