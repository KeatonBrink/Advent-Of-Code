package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	height = 103
	width  = 101
)

// const (
// 	height       = 7
// 	width        = 11
// 	time_elapsed = 100
// )

type Robot struct {
	Px int
	Py int
	Vx int
	Vy int
}

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)

	var robots []Robot

	for _, line := range input {
		var px, py, vx, vy int
		n, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		if err != nil {
			panic(err)
		}
		if n != 4 {
			panic("Incorrect string parse")
		}
		robot := Robot{Px: px, Py: py, Vx: vx, Vy: vy}
		robots = append(robots, robot)
	}

	for time_elapsed := 0; true; time_elapsed++ {
		display_picture := generateBase()
		new_robots := copyRobots(robots)
		for _, robot := range new_robots {
			// fmt.Println(robot)
			total_dx := robot.Vx * time_elapsed
			robot.Px = (((total_dx + robot.Px) % width) + width) % width
			total_dy := robot.Vy * time_elapsed
			robot.Py = (((total_dy + robot.Py) % height) + height) % height
			display_picture[robot.Py][robot.Px] = 'X'
		}

		fmt.Println("Time Passe: ", time_elapsed)
		for _, line := range display_picture {
			fmt.Println(string(line))
		}
		if isPotentialTree(display_picture) {
			fmt.Scanln()
		}
		fmt.Println()
	}
}

func copyRobots(robots []Robot) []Robot {
	var new_robots []Robot
	for _, robot := range robots {
		new_robot := robot
		new_robots = append(new_robots, new_robot)
	}
	return new_robots
}

func isPotentialTree(image [][]rune) bool {
	max_concurrency := 0
	for _, line := range image {
		length_of_concurrency := 0
		for _, cur_rune := range line {
			if cur_rune == 'X' {
				length_of_concurrency++
				if length_of_concurrency == 10 {
					return true
				}
			} else {
				max_concurrency = max(length_of_concurrency, max_concurrency)
				length_of_concurrency = 0
			}
		}
	}
	fmt.Println(max_concurrency)
	return false
}

func generateBase() [][]rune {
	ret := make([][]rune, height)
	for ri := range height {
		ret[ri] = make([]rune, width)
		for ci := range width {
			ret[ri][ci] = ' '
		}
	}
	return ret
}

func getInputAsLines() ([]string, error) {
	// Read in files
	// f, err := os.Open("test.txt")
	// f, err := os.Open("test1.txt")
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
