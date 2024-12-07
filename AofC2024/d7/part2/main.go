package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type ParsedEquation struct {
	goal   int
	inputs []int
}

const (
	plus     = iota
	multiply = iota
	concat   = iota
)

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)

	var equations []ParsedEquation

	for _, line := range input {
		no_colon_line := strings.ReplaceAll(line, ":", "")
		split_line := strings.Fields(no_colon_line)
		var slice_int_line []int
		for _, elm := range split_line {
			int_elm, err := strconv.Atoi(elm)
			if err != nil {
				fmt.Println(err)
				return
			}
			slice_int_line = append(slice_int_line, int_elm)
		}
		new_parsed_line := ParsedEquation{goal: slice_int_line[0], inputs: slice_int_line[1:]}
		equations = append(equations, new_parsed_line)
	}

	for _, equation := range equations {
		fmt.Println("Equation: ", equation)
	}

	sum := 0
	for _, equation := range equations {
		upper_limit := int(math.Pow(2, float64(len(equation.inputs)-1)))
		fmt.Println("Upper limit: ", upper_limit)
		encoded_operations := make([]int, len(equation.inputs)-1)
		for !isOverCapacity(encoded_operations) {
			// fmt.Println("Encoded: ", encoded_operations)
			if equation.isSumEqual(encoded_operations) {
				fmt.Println("Success: ", equation)
				sum += equation.goal
				break
			}
			stepEncodedOperations(encoded_operations)
		}
	}
	fmt.Println("Sum: ", sum)
}

func stepEncodedOperations(encoded_operations []int) {
	for i := 0; i < len(encoded_operations); i++ {
		encoded_operations[i]++
		if i == len(encoded_operations)-1 && encoded_operations[i] > 2 {
			break
		}
		if encoded_operations[i] <= 2 {
			break
		}
		encoded_operations[i] = 0
	}
}

func isOverCapacity(encoded_operations []int) bool {
	for _, operation := range encoded_operations {
		if operation > 2 {
			return true
		}
	}
	return false
}

func (pe ParsedEquation) isSumEqual(operations []int) bool {
	input_copy := make([]int, len(pe.inputs)-1)
	ret_val := pe.inputs[0]

	copy(input_copy, pe.inputs[1:])

	// fmt.Println(input_copy)
	// fmt.Println(pe.inputs)
	// fmt.Println(operations)

	for i := 0; i < len(input_copy); i++ {
		temp_val := input_copy[i]
		temp_operation := operations[i]
		if temp_operation == plus {
			ret_val += temp_val
		} else if temp_operation == multiply {
			ret_val *= temp_val
		} else if temp_operation == concat {
			concat, err := strconv.Atoi(fmt.Sprintf("%d%d", ret_val, temp_val))
			if err != nil {
				fmt.Println(err)
				return false
			}
			ret_val = concat
		}
	}

	return ret_val == pe.goal
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
