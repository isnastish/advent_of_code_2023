package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		fmt.Println("Closing the file.")
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	const red_total = 12
	const green_total = 13
	const blue_total = 14

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ":")[1]

		green_min := 0
		red_min := 0
		blue_min := 0

		var sets = strings.Split(data, ";")
		for i := 0; i < len(sets); i++ {
			cube_sets := strings.Split(sets[i], ",")
			for j := 0; j < len(cube_sets); j++ {
				pair := strings.Split(cube_sets[j], " ")
				cubes_count, _ := strconv.Atoi(pair[1])
				if strings.Compare(pair[2], "blue") == 0 {
					blue_min = int(math.Max(float64(blue_min), float64(cubes_count)))
				} else if strings.Compare(pair[2], "red") == 0 {
					red_min = int(math.Max(float64(red_min), float64(cubes_count)))
				} else if strings.Compare(pair[2], "green") == 0 {
					green_min = int(math.Max(float64(green_min), float64(cubes_count)))
				}
			}
		}
		sum += (green_min * red_min * blue_min)
	}
	fmt.Printf("sum: %d\n", sum)
}
