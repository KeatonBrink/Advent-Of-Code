package main

import (
	"bufio"
	"math"
	"os"
)

type Coordinate struct {
	Row, Col int
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

	// Grab galaxies and create new
	var galaxies []Coordinate
	var galaxy_image [][]byte
	for row, line := range text {
		var line_bytes []byte
		for col, character := range line {
			line_bytes = append(line_bytes, byte(character))
			if character == '#' {
				galaxies = append(galaxies, Coordinate{Row: row, Col: col})
			}
		}
		galaxy_image = append(galaxy_image, line_bytes)
	}

	// Find any row that has to be repeated
	double_row := make(map[int]bool)
	for i_row, row := range galaxy_image {
		found_galaxy := false
		for _, char := range row {
			if char == '#' {
				found_galaxy = true
				break
			}
		}
		double_row[i_row] = !found_galaxy
	}

	// Find any column that has to be repeated
	double_col := make(map[int]bool)
	for i_col := 0; i_col < len(galaxy_image[0]); i_col++ {
		found_galaxy := false
		for i_row := 0; i_row < len(galaxy_image); i_row++ {
			if galaxy_image[i_row][i_col] == '#' {
				found_galaxy = true
				break
			}
		}
		double_col[i_col] = !found_galaxy
	}

	total_distance := 0

	// Start Calculating Distances
	for i_start_galaxy, start_galaxy := range galaxies {
		for i_other_galaxy := i_start_galaxy + 1; i_other_galaxy < len(galaxies); i_other_galaxy++ {
			other_galaxy := galaxies[i_other_galaxy]
			total_distance += int(math.Abs(float64(start_galaxy.Col)-float64(other_galaxy.Col))) + int(math.Abs(float64(start_galaxy.Row)-float64(other_galaxy.Row))) + FindDoubleDistances(start_galaxy, other_galaxy, double_row, double_col)

		}
	}

	println(total_distance)
}

func FindDoubleDistances(start_galaxy, other_galaxy Coordinate, double_row, double_col map[int]bool) int {
	ret_doubles := 0
	var cur_i int
	var end_i int
	if start_galaxy.Row < other_galaxy.Row {
		cur_i = start_galaxy.Row + 1
		end_i = other_galaxy.Row - 1
	} else {
		cur_i = other_galaxy.Row + 1
		end_i = start_galaxy.Row - 1
	}

	for ; cur_i <= end_i; cur_i++ {
		if double_row[cur_i] {
			ret_doubles += 1
			// ret_doubles += 100 - 1
		}
	}

	if start_galaxy.Col < other_galaxy.Col {
		cur_i = start_galaxy.Col + 1
		end_i = other_galaxy.Col - 1
	} else {
		cur_i = other_galaxy.Col + 1
		end_i = start_galaxy.Col - 1
	}

	for ; cur_i <= end_i; cur_i++ {
		if double_col[cur_i] {
			ret_doubles += 1
			// ret_doubles += 100 - 1
		}
	}

	return ret_doubles
}
