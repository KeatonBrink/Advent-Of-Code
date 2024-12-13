package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ClawMachine struct {
	Ax int
	Ay int
	Bx int
	By int
	Px int
	Py int
}

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)
	var claw_machines []ClawMachine

	cur_claw_machine := ClawMachine{}
	for _, line := range input {
		if line == "" {
			continue
		}
		var temp1, temp2 int
		if line[:6] == "Button" {
			var input_name string
			temp_line := strings.ReplaceAll(line, ":", "")
			temp_line = strings.ReplaceAll(temp_line, "+", " ")
			temp_line = strings.ReplaceAll(temp_line, ",", "")
			n, err := fmt.Sscanf(temp_line, "Button %s X %d Y %d", &input_name, &temp1, &temp2)
			if err != nil {
				fmt.Println(n, input_name)
				panic(err)
			}
			if n == 3 {
				if input_name == "A" {
					cur_claw_machine.Ax = temp1
					cur_claw_machine.Ay = temp2
				} else if input_name == "B" {
					cur_claw_machine.Bx = temp1
					cur_claw_machine.By = temp2
				} else {
					panic("Unknown input type")
				}
				continue
			}
		}
		temp_line := strings.ReplaceAll(line, "=", " ")
		temp_line = strings.ReplaceAll(temp_line, ",", "")
		n, err := fmt.Sscanf(temp_line, "Prize: X %d Y %d", &temp1, &temp2)
		if err != nil {
			panic(err)
		}
		if n == 2 {
			cur_claw_machine.Px = temp1
			cur_claw_machine.Py = temp2
			claw_machines = append(claw_machines, cur_claw_machine)
			cur_claw_machine = ClawMachine{}
		}
	}
	for _, claw_machine := range claw_machines {
		fmt.Println(claw_machine)
	}
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
