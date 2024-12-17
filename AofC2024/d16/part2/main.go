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
	Path   []Person
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

	prev_cost := 0

	var best_paths_spots []Person

	for len(queue) > 0 {
		var item QueueItem

		item, queue = queue[0], queue[1:]
		if item.Cost > prev_cost {
			fmt.Println(item.Cost)
			prev_cost = item.Cost
		}
		if item.Person.Pr == end.Pr && item.Person.Pc == end.Pc {
			final_cost = item.Cost
			best_paths_spots = addUniquePathToSpots(best_paths_spots, item.Path)
		}

		if item.Cost > final_cost && final_cost > 0 {
			break
		}
		if !isPersonInVisited(item.Path, item.Person) {
			item.Path = append(item.Path, item.Person)
		} else {
			continue
		}

		switch item.Person.Direction {
		case up:
			if input[item.Person.Pr-1][item.Person.Pc] == r_space {
				temp_person := item.Person
				temp_person.Pr--
				temp_queue_item := QueueItem{Person: temp_person, Cost: item.Cost + 1}
				// fmt.Println("temp path", temp_queue_item.Path)
				// fmt.Println(item.Path)
				temp_queue_item.Path = copyPath(temp_queue_item.Path, item.Path)
				// fmt.Println("temp path2", temp_queue_item.Path)
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
				temp_queue_item.Path = copyPath(temp_queue_item.Path, item.Path)
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
				temp_queue_item.Path = copyPath(temp_queue_item.Path, item.Path)
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
				temp_queue_item.Path = copyPath(temp_queue_item.Path, item.Path)
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
		temp_queue_item.Path = copyPath(temp_queue_item.Path, item.Path)
		queue = append(queue, temp_queue_item)
		temp_person = item.Person
		if temp_person.Direction < left {
			temp_person.Direction++
		} else {
			temp_person.Direction = up
		}
		temp_queue_item = QueueItem{Person: temp_person, Cost: item.Cost + 1000}
		temp_queue_item.Path = copyPath(temp_queue_item.Path, item.Path)
		queue = append(queue, temp_queue_item)
		sort.Slice(queue[:], func(i, j int) bool {
			return queue[i].Cost < queue[j].Cost
		})
		// time.Sleep(20 * time.Millisecond)
	}

	fmt.Println("Final: ", len(best_paths_spots))
}

func copyPath(dst []Person, src []Person) []Person {
	for _, spot := range src {
		dst = append(dst, spot)
	}
	return dst
}

func addUniquePathToSpots(best_paths_spots []Person, new_path []Person) []Person {
	for _, new_spot := range new_path {
		// fmt.Println(new_spot)
		is_new := true
		for _, old_spot := range best_paths_spots {
			// fmt.Println(old_spot)
			if old_spot.Pr == new_spot.Pr && old_spot.Pc == new_spot.Pc {
				is_new = false
			}
		}
		if is_new {
			best_paths_spots = append(best_paths_spots, new_spot)
		}
	}
	return best_paths_spots
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
	f, err := os.Open("test2.txt")
	// f, err := os.Open("test.txt")
	// f, err := os.Open("input.txt")
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
