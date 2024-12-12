package day10

import (
	"fmt"
	"strconv"

	reader "github.com/luigigil/aoc2024/utils"
)

func Part2() {
	s := reader.GetScanner("day10/input.txt")

	input := [][]int{}
	for s.Scan() != false {
		line := []int{}
		for _, l := range s.Text() {
			val, err := strconv.Atoi(string(l))
			if err != nil {
				line = append(line, -1)
				continue
			}
			line = append(line, val)
		}
		input = append(input, line)
	}

	// find starts
	starts := [][]int{}
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == 0 {
				starts = append(starts, []int{i, j})
			}
		}
	}

	var dfs func([]int) int
	dfs = func(curr []int) int {
		y, x := curr[0], curr[1]

		if input[y][x] == 9 {
			return 1
		}

		acc := 0
		// dfs up
		if y-1 >= 0 && input[y-1][x] == input[y][x]+1 {
			acc += dfs([]int{y - 1, x})
		}
		// dfs down
		if y+1 < len(input) && input[y+1][x] == input[y][x]+1 {
			acc += dfs([]int{y + 1, x})
		}
		// dfs left
		if x-1 >= 0 && input[y][x-1] == input[y][x]+1 {
			acc += dfs([]int{y, x - 1})
		}
		// dfs right
		if x+1 < len(input[0]) && input[y][x+1] == input[y][x]+1 {
			acc += dfs([]int{y, x + 1})
		}

		return acc
	}

	ans := 0

	for i := 0; i < len(starts); i++ {
		ans += dfs(starts[i])
	}

	fmt.Println(ans)
}

func Part1() {
	s := reader.GetScanner("day10/input.txt")

	input := [][]int{}
	for s.Scan() != false {
		line := []int{}
		for _, l := range s.Text() {
			val, err := strconv.Atoi(string(l))
			if err != nil {
				line = append(line, -1)
				continue
			}
			line = append(line, val)
		}
		input = append(input, line)
	}

	// find starts
	starts := [][]int{}
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == 0 {
				starts = append(starts, []int{i, j})
			}
		}
	}

	var dfs func([]int, map[string]bool) int
	dfs = func(curr []int, seen map[string]bool) int {
		y, x := curr[0], curr[1]
		if seen[fmt.Sprintf("%d,%d", y, x)] {
			return 0
		}

		seen[fmt.Sprintf("%d,%d", y, x)] = true

		if input[y][x] == 9 {
			return 1
		}

		acc := 0
		// dfs up
		if y-1 >= 0 && input[y-1][x] == input[y][x]+1 {
			acc += dfs([]int{y - 1, x}, seen)
		}
		// dfs down
		if y+1 < len(input) && input[y+1][x] == input[y][x]+1 {
			acc += dfs([]int{y + 1, x}, seen)
		}
		// dfs left
		if x-1 >= 0 && input[y][x-1] == input[y][x]+1 {
			acc += dfs([]int{y, x - 1}, seen)
		}
		// dfs right
		if x+1 < len(input[0]) && input[y][x+1] == input[y][x]+1 {
			acc += dfs([]int{y, x + 1}, seen)
		}

		return acc
	}

	ans := 0

	for i := 0; i < len(starts); i++ {
		seen := make(map[string]bool)
		ans += dfs(starts[i], seen)
	}

	fmt.Println(ans)
}
