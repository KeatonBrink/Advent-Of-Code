package main

// Notes
// Gardens not grouped together are basically unique, so I don't care if A and A have the same name, treat them as different.
// The twist is to start storing the fence sides of each location
// Then check for sequential fence sides for each row, and for each col
// A break is a new side
// Each row is checked for the up and down direction
// Each col is checked fo the right and left direction

import (
	"bufio"
	"fmt"
	"os"
)

const (
	up    = iota
	right = iota
	down  = iota
	left  = iota
)

type Plant struct {
	perimeter [4]bool
	area      int
}

type PlantGardens [][]Plant

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)

	plant_gardens := make(PlantGardens, len(input))

	gardens_counted := make([][]bool, len(input))
	for ri, line := range input {
		plant_gardens[ri] = make([]Plant, len(line))
		gardens_counted[ri] = make([]bool, len(line))

		for ci := 0; ci < len(line); ci++ {
			plant_gardens[ri][ci] = Plant{}
		}
	}

	total_price := 0
	for ri, line := range input {
		for ci, elem := range line {
			if gardens_counted[ri][ci] {
				continue
			}
			temp_plant_gardens := plant_gardens.copyPlantGardens()
			temp_area := recurseGarden(ri, ci, elem, gardens_counted, input, temp_plant_gardens)

			temp_perim := 0

			for ri2, row := range temp_plant_gardens {
				directions := []int{up, down}
				for _, direction := range directions {
					prev_fence := -2
					for ci2, garden := range row {
						is_direction := garden.perimeter[direction]
						if is_direction {
							if prev_fence < ci2-1 {
								temp_perim++
								fmt.Println("New Side, ", ri2, ci2, directionIOTAToString(direction))
							}
							prev_fence = ci2
						}
					}
				}
			}

			for ci2 := 0; ci2 < len(temp_plant_gardens[0]); ci2++ {
				directions := []int{right, left}
				for _, direction := range directions {
					prev_fence := -2
					for ri2 := 0; ri2 < len(temp_plant_gardens); ri2++ {
						is_direction := temp_plant_gardens[ri2][ci2].perimeter[direction]
						if is_direction {
							if prev_fence < ri2-1 {
								temp_perim++
								fmt.Println("New Side, ", ri2, ci2, directionIOTAToString(direction))
							}
							prev_fence = ri2
						}
					}
				}
			}

			temp_price := temp_perim * temp_area

			total_price += temp_perim * temp_area

			fmt.Println("New Total Price, elem, garden-price, perim, area", total_price, string(elem), temp_price, temp_perim, temp_area)
		}
	}

	// total_price := 0

	// for plant_name, plant_stats := range plant_gardens {
	// 	plant_price := plant_stats.area * plant_stats.perimeter
	// 	fmt.Println("Plant, perimeter, area, cost ", plant_name, plant_stats.perimeter, plant_stats.area, plant_price)
	// 	total_price += plant_price
	// }

	fmt.Println("Final: ", total_price)
}

func directionIOTAToString(direction int) string {
	ds := ""
	switch direction {
	case up:
		ds = "up"
	case right:
		ds = "right"
	case down:
		ds = "down"
	case left:
		ds = "left"
	}
	return ds
}

func (pg PlantGardens) copyPlantGardens() PlantGardens {
	temp_plant_gardens := make(PlantGardens, len(pg))
	for ri, row := range pg {
		temp_plant_gardens[ri] = make([]Plant, len(row))
		for ci, garden := range row {
			temp_garden := Plant{}
			for direction_ind, is_direction := range garden.perimeter {
				temp_garden.perimeter[direction_ind] = is_direction
			}
			temp_garden.area = garden.area
			temp_plant_gardens[ri][ci] = temp_garden
		}
	}
	return temp_plant_gardens
}

func (p Plant) copyPlant() Plant {
	temp_plant := Plant{area: p.area}
	for i, val := range p.perimeter {
		temp_plant.perimeter[i] = val
	}
	return temp_plant
}

func recurseGarden(ri, ci int, elem rune, is_visited_gardens [][]bool, input []string, plant_gardens PlantGardens) int {
	is_visited_gardens[ri][ci] = true
	cur_area := 1
	if ri == 0 || (rune(input[ri-1][ci]) != elem) {
		plant_gardens[ri][ci].perimeter[up] = true
	} else if ri > 0 && !is_visited_gardens[ri-1][ci] {
		temp_area := recurseGarden(ri-1, ci, elem, is_visited_gardens, input, plant_gardens)
		cur_area += temp_area
	}
	if ci == len(input[ri])-1 || (rune(input[ri][ci+1]) != elem) {
		plant_gardens[ri][ci].perimeter[right] = true

	} else if ci < len(input[ri])-1 && !is_visited_gardens[ri][ci+1] {
		temp_area := recurseGarden(ri, ci+1, elem, is_visited_gardens, input, plant_gardens)
		cur_area += temp_area
	}
	if ri == len(input)-1 || (rune(input[ri+1][ci]) != elem) {
		plant_gardens[ri][ci].perimeter[down] = true
	} else if ri < len(input)-1 && !is_visited_gardens[ri+1][ci] {
		temp_area := recurseGarden(ri+1, ci, elem, is_visited_gardens, input, plant_gardens)
		cur_area += temp_area
	}
	if ci == 0 || (rune(input[ri][ci-1]) != elem) {
		plant_gardens[ri][ci].perimeter[left] = true
	} else if ci > 0 && !is_visited_gardens[ri][ci-1] {
		temp_area := recurseGarden(ri, ci-1, elem, is_visited_gardens, input, plant_gardens)
		cur_area += temp_area
	}
	return cur_area
}

func getInputAsLines() ([]string, error) {
	// Read in files
	// f, err := os.Open("test3.txt")
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
	var text []string
	for file_scanner.Scan() {
		text = append(text, file_scanner.Text())
	}
	return text, nil
}
