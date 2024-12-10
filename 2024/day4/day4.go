package day4

import (
	"fmt"

	reader "github.com/luigigil/aoc2024/utils"
)

func Part2() {
	s := reader.GetScanner("day4/input.txt")
	ans := 0

	input := []string{}
	for s.Scan() != false {
		input = append(input, s.Text())
	}

	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[0])-1; j++ {
			if input[i][j] != 'A' {
				continue
			}

			hasOneMAS := (input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S') || (input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M')

			if !hasOneMAS {
				continue
			}

			hasTwoMAS := (input[i+1][j-1] == 'M' && input[i-1][j+1] == 'S') || (input[i+1][j-1] == 'S' && input[i-1][j+1] == 'M')

			if !hasTwoMAS {
				continue
			}
			fmt.Printf("%d,%d\n", i, j)

			ans += 1
		}
	}

	fmt.Println(ans)
}

func Part1() {
	s := reader.GetScanner("day4/input.txt")
	ans := 0

	input := []string{}
	for s.Scan() != false {
		input = append(input, s.Text())
	}

	var process func(i, j int, target byte, direction string)
	process = func(i, j int, target byte, direction string) {
		if input[i][j] != target {
			return
		}

		if input[i][j] == 'S' {
			ans = ans + 1
			return
		}

		next := nextTarget(target)
		// up
		if i-1 >= 0 && direction == "up" {
			process(i-1, j, next, "up")
		}
		// left up
		if i-1 >= 0 && j-1 >= 0 && direction == "left up" {
			process(i-1, j-1, next, "left up")
		}
		// left
		if j-1 >= 0 && direction == "left" {
			process(i, j-1, next, "left")
		}
		// left down
		if j-1 >= 0 && i+1 < len(input) && direction == "left down" {
			process(i+1, j-1, next, "left down")
		}
		// down
		if i+1 < len(input) && direction == "down" {
			process(i+1, j, next, "down")
		}
		// right down
		if i+1 < len(input) && j+1 < len(input[0]) && direction == "right down" {
			process(i+1, j+1, next, "right down")
		}
		// right
		if j+1 < len(input[0]) && direction == "right" {
			process(i, j+1, next, "right")
		}
		// right up
		if i-1 >= 0 && j+1 < len(input[0]) && direction == "right up" {
			process(i-1, j+1, next, "right up")
		}
		return
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			process(i, j, 'X', "up")
			process(i, j, 'X', "left up")
			process(i, j, 'X', "left")
			process(i, j, 'X', "left down")
			process(i, j, 'X', "down")
			process(i, j, 'X', "right down")
			process(i, j, 'X', "right")
			process(i, j, 'X', "right up")
		}
	}

	fmt.Println(ans)
}

func nextTarget(target byte) byte {
	switch target {
	case 'X':
		return 'M'
	case 'M':
		return 'A'
	case 'A':
		return 'S'
	default:
		return ' '
	}
}
