package day8

import (
	"fmt"

	reader "github.com/luigigil/aoc2024/utils"
)

func Part2() {
	s := reader.GetScanner("day8/input.txt")

	frequencies := make(map[byte][][]int)
	input := [][]byte{}
	row := 0
	for s.Scan() != false {
		line := s.Bytes()
		for column, i := range line {
			if i == '.' || i == '#' {
				continue
			}
			a, ok := frequencies[i]
			if !ok {
				frequencies[i] = [][]int{}
			}
			frequencies[i] = append(a, []int{row, column})
		}
		input = append(input, line)
		row += 1
	}

	fmt.Println()
	for _, i := range input {
		fmt.Println(string(i))
	}
	fmt.Println()

	aNodes := make(map[string]bool)
	for _, frequency := range frequencies {
		for i := 0; i < len(frequency); i++ {
			f1 := frequency[i]
			for j := 0; j < len(frequency); j++ {
				if i == j {
					continue
				}
				if !aNodes[fmt.Sprintf("%d,%d", f1[0], f1[1])] {
					aNodes[fmt.Sprintf("%d,%d", f1[0], f1[1])] = true
				}
				f2 := frequency[j]
				diffY, diffX := f1[0]-f2[0], f1[1]-f2[1]
				newY, newX := diffY+f1[0], diffX+f1[1]

				for newY >= 0 && newY < len(input) && newX >= 0 && newX < len(input[0]) {
					aNodes[fmt.Sprintf("%d,%d", newY, newX)] = true
					newY, newX = diffY+newY, diffX+newX
				}
			}
		}
	}
	for _, i := range input {
		fmt.Println(string(i))
	}
	fmt.Println()

	fmt.Println(len(aNodes))
}

func Part1() {
	s := reader.GetScanner("day8/input.txt")

	frequencies := make(map[byte][][]int)
	input := [][]byte{}
	row := 0
	for s.Scan() != false {
		line := s.Bytes()
		input = append(input, line)
		for column, i := range line {
			if i == '.' || i == '#' {
				continue
			}
			a, ok := frequencies[i]
			if !ok {
				frequencies[i] = [][]int{}
			}
			frequencies[i] = append(a, []int{row, column})
		}
		row += 1
	}

	count := 0
	for _, frequency := range frequencies {
		for i := 0; i < len(frequency); i++ {
			f1 := frequency[i]
			for j := 0; j < len(frequency); j++ {
				if i == j {
					continue
				}
				f2 := frequency[j]
				newY, newX := (f1[0]-f2[0])+f1[0], (f1[1]-f2[1])+f1[1]
				if newY >= 0 && newY < len(input) && newX >= 0 && newX < len(input[0]) {
					if input[newY][newX] == '#' {
						continue
					}
					input[newY][newX] = '#'
					count += 1
				}
			}
		}
	}

	fmt.Println(count)
}
