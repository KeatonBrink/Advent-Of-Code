package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type JunctionBox struct {
	x float64
	y float64
	z float64
}

func (jb1 JunctionBox) IsSame(jb2 JunctionBox) bool {
	return jb1.x == jb2.x && jb1.y == jb2.y && jb1.z == jb2.z
}

func (jb JunctionBox) Print() {
	fmt.Println(jb, jb.x, jb.y, jb.z)
}
func (jb JunctionBox) Stringify() string {
	return fmt.Sprint(jb.x, jb.y, jb.z)
}

type JunctionBoxConnection struct {
	a *JunctionBox
	b *JunctionBox
}

func (jbc JunctionBoxConnection) Distance() float64 {
	x_squared := math.Pow(jbc.a.x-jbc.b.x, 2)
	y_squared := math.Pow(jbc.a.y-jbc.b.y, 2)
	z_squared := math.Pow(jbc.a.z-jbc.b.z, 2)
	return math.Pow(x_squared+y_squared+z_squared, 0.5)
}

func (jbc JunctionBoxConnection) Print() {
	fmt.Println(jbc.a.Stringify(), ", ", jbc.b.Stringify(), ":::", jbc.Distance())
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
	var junction_boxes []JunctionBox
	for _, line := range input {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		z, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		junction_boxes = append(junction_boxes, JunctionBox{
			float64(x),
			float64(y),
			float64(z)})
	}
	fmt.Println("Junction Boxes Created")
	start_time := time.Now()
	var junction_pairs []JunctionBoxConnection
	for ind1 := 0; ind1 < len(junction_boxes)-1; ind1++ {
		if ind1%100 == 0 {
			fmt.Println("Junction Box Complete: ", ind1, " Time to complete: ", time.Since(start_time))
		}
		for ind2 := ind1 + 1; ind2 < len(junction_boxes); ind2++ {
			// fmt.Println(ind1, ind2)
			jbc := JunctionBoxConnection{&junction_boxes[ind1], &junction_boxes[ind2]}
			jbc_dist := jbc.Distance()

			// Using Golang efficient search
			idx := sort.Search(len(junction_pairs), func(i int) bool {
				return junction_pairs[i].Distance() >= jbc_dist
			})
			junction_pairs = append(junction_pairs[:idx], append([]JunctionBoxConnection{jbc}, junction_pairs[idx:]...)...)
		}

	}
	fmt.Println("Junction Box Connections completed")
	pairs := 0
	total_junctions := len(input)
	var circuits [][]*JunctionBox
	for _, jbc := range junction_pairs {
		// fmt.Println(jbc.a, jbc.b)
		// if pairs == max_pairs {
		// 	break
		// }
		a_ind := -1
		b_ind := -1
		for ind, circuit := range circuits {
			for _, junction_box := range circuit {
				// Probably could have compared address, but this works as well
				if junction_box.IsSame(*jbc.a) {
					a_ind = ind
				}
				if junction_box.IsSame(*jbc.b) {
					b_ind = ind
				}
			}
		}
		if a_ind == -1 {
			if b_ind == -1 {
				circuits = append(circuits, []*JunctionBox{jbc.a, jbc.b})
			} else if b_ind != -1 {
				circuits[b_ind] = append(circuits[b_ind], jbc.a)
			}
		} else if a_ind != -1 {
			if b_ind == -1 {
				circuits[a_ind] = append(circuits[a_ind], jbc.b)
			} else if b_ind != -1 && a_ind != b_ind {
				// Two circuits need to be combined
				circuits[a_ind] = append(circuits[a_ind], circuits[b_ind]...)
				circuits = append(circuits[:b_ind], circuits[b_ind+1:]...)
			}
		}
		if len(circuits) == 1 && len(circuits[0]) == total_junctions {
			fmt.Println("One Circuit created")
			jbc.Print()
			fmt.Println("Answer: ", jbc.a.x*jbc.b.x)
			break
		}
		pairs++
	}
	fmt.Println("Circuits Created")
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
