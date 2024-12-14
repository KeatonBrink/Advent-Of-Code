package main

import (
	"bufio"
	"fmt"
	"os"
)

// const (
// 	height       = 103
// 	width        = 101
// 	time_elapsed = 100
// )

const (
	height       = 7
	width        = 11
	time_elapsed = 100
)

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

	var quadrants [4]int

	for _, robot := range robots {
		fmt.Println(robot)
		total_dx := robot.Vx * time_elapsed
		robot.Px = (total_dx + robot.Px) % width
		total_dy := robot.Vy * time_elapsed
		robot.Py = (total_dy + robot.Py) % height
		if robot.Px < width/2 {
			if robot.Py < height/2 {
				fmt.Println("q1")
				quadrants[0]++
			} else if robot.Py > height/2 {
				fmt.Println("q2")
				quadrants[1]++
			}
		} else if robot.Px > width/2 {
			if robot.Py < height/2 {
				fmt.Println("q3")
				quadrants[2]++
			} else if robot.Py > height/2 {
				fmt.Println("q4")
				quadrants[3]++
			}
		}
		fmt.Println(robot)
		fmt.Println()
	}

	fmt.Println(quadrants)
	fmt.Println("Final: ", quadrants[0]*quadrants[1]*quadrants[2]*quadrants[3])
}

func getInputAsLines() ([]string, error) {
	// Read in files
	f, err := os.Open("test.txt")
	// f, err := os.Open("input.txt")
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
