package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
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
	// Find the max row and col
	max_row := -1
	max_col := -1
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
		if ind2 > max_row {
			max_row = ind2
		}
		if ind1 > max_col {
			max_col = ind1
		}
		cur_tile_corner := TileCorner{ind2, ind1}
		tile_corners = append(tile_corners, cur_tile_corner)

	}
	// Generating the grid
	var grid [][]rune
	for row := 0; row <= max_row+1; row++ {
		var temp_line []rune
		for col := 0; col <= max_col+1; col++ {
			// fmt.Println(row, col)
			temp_line = append(temp_line, '.')
		}
		grid = append(grid, temp_line)
	}
	for _, tc := range tile_corners {
		grid[tc.row][tc.col] = '#'
	}
	fmt.Println("Finished Generating Grid Corners")

	// Working on green capet outline
	for _, tc1 := range tile_corners {
		// Look at two orthogonal directions
		close_north := TileCorner{-1, -1}
		close_west := TileCorner{-1, -1}
		for _, tc2 := range tile_corners {
			if (tc1.col == tc2.col && tc1.row == tc2.row) || (tc1.col != tc2.col && tc1.row != tc2.row) {
				continue
			}
			// Closest North Calc
			if tc1.col == tc2.col && tc2.row < tc1.row && tc2.row > close_north.row {
				close_north = tc2
			}
			// Closest West Calc
			if tc1.row == tc2.row && tc2.col < tc1.col && tc2.col > close_west.col {
				close_west = tc2
			}
		}
		if close_north.row > -1 {
			for ind_row := close_north.row + 1; ind_row < tc1.row; ind_row++ {
				grid[ind_row][tc1.col] = 'X'
			}
		}
		if close_west.col > -1 {
			for ind_col := close_west.col + 1; ind_col < tc1.col; ind_col++ {
				grid[tc1.row][ind_col] = 'X'
			}
		}
	}
	fmt.Println("Finished Generating Grid Outline")

	// Ray Casting
	for row := 0; row < len(grid); row++ {
		is_casting := false
		is_edge := false
		prior_edge_direction := 0
		for col := 0; col < len(grid[0]); col++ {
			elem := grid[row][col]
			if elem == '#' {

				if !is_edge {
					if grid[row+1][col] != '.' {
						prior_edge_direction = 1
					} else {
						prior_edge_direction = -1
					}
				} else {
					new_edge_direction := 0

					if grid[row+1][col] != '.' {
						new_edge_direction = 1
					} else {
						new_edge_direction = -1
					}
					if new_edge_direction != prior_edge_direction {
						is_casting = !is_casting
					}
				}
				is_edge = !is_edge
			} else if elem == 'X' && !is_edge {
				is_casting = !is_casting
			} else if elem == '.' && is_casting {
				grid[row][col] = 'X'
			}
		}

	}
	fmt.Println("Finished Generating Grid")

	if file_name == "test.txt" {
		for _, line := range grid {
			fmt.Println(string(line))
		}
	}

	// fmt.Println(grid)
	largest_area := 0
	start_time := time.Now()
	for tile_corner1_ind := 0; tile_corner1_ind < len(tile_corners)-1; tile_corner1_ind++ {
		if tile_corner1_ind%2 == 0 {
			fmt.Println(tile_corner1_ind, time.Since(start_time))
		}
		for tile_corner2_ind := tile_corner1_ind + 1; tile_corner2_ind < len(tile_corners); tile_corner2_ind++ {
			tc1 := tile_corners[tile_corner1_ind]
			tc2 := tile_corners[tile_corner2_ind]
			cur_area := tc1.GetArea(tc2)
			if cur_area <= largest_area {
				continue
			}
			is_valid := true
			col1 := min(tc1.col, tc2.col)
			col2 := max(tc1.col, tc2.col)
			for temp_row := min(tc1.row, tc2.row); temp_row <= max(tc1.row, tc2.row) && is_valid; temp_row++ {
				if grid[temp_row][col1] == '.' || grid[temp_row][col2] == '.' {
					is_valid = false
				}
			}
			row1 := min(tc1.row, tc2.row)
			row2 := max(tc1.row, tc2.row)
			for temp_col := min(tc1.col, tc2.col); temp_col <= max(tc1.col, tc2.col) && is_valid; temp_col++ {
				if grid[row1][temp_col] == '.' || grid[row2][temp_col] == '.' {
					is_valid = false
				}
			}
			if is_valid {
				fmt.Println("Largest Area so far", largest_area)
				largest_area = cur_area
			}
		}
	}
	// for tile_corner1_ind := 0; tile_corner1_ind < len(tile_corners)-1; tile_corner1_ind++ {
	// 	for tile_corner2_ind := tile_corner1_ind + 1; tile_corner2_ind < len(tile_corners); tile_corner2_ind++ {
	// 		tc1 := tile_corners[tile_corner1_ind]
	// 		tc2 := tile_corners[tile_corner2_ind]
	// 		found_other_red_tile := false
	// 		for temp_row := min(tc1.row, tc2.row) + 1; temp_row < max(tc1.row, tc2.row) && !found_other_red_tile; temp_row++ {
	// 			for temp_col := min(tc1.col, tc2.col) + 1; temp_col < max(tc1.col, tc2.col) && !found_other_red_tile; temp_col++ {
	// 				for _, temp_tc := range tile_corners {
	// 					if temp_tc.row == temp_row && temp_tc.col == temp_col {
	// 						found_other_red_tile = true
	// 						break
	// 					}
	// 				}
	// 			}
	// 		}
	// 		// Use eventually
	// 		if !found_other_red_tile {

	// 			cur_area := tc1.GetArea(tc2)
	// 			fmt.Println("No other tile found:", tc1.row, tc1.col, tc2.row, tc2.col, "Cur Area", cur_area)
	// 			if cur_area > largest_area {
	// 				largest_area = cur_area
	// 			}
	// 		}
	// 	}
	// }
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
