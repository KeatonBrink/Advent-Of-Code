package main

// Go routines, probably a bad idea, but I need to try

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
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

	go_routine_count := -1

	start := time.Now()

	for ri, row := range lab_layout {
		for ci, elem := range row {
			if elem == '.' && !(ri == guard.row && ci == guard.col) {
				go_routine_count++
				go runSimulation(lab_layout, guard, is_visited, ri, ci, is_loop_ch)
			}
		}
	}

	loop_count := 0

	for go_routine_count > 0 {
		go_routine_count--
		if <-is_loop_ch {
			loop_count++
		}
	}

	elapsed := time.Since(start)

	fmt.Println("Time Passed (ms): ", elapsed.Milliseconds())

	fmt.Println("Final: ", loop_count)
}

func runSimulation(lab_layout_orig []string, guard_orig Person, is_visited_orig [][]VisitedSpot, obstruction_row, obstruction_col int, is_loop_chan chan<- bool) {
	lab_layout := make([]string, len(lab_layout_orig))
	for ri := range lab_layout_orig {
		lab_layout[ri] = lab_layout_orig[ri]
	}
	rows := len(lab_layout)
	cols := len(lab_layout[0])
	guard := guard_orig
	is_visited := make([][]VisitedSpot, rows)
	for ri, row := range is_visited {
		row = make([]VisitedSpot, cols)
		is_visited[ri] = make([]VisitedSpot, cols)
		for ci := range row {
			is_visited[ri][ci] = is_visited_orig[ri][ci]
		}
	}

	// for _, row := range is_visited {
	// 	for _, col := range row {
	// 		fmt.Printf("%v,", col)
	// 	}
	// 	fmt.Println()
	// }

	lab_layout[obstruction_row] = replaceAtIndex(lab_layout[obstruction_row], obstruction_col, '#')

	steps := 0

	for !guard.isEndOfLab(rows, cols) {
		steps++
		// fmt.Println("IN LOOP: ", obstruction_row, obstruction_col)
		// fmt.Println(lab_layout)
		// for _, row := range is_visited {
		// 	for _, col := range row {
		// 		fmt.Printf("%v,", col)
		// 	}
		// 	fmt.Println()
		// }
		// fmt.Println()
		// fmt.Println()
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
		// fmt.Println(guard)
		// fmt.Println()
		// fmt.Println()
		visited_spot := is_visited[guard.row][guard.col]
		if visited_spot.is_spot_visited == true && ((guard.direction == up && visited_spot.direction_up) || (guard.direction == right && visited_spot.direction_right) || (guard.direction == down && visited_spot.direction_down) || (guard.direction == left && visited_spot.direction_left)) {
			// fmt.Println("Is loop: ", obstruction_row, obstruction_col, steps)
			is_loop_chan <- true
			return
		}
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
		is_visited[guard.row][guard.col] = visited_spot
	}

	// fmt.Println("Is not Loop", obstruction_row, obstruction_col, steps)
	is_loop_chan <- false
	return
}

func replaceAtIndex(str string, index int, newChar rune) string {
	runes := []rune(str)
	if index >= 0 && index < len(runes) {
		runes[index] = newChar
	}
	return string(runes)
}

// func testWorker

func (p Person) isEndOfLab(rows int, cols int) bool {
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
