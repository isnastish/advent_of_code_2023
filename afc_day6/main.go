package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	// "github.com/niemeyer/golang/src/pkg/container/vector"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
		log.Println("File was close successfully.")
	}()

	var t int
	var dist int
	// time := vector.IntVector{}
	// distance := vector.IntVector{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		data := strings.Trim(strings.Split(text, ":")[1], " ")
		if strings.Contains(strings.Split(text, ":")[0], "Time") {
			num_str := strings.Join(strings.Split(data, " "), "")
			num, err := strconv.Atoi(num_str)
			if err != nil {
				panic(err)
			}
			t = num
			// for _, v := range strings.Split(data, " ") {
			// 	n, err := strconv.Atoi(v)
			// 	if err != nil {
			// 		continue
			// 	}
			// 	time.Push(n)
			// }
		} else {
			num_str := strings.Join(strings.Split(data, " "), "")
			num, err := strconv.Atoi(num_str)
			if err != nil {
				panic(err)
			}
			dist = num
			// for _, v := range strings.Split(data, " ") {
			// 	n, err := strconv.Atoi(v)
			// 	if err != nil {
			// 		continue
			// 	}
			// 	distance.Push(n)
			// }
		}
	}

	log.Println("time: ", t)
	log.Println("time: ", dist)

	// exclude not holding a button at all, and holding it for the whole time
	// total_ways := 1
	// for i := 0; i < time.Len(); i++ {
	// 	time := time[i]
	// 	distance := distance[i]
	// 	ways := 0
	// 	// amount of milliseconds to hold
	// 	for hold_ml := 1; hold_ml < time-1; hold_ml++ {
	// 		speed := hold_ml
	// 		rem_dist := time - hold_ml
	// 		total_dist := rem_dist * speed
	// 		if total_dist > distance {
	// 			ways++
	// 		}
	// 	}
	// 	total_ways *= ways
	// }

	ways := 0
	for hold_ml := 1; hold_ml < t-1; hold_ml++ {
		speed := hold_ml
		rem_dist := t - hold_ml
		total_dist := rem_dist * speed
		if total_dist > dist {
			ways++
		}
	}
	log.Println("total: ", ways)
}
