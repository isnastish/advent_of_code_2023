package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/niemeyer/golang/src/pkg/container/vector"
)

type Card struct {
	CardNumber int
	MatchCount int
}

func getMatchCount(winning_numbers string, numbers_i_have string) int {
	match_count := 0
	for _, wn := range strings.Split(winning_numbers, " ") {
		first, err := strconv.Atoi(wn)
		if err != nil {
			continue
		}
		for _, nh := range strings.Split(numbers_i_have, " ") {
			second, err := strconv.Atoi(nh)
			if err != nil {
				continue
			}
			if first == second {
				match_count++
				break
			}
		}
	}
	return match_count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
		log.Println("File was close succesfully!")
	}()

	cards := vector.Vector{}
	card_number := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(strings.Split(line, ":")[1], "|")
		winning_numbers := strings.Trim(data[0], " ")
		numbers_i_have := strings.Trim(data[1], " ")
		cards.Push(Card{card_number, getMatchCount(winning_numbers, numbers_i_have)})
		card_number++
	}

	total_cards := uint32(0)
	index := 0
	cards_stack := vector.Vector{}
	for index < cards.Len() {
		taken_from_stack := false
		var card Card
		if cards_stack.Len() != 0 {
			card = cards_stack.Last().(Card)
			taken_from_stack = true
		} else {
			card = cards[index].(Card)
		}

		if taken_from_stack {
			cards_stack.Pop()
		} else {
			index++
			total_cards++
		}

		if card.MatchCount > 0 {
			for j := 0; j < card.MatchCount; j++ {
				cards_stack.Push(cards[j+card.CardNumber])
				total_cards++
			}
		}
	}
	log.Println("total cards: ", total_cards)
}
