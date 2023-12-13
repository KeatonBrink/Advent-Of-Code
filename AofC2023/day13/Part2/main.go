package main

import (
	"bufio"
	"fmt"
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

	var reflection_problems [][]string

	var cur_reflection_problem []string

	for _, line := range text {
		if line != "" {
			cur_reflection_problem = append(cur_reflection_problem, line)
		} else {
			reflection_problems = append(reflection_problems, cur_reflection_problem)
			cur_reflection_problem = make([]string, 0)
		}
	}
	reflection_problems = append(reflection_problems, cur_reflection_problem)

	count := 0

	for _, cur_reflection_problem = range reflection_problems {
		// PrintReflectionProblem(cur_reflection_problem)
		count += FindColumnReflection(cur_reflection_problem)
		count += FindRowReflection(cur_reflection_problem)
		// println(count)
	}

	fmt.Printf("Final reflection count: %d\n", count)
}

func FindColumnReflection(cur_reflection_problem []string) int {
	count := 0
	for start_right := 1; start_right < len(cur_reflection_problem[0]); start_right++ {
		difference_count := 0
		for _, cur_row := range cur_reflection_problem {
			left := reverse(cur_row[:start_right])
			right := cur_row[start_right:]
			for comparator_ind := 0; comparator_ind < len(left) && comparator_ind < len(right); comparator_ind++ {
				if left[comparator_ind] != right[comparator_ind] {
					difference_count++
				}
			}
			if difference_count > 1 {
				break
			}
		}
		if difference_count == 1 {
			count += start_right
		}
	}
	return count
}

func FindRowReflection(cur_reflection_problem []string) int {
	count := 0
	for start_down := 1; start_down < len(cur_reflection_problem); start_down++ {
		difference_count := 0
		for cur_col := 0; cur_col < len(cur_reflection_problem[0]); cur_col++ {
			//Convert the column to a string for copying from above
			column_string := ""
			for i := 0; i < len(cur_reflection_problem); i++ {
				column_string += string(cur_reflection_problem[i][cur_col])
			}
			left := reverse(column_string[:start_down])
			right := column_string[start_down:]
			for comparator_ind := 0; comparator_ind < len(left) && comparator_ind < len(right); comparator_ind++ {
				if left[comparator_ind] != right[comparator_ind] {
					difference_count++
				}
			}
			if difference_count > 1 {
				break
			}
		}
		if difference_count == 1 {
			// fmt.Printf("Valid column %d\n", start_down)
			count += start_down * 100
		}
	}
	return count
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func PrintReflectionProblem(p []string) {
	for _, str := range p {
		println(str)
	}
}
