package day9

import (
	"fmt"
	"slices"
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
	spacesSize := make(map[int]int)
	spacesPosition := []int{}
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
			spacesSize[len(memory)] = val
			spacesPosition = append(spacesPosition, len(memory))
			for j := 0; j < val; j++ {
				memory = append(memory, -1)
			}
		}
	}

	maxId := ids - 1
	filesMoved := make(map[int]bool)
	iSpace := 0
	for {
		if iSpace == len(spacesPosition) {
			break
		}
		pos := spacesPosition[iSpace]
		value := spacesSize[pos]

		iSpace += 1

		for id := maxId; id >= 0; id-- {
			if filesMoved[id] {
				continue
			}
			if filesSize[id] > value {
				continue
			}
			filesMoved[id] = true

			for i := 0; i < filesSize[id]; i++ {
				memory[pos+i] = id
				memory[filesPosition[id]+i] = -1
			}

			if filesSize[id] != value {
				spacesPosition = append(spacesPosition, pos+filesSize[id])
				spacesSize[pos+filesSize[id]] = value - filesSize[id]
			}

			spacesPosition = spacesPosition[1:]
			slices.Sort(spacesPosition)

			iSpace = 0
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
