package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open a file.")
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		rstr := []rune(line)
		result := make([]int, 2)
		digits_found := 0
		rstr_len := len(rstr)
		for i := 0; i < len(rstr); i++ {
			//fmt.Printf("%c\n", rstr[i])
			rem := rstr_len - i - 1

			if unicode.IsDigit(rstr[i]) {
				digit := rstr[i] - 48
				if digits_found == 0 {
					result[0] = int(digit)
				} else {
					result[1] = int(digit)
				}
				digits_found += 1
			} else if rstr[i] == 'o' && rem >= 2 { // one
				if rstr[i + 1] == 'n' && rstr[i + 2] == 'e' {
					//i += 2
					if digits_found == 0{
						result[0] = 1
					} else {
						result[1] = 1
					}
					digits_found += 1
				}
			} else if rstr[i] == 't' { // three
				if rem >= 2 {
					if rstr[i + 1] == 'w' && rstr[i + 2] == 'o' {
						// i += 2
						if digits_found == 0{
							result[0] = 2
						} else {
							result[1] = 2
						}
						digits_found += 1
					} else if rstr[i + 1] == 'h' && rem >= 4 {
						if rstr[i + 2] == 'r' && rstr[i + 3] == 'e' && rstr[i + 4] == 'e' {
							i += 4
							if digits_found == 0{
								result[0] = 3
							} else {
								result[1] = 3
							}
							digits_found += 1
						}
					}
				}
			} else if rstr[i] == 'f' && rem >= 3 { // four, five
				if rstr[i + 1] == 'o' && rstr[i + 2] == 'u' && rstr[i + 3] == 'r' {
					// i += 3
					if digits_found == 0{
						result[0] = 4
					} else {
						result[1] = 4
					}
					digits_found += 1
				} else if rstr[i + 1] == 'i' && rstr[i + 2] == 'v' && rstr[i + 3] == 'e' {
					//i += 3
					if digits_found == 0{
						result[0] = 5
					} else {
						result[1] = 5
					}
					digits_found += 1
				}
			} else if rstr[i] == 's' {
				if rem >= 2 {
					if rstr[i + 1] == 'i' && rstr[i + 2] == 'x' {
						//i += 2
						if digits_found == 0{
							result[0] = 6
						} else {
							result[1] = 6
						}
						digits_found += 1
					} else if rstr[i + 1] == 'e' && rem >= 4 {
						if rstr[i + 2] == 'v' && rstr[i + 3] == 'e' && rstr[i + 4] == 'n' {
							//i += 4
							if digits_found == 0{
								result[0] = 7
							} else {
								result[1] = 7
							}
							digits_found += 1
						}
					}
				}
			} else if rstr[i] == 'e' && rem >= 4{ // eight
				if rstr[i + 1] == 'i' && rstr[i + 2] == 'g' && rstr[i + 3] == 'h' && rstr[i + 4] == 't' {
					//i += 4
					if digits_found == 0{
						result[0] = 8
					} else {
						result[1] = 8
					}
					digits_found += 1
				}
			} else if rstr[i] == 'n' && rem >= 3 { // nine
				if rstr[i + 1] == 'i' && rstr[i + 2] == 'n' && rstr[i + 3] == 'e' {
					//i += 3
					if digits_found == 0{
						result[0] = 9
					} else {
						result[1] = 9
					}
					digits_found += 1
				}
			}
		}
		if digits_found == 1 {
			line_sum := 10 * result[0] + result[0]
			fmt.Println(line_sum)
			sum += line_sum
		}else {
			line_sum := 10 * result[0] + result[1] 
			fmt.Println(line_sum)
			sum += line_sum
		}
	}
	fmt.Printf("Sum: %d", sum)
}