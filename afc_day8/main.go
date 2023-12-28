package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type StringPair struct {
	Start string
	End   string
}

var steps int = 0

func computeSteps(instructions []string, starting_nodes []StringPair, input map[string]StringPair) {
	cache := make(map[string]string)
	// Initialize the cache with default values.
	for k := 0; k < len(starting_nodes); k++ {
		cache[starting_nodes[k].Start] = "None"
	}
	computeStepsAux(instructions, 0, starting_nodes, input, cache)
}

func testNode(instructions []string, node StringPair, input map[string]StringPair) {
	for k := 0; k < len(instructions); k++ {
		if instructions[k] == "L" {
			node.End = input[node.End].Start
		} else {
			node.End = input[node.End].End
		}
	}
	log.Printf("Node: %v, Start: %s, End: %s\n", node, node.Start, node.End)
}

func cacheNode(instructions []string, node StringPair, input map[string]StringPair, cache map[string]string) {
	for k := 0; k < len(instructions); k++ {
		if instructions[k] == "L" {
			node.End = input[node.End].Start
		} else {
			node.End = input[node.End].End
		}
	}
	cache[node.Start] = node.End
}

func dumpNodeCacheIntoFile(cache map[string]string) {
	file, err := os.Create("cache.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
		log.Println("Successfully wrote cache to the file.")
	}()

	for k, v := range cache {
		str := fmt.Sprintf("cache[%s]: %s\n", k, v)
		file.WriteString(str)
	}
}

func computeStepsAux(instructions []string, index int, nodes []StringPair, input map[string]StringPair, cache map[string]string) {
	testNode(instructions, StringPair{"GPG", "GPG"}, input)
	testNode(instructions, StringPair{"DRB", "DRB"}, input)

	// DFA, BLA, TGA, AAA, PQA, CQA
	testNode(instructions, StringPair{"DFA", "DFA"}, input)
	testNode(instructions, StringPair{"BLA", "BLA"}, input)
	testNode(instructions, StringPair{"TGA", "TGA"}, input)
	testNode(instructions, StringPair{"AAA", "AAA"}, input)
	testNode(instructions, StringPair{"PQA", "PQA"}, input)
	testNode(instructions, StringPair{"CQA", "CQA"}, input)

	rel_index := index % len(instructions)
	have_Z := 0
	if rel_index == 0 && index != 0 {
		cache_hits := 0
		old_nodes := nodes[:]
		for i := 0; i < len(nodes); i++ {
			// initialize old nodes.
			old_key := nodes[i].Start
			cache[old_key] = nodes[i].End

			// add new nodes
			new_key := nodes[i].End
			v, ok := cache[new_key]
			if ok && v != "None" {
				cache_hits++
			} else {
				cache[new_key] = "None"
			}

			// update nodes
			nodes[i].Start = nodes[i].End
		}
		if cache_hits == len(nodes) {
			// steps += len(instructions)
			index += len(instructions) - 1
			for k := 0; k < len(nodes); k++ {
				testNode(instructions, old_nodes[k], input)
				if nodes[k].End[2:] == "Z" {
					have_Z++
				}
				nodes[k].End = cache[nodes[k].Start]
			}
		} else {
			for k := 0; k < len(nodes); k++ {
				if nodes[k].End[2:] == "Z" {
					have_Z++
				}

				key := nodes[k].End
				if instructions[rel_index] == "L" {
					left := input[key].Start
					nodes[k].End = left
				} else { // R
					right := input[key].End
					nodes[k].End = right
				}
			}
			// steps++
		}
	} else if index == 0 {
		for i := 0; i < len(nodes); i++ {
			key := nodes[i].End
			if instructions[rel_index] == "L" {
				left := input[key].Start
				nodes[i].End = left
			} else { // R
				right := input[key].End
				nodes[i].End = right
			}
			if nodes[i].End[2:] == "Z" {
				have_Z++
			}
		}
		// steps++
	} else {
		for i := 0; i < len(nodes); i++ {
			if nodes[i].End[2:] == "Z" {
				have_Z++
			}

			key := nodes[i].End
			if instructions[rel_index] == "L" {
				left := input[key].Start
				nodes[i].End = left
			} else { // R
				right := input[key].End
				nodes[i].End = right
			}
		}
	}

	if have_Z == len(nodes) {
		return
	}

	// jump to the next instruction
	computeStepsAux(instructions, index+1, nodes, input, cache)
}

func printInput(m map[string]StringPair) {
	for k, v := range m {
		log.Printf("m[%s]: %v\n", k, v)
	}
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

	input := make(map[string]StringPair)
	instructions := []string{}
	starting_nodes := []StringPair{}
	cache := make(map[string]string)
	all_nodes := []StringPair{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if !strings.Contains(text, "=") {
			for _, v := range text {
				vs := string(v)
				if vs == "L" || vs == "R" {
					instructions = append(instructions, vs)
				}
			}
			continue
		}
		key := strings.Trim(strings.Split(text, "=")[0], " ")
		if key[2:] == "A" {
			starting_nodes = append(starting_nodes, StringPair{key, key})
		}
		rest := strings.Trim(strings.Split(text, "=")[1], " ")
		left := strings.Trim(strings.Split(rest, ",")[0], "( ")
		right := strings.Trim(strings.Split(rest, ",")[1], ") ")

		input[key] = StringPair{left, right}

		all_nodes = append(all_nodes, StringPair{key, key})
		all_nodes = append(all_nodes, StringPair{left, left})
		all_nodes = append(all_nodes, StringPair{right, right})
	}

	for k := 0; k < len(all_nodes); k++ {
		cacheNode(instructions, all_nodes[k], input, cache)
	}

	// only for debugging
	dumpNodeCacheIntoFile(cache)

	// init cache
	growing_cache := make(map[string]string)
	for k := 0; k < len(starting_nodes); k++ {
		growing_cache[starting_nodes[k].Start] = "None"
	}

	for k := 0; k < len(instructions); {
		// update all the nodes
		nodes_with_Z := 0
		for j := 0; j < len(starting_nodes); j++ {
			if instructions[k] == "L" {
				starting_nodes[j].End = input[starting_nodes[j].End].Start
			} else {
				starting_nodes[j].End = input[starting_nodes[j].End].End
			}

			if starting_nodes[j].End[2:] == "Z" {
				nodes_with_Z++
			}
		}

		if nodes_with_Z == len(starting_nodes) {
			break // found the solution
		}

		if k == len(instructions)-1 {
			cache_hits := 0
			for j := 0; j < len(starting_nodes); j++ {
				testNode(instructions, StringPair{starting_nodes[j].Start, starting_nodes[j].Start}, input)
				// cache the previous element
				v, ok := growing_cache[starting_nodes[j].Start]
				if ok && v == "None" { // doesn't exist
					growing_cache[starting_nodes[j].Start] = starting_nodes[j].End
				} else {
					log.Printf("node: %s has already been cached\n", starting_nodes[j].Start)
					if v != starting_nodes[j].End {
						// NOTE(alx): No idea why does it occur.
						// What if we reassign the old cache?
						starting_nodes[j].End = growing_cache[starting_nodes[j].Start]
						log.Println("Cache is not the same!")
					}
				}

				// TODO(alx): Make sure that v actually contains a value
				v, ok = growing_cache[starting_nodes[j].End]
				if ok && v != "None" {
					cache_hits++
				} else {
					growing_cache[starting_nodes[j].End] = "None"
				}

				// assign new node
				starting_nodes[j].Start = starting_nodes[j].End
			}

			if cache_hits == len(starting_nodes) {
				for j := 0; j < len(starting_nodes); j++ {
					// although Start and End are the same at this point.
					if starting_nodes[j].End != starting_nodes[j].Start {
						log.Println("Nodes are not the same.")
					}
					starting_nodes[j].End = growing_cache[starting_nodes[j].Start]
				}
			}
			k = 0
			continue
		}
		k++
	}

	// computeSteps(instructions, starting_nodes, input)
	// log.Println("steps: ", steps)

	// For example, following one instruction cycle we can get from 11A -> 11Z
	// we have to memoize that in a table.
	// memoization_table := make(map[string]string)
	// steps := 0
	// for true {
	// 	index := steps % instructions.Len()
	// 	if steps == instructions.Len() {
	// 		for _, v := range nodes {
	// 			_, ok := memoization_table[v.(StringPair).First]
	// 			if ok {

	// 			}
	// 			memoization_table[v.(StringPair).First] = v.(StringPair).Second
	// 		}
	// 	}

	// 	for i := 0; i < nodes.Len(); i++ {
	// 		if instructions[index] == "L" {
	// 			first := nodes[i].(StringPair).First
	// 			second := network[nodes[i].(StringPair).Second].First
	// 			nodes.Set(i, StringPair{first, second})
	// 		} else if instructions[index] == "R" {
	// 			first := nodes[i].(StringPair).First
	// 			second := network[nodes[i].(StringPair).Second].Second
	// 			nodes.Set(i, StringPair{first, second})
	// 		}
	// 	}

	// 	all_nodes_with_Z := true
	// 	for i := 0; i < nodes.Len(); i++ {
	// 		if nodes[i].(StringPair).Second[2:] != "Z" {
	// 			all_nodes_with_Z = false
	// 			break
	// 		}
	// 	}

	// 	steps++

	// 	if all_nodes_with_Z {
	// 		break
	// 	}
	// }
	// log.Println("steps: ", steps)

	log.Printf("instructions: %v\n", instructions)
	log.Printf("starting_nodes: %v\n", starting_nodes)
	printInput(input)
}
