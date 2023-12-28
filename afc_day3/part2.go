package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func parsePrevInt(arr []rune, x, y, width int) int {
	i := x
	for ; unicode.IsDigit(arr[y*width + i]); i-- {}
	result, err := strconv.Atoi(string(arr[(y*width + i+1) : (y*width + x+1)]))
	if err != nil { panic(err) }
	return result
}

func parseNextInt(arr []rune, x, y, width int) int {
	i := x
	for ; unicode.IsDigit(arr[y*width + i]); i++ {}
	result, err := strconv.Atoi(string(arr[(y*width + x) : (y*width + i)]))
	if err != nil { panic(err) }
	return result
}

func exploreAroundStar(arr []rune, x, y, width, height int) int {
	/*
	      . . . 
		  .	* .
		  .	. .
	*/

	var minx int
	var miny int
	var maxx int
	var maxy int

	if x-1 >= 0 {
		minx = x - 1
	} else {
		minx = x
	}

	if y-1 >= 0 {
		miny = y - 1
	} else {
		miny = y
	}

	if x+1 < width {
		maxx = x+1
	} else {
		maxx = x
	} 

	if y+1 < height {
		maxy = y+1
	} else {
		maxy = y
	}

	first_value := 0
	for i := miny; i <= maxy; i++ {
		for j := minx; j <= maxx; j++ {
			fmt.Printf("%c ",arr[i*width + j])
			if unicode.IsDigit(arr[i*width + j]) {
				s := j
				for ; s > 0 && unicode.IsDigit(arr[i*width + s]); {
					s--
				}
				if !unicode.IsDigit(arr[i*width + s]) {
					s++
				}

				var number []rune
				k := s
				for ; unicode.IsDigit(arr[i*width + k]); k++ {
					number = append(number, arr[i*width + k])
				}
				r, err := strconv.Atoi(string(number))
				if err != nil { panic(err) }

				if first_value == 0 {
					first_value = r
					fmt.Printf("first value: %d\n", first_value)
					j = k
				} else{
					fmt.Println("first: ", first_value, "second: ", r)
					return first_value * r
				}
			}
		}
	}
	return 0
}

func Part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	
	defer func () {
		err := file.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("File was closed succesfully.")
	}()

	arr := make([]rune, 0)
	
	// read file contents into rune.
	scannar := bufio.NewScanner(file)
	for scannar.Scan() {
		line := scannar.Text()
		for i := 0; i < len(line); i++ {
			arr = append(arr, rune(line[i]))
		}
	}

	const width = 140
	const height = 140
	var sum uint64 = 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			i := y*width + x
			if arr[i] == '*'{
				sum += uint64(exploreAroundStar(arr, x, y, width, height))

				// lx := x - 1
				// rx := x + 1
				// ty := y - 1
				// by := y + 1

				// if lx >= 0 && rx < width && ty >= 0 && by < height {
				// 	if unicode.IsDigit(arr[y*width + lx]) { // 235*
				// 		prev := parsePrevInt(arr, lx, y, width)
				// 		if unicode.IsDigit(arr[y*width + rx]) { // 235*145
				// 			next := parseNextInt(arr, rx, y, width)
				// 			sum += uint64(prev*next)
				// 		} else if unicode.IsDigit(arr[ty*width + x]) { 
				// 			//  ..3..
				// 			// 234*
				// 			start := x
				// 			end := x 
				// 			for ; unicode.IsDigit(arr[ty*width + start]); start--{}
				// 			for ; unicode.IsDigit(arr[ty*width + end]); end++{}
				// 			result, err := strconv.Atoi(string(arr[(ty*width + start+1) : (ty*width + end)]))
				// 			if err != nil { panic(err) }
				// 			sum += uint64(prev * result)
				// 		} else if unicode.IsDigit(arr[ty*width + rx]) {
				// 			//  ...534
				// 			// 234*
				// 			result := parseNextInt(arr, rx, ty, width)
				// 			sum += uint64(prev * result)
				// 		} else if unicode.IsDigit(arr[ty*width + lx]) {
				// 			// 548..
				// 			// 234*
				// 			result := parsePrevInt(arr, lx, ty, width)
				// 			sum += uint64(prev * result)
				// 		} else if unicode.IsDigit(arr[by*width + x]) {
				// 			// 234*
				// 			//  ..4..
				// 			start := x
				// 			end := x 
				// 			for ; unicode.IsDigit(arr[by*width + start]); start--{}
				// 			for ; unicode.IsDigit(arr[by*width + end]); end++{}
				// 			result, err := strconv.Atoi(string(arr[(by*width + start+1) : (by*width + end)]))
				// 			if err != nil { panic(err) }
				// 			sum += uint64(prev * result)
				// 		} else if unicode.IsDigit(arr[by*width + rx]) {
				// 			// 234*
				// 			//  ...2345
				// 			result := parseNextInt(arr, rx, by, width)
				// 			sum += uint64(prev * result)
				// 		} else if unicode.IsDigit(arr[by*width + lx]) {
				// 			// 234*
				// 			// ..2...
				// 			result := parsePrevInt(arr, lx, by, width)
				// 			sum += uint64(prev * result)
				// 		}
				// 	}
				// }
			}
		}
	}
	fmt.Println(sum)
}