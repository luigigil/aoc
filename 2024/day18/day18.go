package day18

import (
	"fmt"
	"strconv"
	"strings"

	reader "github.com/luigigil/aoc2024/utils"
)

var (
	// MAX_X = 6  // 70
	// MAX_Y = 6  // 70
	// BYTES = 12 // 1024

	MAX_X = 70   // 70
	MAX_Y = 70   // 70
	BYTES = 1024 // 1024
)

func Part2() {
	s := reader.GetScanner("day18/input.txt")

	isValidPoint := func(m, n int, grid [][]bool) bool {
		if m < 0 || m >= len(grid) || n < 0 || n >= len(grid[0]) || grid[m][n] {
			return false
		}
		return true
	}

	bfs := func(grid [][]bool) int {
		curr := [][2]int{{0, 0}}
		var next [][2]int

		for depth := 0; len(curr) > 0; curr, next, depth = next, curr[:0], depth+1 {
			for _, c := range curr {
				m, n := c[0], c[1]
				if grid[m][n] {
					continue
				}

				grid[m][n] = true

				if m == MAX_Y && n == MAX_X {
					return depth
				}

				if isValidPoint(m-1, n, grid) { // up
					next = append(next, [2]int{m - 1, n})
				}
				if isValidPoint(m+1, n, grid) { // down
					next = append(next, [2]int{m + 1, n})
				}
				if isValidPoint(m, n-1, grid) { // left
					next = append(next, [2]int{m, n - 1})
				}
				if isValidPoint(m, n+1, grid) { // right
					next = append(next, [2]int{m, n + 1})
				}
			}
		}
		return -1
	}

	grid := make([][]bool, MAX_Y+1)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]bool, MAX_X+1)
	}

	ans := [2]int{-1, -1}
	for s.Scan() != false {
		split := strings.Split(s.Text(), ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		grid[y][x] = true

		gridCopy := make([][]bool, MAX_Y+1)
		for i := 0; i < len(gridCopy); i++ {
			gridCopy[i] = make([]bool, MAX_X+1)
			for j := range grid[i] {
				gridCopy[i][j] = grid[i][j]
			}
		}

		result := bfs(gridCopy)
		if result == -1 {
			ans = [2]int{x, y}
			break
		}
	}

	fmt.Println(ans)
}
func Part1() {
	s := reader.GetScanner("day18/input.txt")

	grid := make([][]bool, MAX_Y+1)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]bool, MAX_X+1)
	}

	count := BYTES
	for s.Scan() != false {
		split := strings.Split(s.Text(), ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		grid[y][x] = true
		count--
		if count == 0 {
			break
		}
	}

	isValidPoint := func(m, n int) bool {
		if m < 0 || m >= len(grid) || n < 0 || n >= len(grid[0]) || grid[m][n] {
			return false
		}
		return true
	}
	bfs := func() int {
		curr := [][2]int{{0, 0}}
		var next [][2]int

		for depth := 0; len(curr) > 0; curr, next, depth = next, curr[:0], depth+1 {
			for _, c := range curr {
				m, n := c[0], c[1]
				if grid[m][n] {
					continue
				}

				grid[m][n] = true

				if m == MAX_Y && n == MAX_X {
					return depth
				}

				if isValidPoint(m-1, n) { // up
					next = append(next, [2]int{m - 1, n})
				}
				if isValidPoint(m+1, n) { // down
					next = append(next, [2]int{m + 1, n})
				}
				if isValidPoint(m, n-1) { // left
					next = append(next, [2]int{m, n - 1})
				}
				if isValidPoint(m, n+1) { // right
					next = append(next, [2]int{m, n + 1})
				}
			}
		}
		return -1
	}
	result := bfs()
	fmt.Println(result)
}
