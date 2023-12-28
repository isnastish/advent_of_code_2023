package main

// import (
// 	// "bufio"
// 	// "fmt"
// 	// "os"
// 	// "strconv"

// 	// //"strings"
// 	// "unicode"
// 	// //"strconv"
// )

func main() {
	Part2()
	/*
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	
	defer func () {
		err := file.Close()
		if err != nil {
			panic(err)
		}
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

	sum := 0
	// go through each character
	const width = 140
	const height = 140
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			xi := y*width + x
			if unicode.IsDigit(arr[xi]) {
				//fmt.Printf("%c\n", arr[xi])
				valid_number := false
				// check whether the previous character was a symbol
				if x != 0 {
					pi := y*width + x-1
					if !unicode.IsDigit(arr[pi]) && arr[pi] != '.' {
						valid_number = true
					}
				}
				k := x
				for ; (k < width) && unicode.IsDigit(arr[y * width + k]); k++ {
					nc := '.'
					if y < (height - 1) {
						nc = arr[(y + 1) * width + k]
					}
					pc := '.'
					if y != 0 {
						pc = arr[(y - 1) * width + k]
					}
					if !valid_number {
						if (!unicode.IsDigit(nc) && nc != '.') || (!unicode.IsDigit(pc) && pc != '.') {
							valid_number = true
						}
					}
				}
				// check whether the next character is a symbol
				if !valid_number {
					if k <= (width - 1) { // k < width ?
						ki := y*width + k
						if !unicode.IsDigit(arr[ki]) && arr[ki] != '.' {
							valid_number = true
						}
					}
				}
				// check diagonals
				if !valid_number {
					lx := x - 1
					rx := k
					tly := y - 1
					bly := y + 1
					// left
					if lx >= 0 {
						// top
						if tly >= 0 {
							idx := tly*width + lx
							if !unicode.IsDigit(arr[idx]) && arr[idx] != '.' {
								valid_number = true
							}
						}
						// bottom
						if !valid_number && (bly < height) {
							idx := bly*width + lx
							if !unicode.IsDigit(arr[idx]) && arr[idx] != '.' {
								valid_number = true
							}
						}
					}
					// right
					if !valid_number && rx < width {
						// top
						if tly >= 0 {
							idx := tly*width + rx
							if !unicode.IsDigit(arr[idx]) && arr[idx] != '.' {
								valid_number = true
							}
						}
						// bottom
						if !valid_number && (bly < height) {
							idx := bly*width + rx
							if !unicode.IsDigit(arr[idx]) && arr[idx] != '.' {
								valid_number = true
							}
						}
					}
				}
				if valid_number {
					//fmt.Println(string(arr[x:k]), "y: ", y, "x: ", x)
					number, err := strconv.Atoi(string(arr[(y*width + x) : (y*width + k)]))
					if err != nil {
						panic(err)
					}
					sum += number
				}
				x = k - 1
			}
		}
	}
	fmt.Println(sum)
*/
}