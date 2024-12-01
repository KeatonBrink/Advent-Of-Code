package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	// Read in files
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create scanner
	file_scanner := bufio.NewScanner(f)
	file_scanner.Split(bufio.ScanLines)

	// Get lines
	var text []string
	for file_scanner.Scan() {
		text = append(text, file_scanner.Text())
	}

	var listA []int
	var listB []int
	for i := 0; i < len(text); i++ {
		var ElemA, ElemB int
		_, err := fmt.Sscanf(text[i], "%d   %d", &ElemA, &ElemB)
		if err != nil {
			fmt.Println(err)
			return
		}
		listA = append(listA, ElemA)
		listB = append(listB, ElemB)
	}

	// Sort
	sort.Ints(listA)
	sort.Ints(listB)

	dist := 0

	for i := 0; i < len(listA) && i < len(listB); i++ {
		ElemA := listA[i]
		ElemB := listB[i]
		iterDist := int(math.Max(float64(ElemA), float64(ElemB)) - math.Min(float64(ElemA), float64(ElemB)))
		dist += iterDist
	}

	fmt.Println("Final Distance ", dist)
}
