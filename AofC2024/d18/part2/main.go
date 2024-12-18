package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	h            = 71
	w            = 71
	initial_bits = 1024
)

// const (
// 	h = 7
// 	w = 7
// )

const (
	r_bad_bit     = '#'
	r_empty_space = '.'
	r_person      = '@'
)

type Spot struct {
	row int
	col int
}

type Person struct {
	spot     Spot
	distance int
}

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	var final_bit string

	for bits := initial_bits; bits < len(input); bits++ {
		input_map := make([][]rune, h)
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				input_map[i] = append(input_map[i], r_empty_space)
			}
		}

		for lc, line := range input {
			if lc >= bits {
				break
			}
			var col, row int
			n, err := fmt.Sscanf(string(line), "%d,%d", &row, &col)
			if err != nil {
				panic(err)
			} else if n != 2 {
				panic("Not good")
			}
			input_map[row][col] = r_bad_bit
		}

		end := Spot{row: h - 1, col: w - 1}

		var queue []Person
		queue = append(queue, Person{})

		var is_visited []Person

		final_cost := 0

		for len(queue) > 0 {
			var cur_p Person
			cur_p, queue = queue[0], queue[1:]

			if isVisited(is_visited, cur_p) {
				continue
			}

			is_visited = append(is_visited, cur_p)

			// printGrid(input_map, cur_p)
			// fmt.Println()
			// time.Sleep(40 * time.Millisecond)

			if cur_p.spot.col == end.col && cur_p.spot.row == end.row {
				final_cost = cur_p.distance
				break
			}

			if cur_p.spot.row > 0 && input_map[cur_p.spot.row-1][cur_p.spot.col] == r_empty_space {
				new_p := Person{}
				new_p.spot = cur_p.spot
				new_p.distance = cur_p.distance + 1
				new_p.spot.row--
				// fmt.Println(new_p, cur_p)
				queue = append(queue, new_p)
			}
			if cur_p.spot.row < h-1 && input_map[cur_p.spot.row+1][cur_p.spot.col] == r_empty_space {
				new_p := Person{}
				new_p.spot = cur_p.spot
				new_p.distance = cur_p.distance + 1
				new_p.spot.row++
				// fmt.Println(new_p, cur_p)
				queue = append(queue, new_p)
			}
			if cur_p.spot.col > 0 && input_map[cur_p.spot.row][cur_p.spot.col-1] == r_empty_space {
				new_p := Person{}
				new_p.spot = cur_p.spot
				new_p.distance = cur_p.distance + 1
				new_p.spot.col--
				// fmt.Println(new_p, cur_p)
				queue = append(queue, new_p)
			}
			if cur_p.spot.col < w-1 && input_map[cur_p.spot.row][cur_p.spot.col+1] == r_empty_space {
				new_p := Person{}
				new_p.spot = cur_p.spot
				new_p.distance = cur_p.distance + 1
				new_p.spot.col++
				// fmt.Println(new_p, cur_p)
				queue = append(queue, new_p)
			}
			sort.Slice(queue[:], func(i, j int) bool {
				return queue[i].distance < queue[j].distance
			})

			// fmt.Println(queue)
		}
		if final_cost == 0 {
			final_bit = string(input[bits-1])
			fmt.Println(bits)
			printGrid(input_map, Person{})
			break
		}
	}
	fmt.Println("final_bit", final_bit)
}

func isVisited(ps []Person, p Person) bool {
	for _, cur_p := range ps {
		if cur_p.spot.row == p.spot.row && cur_p.spot.col == p.spot.col {
			return true
		}
	}
	return false
}

func printGrid(input [][]rune, p Person) {
	for ri, line := range input {
		for ci, elem := range line {
			if elem == r_bad_bit {
				fmt.Print(string(elem))
			} else if p.spot.row == ri && p.spot.col == ci {
				fmt.Print(string(r_person))
			} else {
				fmt.Print(string(r_empty_space))
			}
		}
		fmt.Println()
	}
}

func getInputAsLines() ([][]rune, error) {
	// Read in files
	// f, err := os.Open("test2.txt")
	// f, err := os.Open("test.txt")
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Create scanner
	file_scanner := bufio.NewScanner(f)
	file_scanner.Split(bufio.ScanLines)

	// Get lines
	var text [][]rune
	for file_scanner.Scan() {
		text = append(text, []rune(file_scanner.Text()))
	}
	return text, nil
}
