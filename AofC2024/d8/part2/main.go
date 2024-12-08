package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Thought process, one run to find all the possible towers
// Then run through each spot again to find all the distances for each type of node, and search for a double
// Not sure what the double looks like at this point, but probably just check if the double exists before adding

// Note: map variables pertain to the data structure maps

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)

	city := input

	blank_ant_distances := make(map[rune][]float64, 0)

	for _, line := range city {
		for _, ant := range line {
			blank_ant_distances[ant] = make([]float64, 0)
		}
	}

	spots_found := 0

	// Lots of logic that should probably be broken down
	// Cycle through every spot to look for distances
	for ri, row := range city {
		for ci := range row {
			// for key, val := range blank_ant_distances {
			// 	fmt.Println(key, val)
			// }
			// fmt.Println()
			cur_ant_distance_map := copyAntDistanceMap(blank_ant_distances)
			cur_ant_slope_map := copyAntDistanceMap(blank_ant_distances)
			spot_found := false
			// Compare current spot against all others
			for cur_ri, cur_row := range city {
				for cur_ci, cur_spot := range cur_row {
					if cur_spot == '.' || (ri == cur_ri || ci == cur_ci) {
						// Not antannae
						continue
					}
					// Unnecessary, but feels clearier
					cur_ant := cur_spot
					dist := getDistanceBetweenPoints(ri, ci, cur_ri, cur_ci)
					slope := getSlopeBetweenPoints(ri, ci, cur_ri, cur_ci)
					prev_ant_distances := cur_ant_distance_map[cur_ant]
					prev_ant_slopes := cur_ant_slope_map[cur_ant]
					for padi, prev_ant_dist := range prev_ant_distances {
						if (dist*2 == prev_ant_dist || dist/2 == prev_ant_dist) && slope == prev_ant_slopes[padi] {
							fmt.Println("ri1, ci1, ri2, ci2", ri, ci, cur_ri, cur_ci)
							fmt.Println("Dist1, Dist2", dist, prev_ant_dist)
							fmt.Println()
							spots_found++
							spot_found = true
							break
						}
					}
					if spot_found {
						break
					} else {
						prev_ant_distances = append(prev_ant_distances, dist)
						cur_ant_distance_map[cur_ant] = prev_ant_distances
						prev_ant_slopes = append(prev_ant_slopes, slope)
						cur_ant_slope_map[cur_ant] = prev_ant_slopes
					}
				}
				if spot_found {
					break
				}
			}
			// for key, val := range cur_ant_distance_map {
			// 	fmt.Println(key, val)
			// }
			// fmt.Println()
		}
	}

	fmt.Println("Final: ", spots_found)
}

func getSlopeBetweenPoints(ri1, ci1, ri2, ci2 int) float64 {
	return float64(ci2-ci1) / float64(ri2-ri1)
}

func getDistanceBetweenPoints(ri1, ci1, ri2, ci2 int) float64 {
	dr := math.Pow(float64(ri1-ri2), 2)
	dc := math.Pow(float64(ci1-ci2), 2)
	return math.Sqrt(dr + dc)
}

func copyAntDistanceMap(og_map map[rune][]float64) map[rune][]float64 {
	ret_map := make(map[rune][]float64, 0)
	for key, value := range og_map {
		ret_map[key] = make([]float64, len(value))
		copy(ret_map[key], value)
	}
	return ret_map
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
