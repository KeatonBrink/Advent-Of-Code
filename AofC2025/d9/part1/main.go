package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type TileCorner struct {
	row int
	col int
}

func (tc1 TileCorner) GetArea(tc2 TileCorner) int {
	return (int(math.Abs((float64(tc1.row - tc2.row)))) + 1) * (int(math.Abs(float64(tc1.col-tc2.col))) + 1)
}

func main() {
	file_name := "input.txt"
	args := os.Args[1:]
	if len(args) >= 1 {
		file_name = args[0]
	}
	input, err := getInputAsLines(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	var tile_corners []TileCorner
	for _, line := range input {
		parts := strings.Split(line, ",")
		ind1, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		ind2, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		cur_tile_corner := TileCorner{ind1, ind2}
		tile_corners = append(tile_corners, cur_tile_corner)

	}
	largest_area := 0
	for tile_corner1_ind := 0; tile_corner1_ind < len(tile_corners)-1; tile_corner1_ind++ {
		for tile_corner2_ind := tile_corner1_ind + 1; tile_corner2_ind < len(tile_corners); tile_corner2_ind++ {
			cur_area := tile_corners[tile_corner1_ind].GetArea(tile_corners[tile_corner2_ind])
			if cur_area > largest_area {
				largest_area = cur_area
			}
		}
	}
	fmt.Println(largest_area)
}

func getInputAsLines(file_name string) ([]string, error) {
	// Read in files
	f, err := os.Open(file_name)
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
