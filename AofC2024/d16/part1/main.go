package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	up    = iota
	right = iota
	down  = iota
	left  = iota
)

const (
	r_wall   = '#'
	r_space  = '.'
	r_person = 'S'
	r_end    = 'E'
)

type Spot struct {
	Pr int
	Pc int
}

type Person struct {
	Pr        int
	Pc        int
	Direction int
}

type QueueItem struct {
	Person Person
	Cost   int
}

func main() {
	input, err := getInputAsLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)

	person := Person{}
	end := Spot{}

	var queue []QueueItem

	for ri, line := range input {
		for ci, elem := range line {
			if elem == r_person {
				person = Person{Pr: ri, Pc: ci, Direction: right}
				queue = append(queue, QueueItem{Person: person, Cost: 0})
			}
			if elem == r_end {
				end = Spot{Pr: ri, Pc: ci}
				input[ri][ci] = r_space
			}
		}
	}
	fmt.Println("End", end)

	final_cost := 0

	var visited []Person

	for len(queue) > 0 {
		var item QueueItem

		item, queue = queue[0], queue[1:]
		if isPersonInVisited(visited, item.Person) {
			fmt.Println("Person Visited")
			continue
		} else {
			fmt.Println("Person not visited")
			visited = append(visited, item.Person)
		}
		fmt.Println("Popped: ", item.Person, item.Cost)
		// printGrid(input, item.Person)
		if item.Person.Pr == end.Pr && item.Person.Pc == end.Pc {
			final_cost = item.Cost
			break
		}

		switch item.Person.Direction {
		case up:
			if input[item.Person.Pr-1][item.Person.Pc] == r_space {
				temp_person := item.Person
				temp_person.Pr--
				temp_queue_item := QueueItem{Person: temp_person, Cost: item.Cost + 1}
				queue = append(queue, temp_queue_item)
				sort.Slice(queue[:], func(i, j int) bool {
					return queue[i].Cost < queue[j].Cost
				})
			}
		case right:
			if input[item.Person.Pr][item.Person.Pc+1] == r_space {
				temp_person := item.Person
				temp_person.Pc++
				temp_queue_item := QueueItem{Person: temp_person, Cost: item.Cost + 1}
				queue = append(queue, temp_queue_item)
				sort.Slice(queue[:], func(i, j int) bool {
					return queue[i].Cost < queue[j].Cost
				})
			}
		case down:
			if input[item.Person.Pr+1][item.Person.Pc] == r_space {
				temp_person := item.Person
				temp_person.Pr++
				temp_queue_item := QueueItem{Person: temp_person, Cost: item.Cost + 1}
				queue = append(queue, temp_queue_item)
				sort.Slice(queue[:], func(i, j int) bool {
					return queue[i].Cost < queue[j].Cost
				})
			}
		case left:
			if input[item.Person.Pr][item.Person.Pc-1] == r_space {
				temp_person := item.Person
				temp_person.Pc--
				temp_queue_item := QueueItem{Person: temp_person, Cost: item.Cost + 1}
				queue = append(queue, temp_queue_item)
				sort.Slice(queue[:], func(i, j int) bool {
					return queue[i].Cost < queue[j].Cost
				})
			}
		}
		temp_person := item.Person
		if temp_person.Direction > up {
			temp_person.Direction--
		} else {
			temp_person.Direction = left
		}
		temp_queue_item := QueueItem{Person: temp_person, Cost: item.Cost + 1000}
		queue = append(queue, temp_queue_item)
		temp_person = item.Person
		if temp_person.Direction < left {
			temp_person.Direction++
		} else {
			temp_person.Direction = up
		}
		temp_queue_item = QueueItem{Person: temp_person, Cost: item.Cost + 1000}
		queue = append(queue, temp_queue_item)
		sort.Slice(queue[:], func(i, j int) bool {
			return queue[i].Cost < queue[j].Cost
		})
		// time.Sleep(20 * time.Millisecond)
	}

	fmt.Println("Final: ", final_cost)
}

func isPersonInVisited(visited []Person, p Person) bool {
	for _, temp_p := range visited {
		if temp_p.Direction == p.Direction && temp_p.Pc == p.Pc && temp_p.Pr == p.Pr {
			return true
		}
	}
	return false
}

func printGrid(input [][]rune, p Person) {
	for ri, line := range input {
		for ci, elem := range line {
			if elem == r_wall {
				fmt.Print(string(elem))
			} else if p.Pr == ri && p.Pc == ci {
				fmt.Print(string(r_person))
			} else {
				fmt.Print(string(r_space))
			}
		}
		fmt.Println()
	}
}

func getInputAsLines() ([][]rune, error) {
	// Read in files
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
	var text [][]rune
	for file_scanner.Scan() {
		text = append(text, []rune(file_scanner.Text()))
	}
	return text, nil
}
