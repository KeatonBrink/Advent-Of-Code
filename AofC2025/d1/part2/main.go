package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	dial := 50
	zeros := 0
	const MaxDial int = 100
	for _, value := range input {
		// fmt.Println("Value:", value, "Dial:", dial)
		// direction := value[0]
		// clicks, err := strconv.Atoi(value[1:])
		// zeros += (clicks / MaxDial)
		// clicks %= MaxDial
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// if direction == 'L' {
		// 	clicks *= -1
		// }
		// dial += clicks

		// if dial < 0 || dial >= MaxDial {
		// 	zeros += 1
		// 	if dial < 0 {
		// 		dial += MaxDial
		// 	}
		// 	dial %= MaxDial
		// }

		// fmt.Println("New Dial", dial, "Zeros:", zeros, "\n")
		fmt.Println("Value:", value, "Dial:", dial)
		direction := value[0]
		clicks, err := strconv.Atoi(value[1:])
		if err != nil {
			fmt.Println(err)
			return
		}
		for range clicks {
			if direction == 'L' {
				dial -= 1
			} else {
				dial += 1
			}
			if dial < 0 {
				dial += MaxDial
			}
			dial %= MaxDial
			if dial == 0 {
				zeros += 1
			}
		}

	}
	fmt.Println("Zeros:", zeros)
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
