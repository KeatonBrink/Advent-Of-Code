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
	for ri, input_line := range input {
		topo_map[ri] = make([]int, len(input_line))
		for ci, elem := range input_line {
			temp_elem := int(elem) - '0'
			if temp_elem > 9 || temp_elem < 0 {
				temp_elem = -1
			}
			topo_map[ri][ci] = temp_elem
		}
	}

	fmt.Println(topo_map)

	var wg sync.WaitGroup
	success_chan := make(chan TrailPath)

	for ri, row := range topo_map {
		for ci, elem := range row {
			if elem == 0 {
				wg.Add(1)
				go moveOnMap(ri, ci, ri, ci, topo_map, success_chan, &wg)
			}
		}
	}

	go func() {
		wg.Wait()
		close(success_chan)
	}()

	success_count := 0

	for range success_chan {
		success_count++
	}

	fmt.Println("Final: ", success_count)
}

func moveOnMap(pos_x, pos_y, start_x, start_y int, topo_map [][]int, ch chan<- TrailPath, wg *sync.WaitGroup) {
	defer wg.Done()
	cur_elem := topo_map[pos_x][pos_y]
	if cur_elem == 9 {
		ch <- TrailPath{head_x: start_x, head_y: start_y, peak_x: pos_x, peak_y: pos_y}
		return
	}
	next_elem := cur_elem + 1
	if pos_x > 0 && topo_map[pos_x-1][pos_y] == next_elem {
		wg.Add(1)
		go moveOnMap(pos_x-1, pos_y, start_x, start_y, topo_map, ch, wg)
	}
	if pos_y < len(topo_map[0])-1 && topo_map[pos_x][pos_y+1] == next_elem {
		wg.Add(1)
		go moveOnMap(pos_x, pos_y+1, start_x, start_y, topo_map, ch, wg)
	}
	if pos_x < len(topo_map)-1 && topo_map[pos_x+1][pos_y] == next_elem {
		wg.Add(1)
		go moveOnMap(pos_x+1, pos_y, start_x, start_y, topo_map, ch, wg)
	}
	if pos_y > 0 && topo_map[pos_x][pos_y-1] == next_elem {
		wg.Add(1)
		go moveOnMap(pos_x, pos_y-1, start_x, start_y, topo_map, ch, wg)
	}
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
