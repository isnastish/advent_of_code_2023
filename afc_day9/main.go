package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func computeNextHistoryValue(seq []int) int {
	slice := [][]int{
		seq,
	}

	all_zeros := false
	for i := 0; i < len(slice); i++ {
		new_slice := []int{}
		zeros_count := 0
		for k := 0; k < len(slice[i])-1; k++ {
			diff := slice[i][k+1] - slice[i][k]
			new_slice = append(new_slice, diff)
			if diff == 0 {
				zeros_count++
			}
		}
		slice = append(slice, new_slice)
		if zeros_count == len(slice[i])-1 {
			all_zeros = true
			break
		}
	}

	next_history := 0
	// NOTE(alx): Uncomment our for the first part.
	// if all_zeros {
	// 	for i := len(slice) - 1; i > 0; i-- {
	// 		next_history = next_history + slice[i-1][len(slice[i-1])-1]
	// 	}
	// }

	// second part
	if all_zeros {
		for i := len(slice) - 1; i > 0; i-- {
			next_history = slice[i-1][0] - next_history
		}
	}
	return next_history
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
		log.Println("Successfully closed the file.")
	}()

	total_history := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		history := []int{}
		for _, v := range strings.Split(text, " ") {
			n, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			history = append(history, n)
		}
		next_history := computeNextHistoryValue(history)
		total_history += next_history
	}
	log.Println("total history: ", total_history)
}
