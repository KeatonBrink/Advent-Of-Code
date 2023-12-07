package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	card string
	bid  int
}

const (
	total_cards = 13
)

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

	grouped_by_win_type := make([][]Pair, 7)

	return_val := 0

	for _, line := range text {
		var cards_from_line string
		var bid_from_line int
		s, err := fmt.Sscanf(line, "%s %d", &cards_from_line, &bid_from_line)
		if err != nil {
			panic(err)
		}
		if s != 2 {
			panic("Not enough found in scan")
		}
		grouped_by_win_type[FindWinType(cards_from_line)] = append(grouped_by_win_type[FindWinType(cards_from_line)], Pair{card: cards_from_line, bid: bid_from_line})
		// fmt.Printf("%s added to group: %d \n", cards_from_line, FindWinType(cards_from_line))
	}

	card_count := 1
	for win_type, slice_of_win_type := range grouped_by_win_type {
		for current_card_index := 0; current_card_index < len(grouped_by_win_type[win_type])-1; current_card_index++ {
			for potential_switch_card_index := current_card_index + 1; potential_switch_card_index < len(grouped_by_win_type[win_type]); potential_switch_card_index++ {
				if IsLeftHandBetterThanRight(slice_of_win_type[current_card_index].card, slice_of_win_type[potential_switch_card_index].card) {
					slice_of_win_type[potential_switch_card_index], slice_of_win_type[current_card_index] = slice_of_win_type[current_card_index], slice_of_win_type[potential_switch_card_index]
				}
			}
			return_val += card_count * slice_of_win_type[current_card_index].bid
			// fmt.Printf("Cards %s rank %d with bid %d, new ret_val %d\n", slice_of_win_type[current_card_index].card, card_count, slice_of_win_type[current_card_index].bid, return_val)
			card_count++
		}
		// Getting the last element
		i := len(slice_of_win_type) - 1
		if i >= 0 {
			return_val += card_count * slice_of_win_type[i].bid
			// fmt.Printf("Cards %s rank %d with bid %d, new ret_val %d\n", slice_of_win_type[i].card, card_count, slice_of_win_type[i].bid, return_val)
			card_count++
		}
	}

	println("Valid attempts", return_val)
}

func FindWinType(cards string) int {
	if IsFiveOfKind(cards) {
		return 6
	} else if IsFourOfKind(cards) {
		return 5
	} else if IsFullHouse(cards) {
		return 4
	} else if IsThreeOfKind(cards) {
		return 3
	} else if IsTwoPair(cards) {
		return 2
	} else if IsOnePair(cards) {
		return 1
	}
	return 0
}

func IsLeftHandBetterThanRight(left string, right string) bool {
	for char_index := range left {
		if CharToInt(left[char_index]) > CharToInt(right[char_index]) {
			// fmt.Printf("Left hand is better: left %s right %s \n", left, right)
			return true
		} else if CharToInt(left[char_index]) < CharToInt(right[char_index]) {
			// fmt.Printf("Right hand is better: left %s right %s \n", left, right)
			return false
		}
	}
	println(string(left[3]))
	println(CharToInt(left[3]))
	println(string(right[3]))
	println(CharToInt(right[3]))
	fmt.Printf("Hand left %s, Hand right %s \n", left, right)
	panic("I think all hands should be unique?")
}

func CharToInt(card byte) int {
	switch {
	case card >= '2' && card <= '9':
		return int(card - '1')
	case card == 'T':
		return 9
	case card == 'J':
		return 0
	case card == 'Q':
		return 10
	case card == 'K':
		return 11
	case card == 'A':
		return 12
	default:
		panic(fmt.Sprintf("Invalid byte %b", card))
	}
}

func IsFiveOfKind(cards string) bool {
	kind := 'a'
	for _, char := range cards {
		if char == 'J' {
			continue
		}
		if kind == 'a' {
			kind = char
		} else if kind != char {
			return false
		}
	}
	return true
}

func IsFourOfKind(cards string) bool {
	potential_cards := make([]int, total_cards)
	for _, char := range cards {
		integer_card := CharToInt(byte(char))
		potential_cards[integer_card] += 1
	}

	for i := 1; i < len(potential_cards); i++ {
		if potential_cards[i]+potential_cards[0] >= 4 {
			return true
		}
	}
	return false
}

type HighCard struct {
	Card  int
	Count int
}

func IsFullHouse(cards string) bool {
	potential_cards := make([]int, total_cards)
	for _, char := range cards {
		integer_card := CharToInt(byte(char))
		potential_cards[integer_card] += 1
	}
	first := HighCard{Card: -1, Count: -1}
	second := HighCard{Card: -1, Count: -1}
	for current_card_as_int, current_card_count := range potential_cards {
		if current_card_as_int == 0 {
			continue
		}
		if first.Count < current_card_count {
			second.Count = first.Count
			second.Card = first.Count
			first.Count = current_card_count
			first.Card = current_card_as_int
		} else if second.Count < current_card_count {
			second.Count = current_card_count
			second.Card = current_card_as_int
		}
	}
	if first.Count+second.Count+potential_cards[0] == 5 {
		return true
	}
	return false
}

func IsThreeOfKind(cards string) bool {
	potential_cards := make([]int, total_cards)
	for _, char := range cards {
		integer_card := CharToInt(byte(char))
		potential_cards[integer_card] += 1
	}

	for i := 1; i < len(potential_cards); i++ {
		if potential_cards[i]+potential_cards[0] == 3 {
			return true
		}
	}
	return false
}

func IsTwoPair(cards string) bool {
	potential_cards := make([]int, total_cards)
	for _, char := range cards {
		integer_card := CharToInt(byte(char))
		potential_cards[integer_card] += 1
	}
	first := HighCard{Card: -1, Count: -1}
	second := HighCard{Card: -1, Count: -1}
	for current_card_as_int, current_card_count := range potential_cards {
		if current_card_as_int == 0 {
			continue
		}
		if first.Count < current_card_count {
			second.Count = first.Count
			second.Card = first.Count
			first.Count = current_card_count
			first.Card = current_card_as_int
		} else if second.Count < current_card_count {
			second.Count = current_card_count
			second.Card = current_card_as_int
		}
	}
	if first.Count+second.Count+potential_cards[0] == 4 {
		return true
	}
	return false
}

func IsOnePair(cards string) bool {
	potential_cards := make([]int, total_cards)
	for _, char := range cards {
		integer_card := CharToInt(byte(char))
		potential_cards[integer_card] += 1
	}
	for i, card_int := range potential_cards {
		if i == 0 {
			continue
		}
		if card_int+potential_cards[0] == 2 {
			return true
		}
	}
	return false
}
