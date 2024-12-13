package day9

import (
	"fmt"
	"strconv"

	reader "github.com/luigigil/aoc2024/utils"
)

func Part2() {
	s := reader.GetScanner("day9/input.txt")

	input := []byte{}
	for s.Scan() != false {
		input = s.Bytes()
	}

	ids := 0
	memory := []int{}
	filesSize := make(map[int]int)
	filesPosition := make(map[int]int)
	for i, v := range input {
		val, _ := strconv.Atoi(string(v))
		if i%2 == 0 {
			// file
			filesSize[ids] = val
			filesPosition[ids] = len(memory)
			for j := 0; j < val; j++ {
				memory = append(memory, ids)
			}
			ids += 1
		} else {
			// free space
			for j := 0; j < val; j++ {
				memory = append(memory, -1)
			}
		}
	}

	maxId := ids - 1

	for id := maxId; id >= 0; id-- {
		init := 0
		end := 0
		for init = 0; init < len(memory); init++ {
			if memory[init] != -1 {
				continue
			}

			for end = init; end < len(memory); end++ {
				if memory[end] != -1 {
					break
				}
			}

			freeSize := end - init

			if filesPosition[id] <= init {
				break
			}

			if filesSize[id] > freeSize {
				continue
			}

			for k := 0; k < filesSize[id]; k++ {
				memory[init+k] = id
				memory[filesPosition[id]+k] = -1
			}
			break
		}
	}

	checksum := 0
	for i, id := range memory {
		if id == -1 {
			continue
		}
		checksum += i * id
	}

	fmt.Println(checksum)
}

func Part1() {
	s := reader.GetScanner("day9/input.txt")

	input := []byte{}
	for s.Scan() != false {
		input = s.Bytes()
	}

	id := 0
	memory := []int{}
	for i, v := range input {
		val, _ := strconv.Atoi(string(v))
		if i%2 == 0 {
			// file
			for j := 0; j < val; j++ {
				memory = append(memory, id)
			}
			id += 1
		} else {
			// free space
			for j := 0; j < val; j++ {
				memory = append(memory, -1)
			}
		}

	}

	for j := len(memory) - 1; j > 0; j-- {
		for i := 0; i < len(memory); i++ {
			if j == i {
				break
			}
			if memory[i] != -1 {
				continue
			}
			if memory[j] == -1 {
				j -= 1
				i -= 1
				continue
			}
			memory[i] = memory[j]
			memory[j] = -1
			j -= 1
		}
	}

	checksum := 0
	for i := 0; i < len(memory); i++ {
		if memory[i] == -1 {
			break
		}
		checksum += i * memory[i]
	}

	fmt.Println(string(input))
	fmt.Println(memory)
	fmt.Println(checksum)
}
