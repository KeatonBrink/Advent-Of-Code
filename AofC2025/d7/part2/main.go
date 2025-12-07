package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Laser struct {
	Location int
	Total    int
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
	// Location of active lasers and the total for the given spot
	var cur_lasers []Laser
	for _, line := range input {
		var next_lasers []Laser
		if len(cur_lasers) == 0 {
			for ind, char := range line {
				if char == 'S' {
					next_lasers = []Laser{{ind, 1}}
					break
				}
			}
		} else {
			if !strings.ContainsRune(line, '^') {
				continue
			}
			// printLaserTotal(cur_lasers)
			for _, laser := range cur_lasers {
				if line[laser.Location] == '^' {
					if laser.Location > 0 {
						if !isLaserInSlice(Laser{laser.Location - 1, laser.Total}, next_lasers) {
							next_lasers = append(next_lasers, Laser{laser.Location - 1, laser.Total})
						} else {
							next_lasers = incrementTotalInSlice(Laser{laser.Location - 1, laser.Total}, next_lasers)
						}
					}
					if laser.Location < len(line)-1 {
						if !isLaserInSlice(Laser{laser.Location + 1, laser.Total}, next_lasers) {
							next_lasers = append(next_lasers, Laser{laser.Location + 1, laser.Total})
						} else {
							next_lasers = incrementTotalInSlice(Laser{laser.Location + 1, laser.Total}, next_lasers)
						}
					}
				} else if line[laser.Location] == '.' {
					if !isLaserInSlice(Laser{laser.Location, laser.Total}, next_lasers) {
						next_lasers = append(next_lasers, Laser{laser.Location, laser.Total})
					} else {
						next_lasers = incrementTotalInSlice(Laser{laser.Location, laser.Total}, next_lasers)
					}
				}
			}
		}
		cur_lasers = next_lasers
	}
	printLaserTotal(cur_lasers)
}

func printLaserTotal(cur_lasers []Laser) {
	temp_total := 0
	for _, cur_laser := range cur_lasers {
		temp_total += cur_laser.Total
	}
	fmt.Println(temp_total)
}

func incrementTotalInSlice(target_laser Laser, cur_slice []Laser) []Laser {
	for ind, cur_laser := range cur_slice {
		if cur_laser.Location == target_laser.Location {
			cur_slice[ind].Total += target_laser.Total
			break
		}
	}
	return cur_slice
}

func isLaserInSlice(target_laser Laser, cur_slice []Laser) bool {
	for _, cur_laser := range cur_slice {
		if cur_laser.Location == target_laser.Location {
			return true
		}
	}
	return false
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
