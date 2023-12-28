package main

import (
	"bufio"
	"github.com/niemeyer/golang/src/pkg/container/vector"
	"log"
	"os"
	"strconv"
	"strings"
)

type HandBidPair struct {
	Hand string
	Bid  int
}

type HandKind int

const (
	HandKind_FiveOfKind  HandKind = 1
	HandKind_FourOfKind           = 2
	HandKind_FullHouse            = 3
	HandKind_ThreeOfKind          = 4
	HandKind_TwoPair              = 5
	HandKind_OnePair              = 6
	HandKind_HighCard             = 7
)

func constructLabelMap(hand string) map[string]int {
	label_map := make(map[string]int)
	for _, label := range hand {
		_, ok := label_map[string(label)]
		if ok {
			label_map[string(label)] += 1
		} else {
			label_map[string(label)] = 1
		}
	}
	return label_map
}

// Joker is "J" letter, which can pretend to be any letter to make the hand strongest.
func detectHandKindWithJoker(hand string) HandKind {
	label_map := constructLabelMap(hand)
	number_of_Js, ok := label_map["J"]
	if !ok { // normal hand, "J" is not present
		if len(label_map) == 1 { // AAAAA
			return HandKind_FiveOfKind
		} else if len(label_map) == 2 { // AA8AA | 23332
			for _, v := range label_map {
				if v == 4 {
					return HandKind_FourOfKind
				}
			}
			return HandKind_FullHouse
		} else if len(label_map) == 3 { // ThreeOfAKind | TwoPair
			for _, v := range label_map {
				if v == 3 {
					return HandKind_ThreeOfKind
				}
			}
			return HandKind_TwoPair
		} else if len(label_map) == 4 { // One Pair A23A4
			return HandKind_OnePair
		} else { // High card 12345
			return HandKind_HighCard
		}
	} else {
		if len(label_map) == 1 { // all Js
			return HandKind_FiveOfKind
		} else if len(label_map) == 2 { // AAJAA, J333J, JJJJ8
			return HandKind_FiveOfKind
		} else if len(label_map) == 3 {
			// T55J5, either ThreeOfKind | TwoPair
			// Which means that can either become FourOfKind | ThreeOfKind
			// TTKKJ -> TTKKK
			// TTKJJ -> TTKTT
			if number_of_Js == 1 {
				for _, count := range label_map {
					if count == 3 {
						return HandKind_FourOfKind
					}
				}
				return HandKind_FullHouse
			}
			return HandKind_FourOfKind
		} else if len(label_map) == 4 { // A23AJ -> A23AA (ThreeOfKind)
			// one of them is J
			// AA14J
			return HandKind_ThreeOfKind
		} else {
			// 2346J
			return HandKind_OnePair
		}
	}
}

func detectHandKind(hand string) HandKind {
	label_map := constructLabelMap(hand)

	if len(label_map) == 1 { // AAAAA
		return HandKind_FiveOfKind
	} else if len(label_map) == 2 { // AA8AA | 23332
		for _, v := range label_map {
			if v == 4 {
				return HandKind_FourOfKind
			}
		}
		return HandKind_FullHouse
	} else if len(label_map) == 3 { // ThreeOfAKind | TwoPair
		for _, v := range label_map {
			if v == 3 {
				return HandKind_ThreeOfKind
			}
		}
		return HandKind_TwoPair
	} else if len(label_map) == 4 { // One Pair A23A4
		return HandKind_OnePair
	} else { // High card 12345
		return HandKind_HighCard
	}
}

func printHand(hand *vector.Vector, hand_kind string) {
	for i, v := range *hand {
		log.Printf("%s[%d]: %v\n", hand_kind, i, v)
	}
}

func compareHands(hand_a string, hand_b string, labels []string) int {
	// Each hand has a length of 5
	// 0  - hands are equal
	// 1  - first hand is stronger than the second one
	// -1 - second hand is stronger than the first one
	for i := 0; i < 5; i++ {
		a := string(hand_a[i])
		b := string(hand_b[i])
		if a == b {
			continue
		}
		a_index := -1
		b_index := -1
		for j, l := range labels {
			if b_index == -1 && b == l {
				b_index = j
			}
			if a_index == -1 && a == l {
				a_index = j
			}

			if a_index != -1 && b_index != -1 {
				break
			}
		}
		if a_index < b_index { // hand a is stronger
			return 1
		} else if b_index < a_index { // hand b is stronger
			return -1
		}
	}
	return 0
}

func sortHand(hand_list *vector.Vector) {
	labels := []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}
	// labels := []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

	for i := 0; i < hand_list.Len(); i++ {
		if i == 0 {
			continue
		}
		cur_hand := (*hand_list)[i].(HandBidPair).Hand
		prev_hand := (*hand_list)[i-1].(HandBidPair).Hand
		result := compareHands(cur_hand, prev_hand, labels)
		if result == 1 { // cur_hand is stronger than prev_hand, swap is required
			hand_list.Swap(i-1, i)
			for j := i - 1; j > 0; j-- {
				prev_hand = (*hand_list)[j-1].(HandBidPair).Hand
				result = compareHands(cur_hand, prev_hand, labels)
				if result == 1 {
					hand_list.Swap(j, j-1)
				}
			}
		}
	}
}

func main() {
	file_name := "input.txt"
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
		log.Println("Successfully closed the file.")
	}()

	five_of_kind := vector.Vector{}
	four_of_kind := vector.Vector{}
	full_house := vector.Vector{}
	three_of_kind := vector.Vector{}
	two_pair := vector.Vector{}
	one_pair := vector.Vector{}
	high_card := vector.Vector{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		bid_str := strings.Trim(strings.Split(text, " ")[1], " ")
		bid, err := strconv.Atoi(bid_str)
		if err != nil {
			panic(err)
		}
		hand_str := strings.Trim(strings.Split(text, " ")[0], " ")
		// hand_kind := detectHandKind(hand_str)
		hand_kind := detectHandKindWithJoker(hand_str)
		switch hand_kind {
		case HandKind_FiveOfKind:
			five_of_kind.Push(HandBidPair{hand_str, bid})
		case HandKind_FourOfKind:
			four_of_kind.Push(HandBidPair{hand_str, bid})
		case HandKind_FullHouse:
			full_house.Push(HandBidPair{hand_str, bid})
		case HandKind_ThreeOfKind:
			three_of_kind.Push(HandBidPair{hand_str, bid})
		case HandKind_TwoPair:
			two_pair.Push(HandBidPair{hand_str, bid})
		case HandKind_OnePair:
			one_pair.Push(HandBidPair{hand_str, bid})
		case HandKind_HighCard:
			high_card.Push(HandBidPair{hand_str, bid})
		}
	}

	sortHand(&five_of_kind)
	sortHand(&four_of_kind)
	sortHand(&full_house)
	sortHand(&three_of_kind)
	sortHand(&two_pair)
	sortHand(&one_pair)
	sortHand(&high_card)

	final_hand := vector.Vector{}

	final_hand.AppendVector(&five_of_kind)
	final_hand.AppendVector(&four_of_kind)
	final_hand.AppendVector(&full_house)
	final_hand.AppendVector(&three_of_kind)
	final_hand.AppendVector(&two_pair)
	final_hand.AppendVector(&one_pair)
	final_hand.AppendVector(&high_card)

	total := int64(0)
	rank := 1
	for i := final_hand.Len() - 1; i >= 0; i-- {
		bid := final_hand[i].(HandBidPair).Bid
		total += int64(bid * rank)
		rank++
	}
	log.Println("total: ", total)
}
