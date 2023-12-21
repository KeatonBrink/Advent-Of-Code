package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

const (
	Up = iota
	Down
	Left
	Right
)

type EndOfPath struct {
	Row, Col, StraightLineDistance, HeatLoss, InputDirection int
}

// An Item is something we manage in a priority queue.
type HeapItem struct {
	value    EndOfPath // The value of the item; arbitrary.
	priority int       // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*HeapItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*HeapItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) PushOrdered(x any) {
	n := len(*pq)
	heap.Push(pq, x)
	heap.Fix(pq, n)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *HeapItem, value EndOfPath, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
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

	var heat_map = make([][]int, len(text))

	for ind_row, row := range text {
		for _, num_as_byte := range row {
			cur_int, err := strconv.Atoi(string(num_as_byte))
			if err != nil {
				panic(err)
			}
			heat_map[ind_row] = append(heat_map[ind_row], cur_int)
		}
	}

	var visited_spots = make([][]bool, len(heat_map))

	for i := 0; i < len(heat_map); i++ {
		visited_spots[i] = make([]bool, len(heat_map[i]))
	}

	visited_spots[0][0] = true

	//Will be initially 2 items on heap
	pq := make(PriorityQueue, 2)
	pq_0_val := EndOfPath{Row: 0, Col: 1, StraightLineDistance: 2, HeatLoss: heat_map[0][1], InputDirection: Left}
	pq_1_val := EndOfPath{Row: 1, Col: 0, StraightLineDistance: 2, HeatLoss: heat_map[1][0], InputDirection: Up}

	pq[0] = &HeapItem{
		value:    pq_0_val,
		priority: pq_0_val.HeatLoss,
		index:    0}

	pq[1] = &HeapItem{
		value:    pq_1_val,
		priority: pq_1_val.HeatLoss,
		index:    1}

	heap.Init(&pq)

	best_path_val := 0
	bpd := -1
	prior := 0

	for r_counter := 0; len(pq) > 0; r_counter++ {
		// println("StartOfLoop")
		heap.Init(&pq)
		// for _, hi := range pq {
		// 	PrintHeapItem(*hi)
		// }
		// PrintSliceSliceBool(visited_spots)
		cur_heap_item := heap.Pop(&pq).(*HeapItem)
		if visited_spots[cur_heap_item.value.Row][cur_heap_item.value.Col] {
			continue
		}
		visited_spots[cur_heap_item.value.Row][cur_heap_item.value.Col] = true
		if prior > cur_heap_item.priority {
			fmt.Printf("Prior %d, current %d", prior, cur_heap_item.priority)
			panic("Uh")
		}
		prior = cur_heap_item.priority
		if cur_heap_item.value.Row == len(heat_map)-1 && cur_heap_item.value.Col == len(heat_map[0])-1 {
			best_path_val = cur_heap_item.priority
			bpd = cur_heap_item.value.InputDirection
			break
		}
		// println("Hello")
		// PrintHeapItem(*cur_heap_item)
		switch cur_heap_item.value.InputDirection {
		case Up:
			if cur_heap_item.value.StraightLineDistance <= 3 && cur_heap_item.value.Row < len(heat_map)-1 {
				next_heap_item := CopyHeapItem(*cur_heap_item)
				next_heap_item.value.Row += 1
				next_heap_item.value.StraightLineDistance++
				next_heap_item.value.HeatLoss += heat_map[next_heap_item.value.Row][next_heap_item.value.Col]
				next_heap_item.priority = next_heap_item.value.HeatLoss
				pq.PushOrdered(next_heap_item)
			}
			if cur_heap_item.value.Col < len(heat_map[0])-1 {
				next_heap_item := CopyHeapItem(*cur_heap_item)
				next_heap_item.value.Col += 1
				next_heap_item.value.StraightLineDistance = 1
				next_heap_item.value.InputDirection = Left
				next_heap_item.value.HeatLoss += heat_map[next_heap_item.value.Row][next_heap_item.value.Col]
				next_heap_item.priority = next_heap_item.value.HeatLoss
				pq.PushOrdered(next_heap_item)
			}
			if cur_heap_item.value.Col > 0 {
				next_heap_item := CopyHeapItem(*cur_heap_item)
				next_heap_item.value.Col -= 1
				next_heap_item.value.StraightLineDistance = 1
				next_heap_item.value.InputDirection = Right
				next_heap_item.value.HeatLoss += heat_map[next_heap_item.value.Row][next_heap_item.value.Col]
				next_heap_item.priority = next_heap_item.value.HeatLoss
				pq.PushOrdered(next_heap_item)
			}
		case Down:
			if cur_heap_item.value.StraightLineDistance <= 3 && cur_heap_item.value.Row > 0 {
				next_heap_item := CopyHeapItem(*cur_heap_item)
				next_heap_item.value.Row -= 1
				next_heap_item.value.StraightLineDistance++
				next_heap_item.value.HeatLoss += heat_map[next_heap_item.value.Row][next_heap_item.value.Col]
				next_heap_item.priority = next_heap_item.value.HeatLoss
				pq.PushOrdered(next_heap_item)
			}
			if cur_heap_item.value.Col < len(heat_map[0])-1 {
				next_heap_item := CopyHeapItem(*cur_heap_item)
				next_heap_item.value.Col += 1
				next_heap_item.value.StraightLineDistance = 1
				next_heap_item.value.InputDirection = Left
				next_heap_item.value.HeatLoss += heat_map[next_heap_item.value.Row][next_heap_item.value.Col]
				next_heap_item.priority = next_heap_item.value.HeatLoss
				pq.PushOrdered(next_heap_item)
			}
			if cur_heap_item.value.Col > 0 {
				next_heap_item := CopyHeapItem(*cur_heap_item)
				next_heap_item.value.Col -= 1
				next_heap_item.value.StraightLineDistance = 1
				next_heap_item.value.InputDirection = Right
				next_heap_item.value.HeatLoss += heat_map[next_heap_item.value.Row][next_heap_item.value.Col]
				next_heap_item.priority = next_heap_item.value.HeatLoss
				pq.PushOrdered(next_heap_item)
			}
		case Right:
			if cur_heap_item.value.StraightLineDistance <= 3 && cur_heap_item.value.Col > 0 {
				next_heap_item := CopyHeapItem(*cur_heap_item)
				next_heap_item.value.Col -= 1
				next_heap_item.value.StraightLineDistance++
				next_heap_item.value.HeatLoss += heat_map[next_heap_item.value.Row][next_heap_item.value.Col]
				next_heap_item.priority = next_heap_item.value.HeatLoss
				pq.PushOrdered(next_heap_item)
			}
			if cur_heap_item.value.Row < len(heat_map[0])-1 {
				next_heap_item := CopyHeapItem(*cur_heap_item)
				next_heap_item.value.Row += 1
				next_heap_item.value.StraightLineDistance = 1
				next_heap_item.value.InputDirection = Up
				next_heap_item.value.HeatLoss += heat_map[next_heap_item.value.Row][next_heap_item.value.Col]
				next_heap_item.priority = next_heap_item.value.HeatLoss
				pq.PushOrdered(next_heap_item)
			}
			if cur_heap_item.value.Row > 0 {
				next_heap_item := CopyHeapItem(*cur_heap_item)
				next_heap_item.value.Row -= 1
				next_heap_item.value.StraightLineDistance = 1
				next_heap_item.value.InputDirection = Down
				next_heap_item.value.HeatLoss += heat_map[next_heap_item.value.Row][next_heap_item.value.Col]
				next_heap_item.priority = next_heap_item.value.HeatLoss
				pq.PushOrdered(next_heap_item)
			}
		case Left:
			if cur_heap_item.value.StraightLineDistance <= 3 && cur_heap_item.value.Col < len(heat_map[0])-1 {
				next_heap_item := CopyHeapItem(*cur_heap_item)
				next_heap_item.value.Col += 1
				next_heap_item.value.StraightLineDistance++
				next_heap_item.value.HeatLoss += heat_map[next_heap_item.value.Row][next_heap_item.value.Col]
				next_heap_item.priority = next_heap_item.value.HeatLoss
				pq.PushOrdered(next_heap_item)
			}
			if cur_heap_item.value.Row < len(heat_map[0])-1 {
				next_heap_item := CopyHeapItem(*cur_heap_item)
				next_heap_item.value.Row += 1
				next_heap_item.value.StraightLineDistance = 1
				next_heap_item.value.InputDirection = Up
				next_heap_item.value.HeatLoss += heat_map[next_heap_item.value.Row][next_heap_item.value.Col]
				next_heap_item.priority = next_heap_item.value.HeatLoss
				pq.PushOrdered(next_heap_item)
			}
			if cur_heap_item.value.Row > 0 {
				next_heap_item := CopyHeapItem(*cur_heap_item)
				next_heap_item.value.Row -= 1
				next_heap_item.value.StraightLineDistance = 1
				next_heap_item.value.InputDirection = Down
				next_heap_item.value.HeatLoss += heat_map[next_heap_item.value.Row][next_heap_item.value.Col]
				next_heap_item.priority = next_heap_item.value.HeatLoss
				pq.PushOrdered(next_heap_item)
			}
		}
	}

	// PrintSliceSliceBool(visited_spots)

	fmt.Printf("Best Path Val: %d %d\n", best_path_val, bpd)
	println(Right)
}

func CopyHeapItem(hi HeapItem) *HeapItem {
	var new_hi HeapItem
	new_hi.value.Row = hi.value.Row
	new_hi.value.Col = hi.value.Col
	new_hi.value.StraightLineDistance = hi.value.StraightLineDistance
	new_hi.value.HeatLoss = hi.value.HeatLoss
	new_hi.value.InputDirection = hi.value.InputDirection
	new_hi.priority = hi.priority
	new_hi.index = hi.index
	return &new_hi
}

func PrintHeapItem(hi HeapItem) {
	fmt.Printf("Row %d Col %d SLD %d HeatLoss %d ID %d Priority %d index %d\n", hi.value.Row, hi.value.Col, hi.value.StraightLineDistance, hi.value.HeatLoss, hi.value.InputDirection, hi.priority, hi.index)
}

func PrintSliceSliceBool(ssb [][]bool) {
	for _, line := range ssb {
		for _, b_val := range line {
			if b_val {
				print("T")
			} else {
				print("F")
			}
		}
		println()
	}
}
