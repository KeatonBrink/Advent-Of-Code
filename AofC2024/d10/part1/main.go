package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

const (
	up    = iota
	right = iota
	down  = iota
	left  = iota
)

type TrailPath struct {
	head_x int
	head_y int
	peak_x int
	peak_y int
}

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)

	// This will be left unmodified
	topo_map := make([][]int, len(input))

	// Copy over ints to map
	for _, input_line := range input {
		topo_map_line := make([]int, len(input_line))
		for _, elem := range input_line {
			topo_map_line = append(topo_map_line, int(elem)-'0')
		}
	}

}

func moveOnMap(pos_x, pos_y int, topo_map [][]int, ch chan<- TrailPath, wg *sync.WaitGroup) {
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
