package day12

import (
	"fmt"
	"slices"

	reader "github.com/luigigil/aoc2024/utils"
)

func getNeighbors2(i, j int) [][3]int {
	// third element is the direction
	return [][3]int{
		{i - 1, j, 0},
		{i + 1, j, 1},
		{i, j - 1, 2},
		{i, j + 1, 3},
	}
}

func countSides(sides [][3]int) int {
	processed := make(map[[3]int]bool)
	sideCount := 0

	// order sides taking preference to y then x
	slices.SortFunc(sides, func(a, b [3]int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})

	for _, s := range sides {
		sideNeighbors := getNeighbors2(s[0], s[1])
		found := false

		for _, c := range sideNeighbors {
			// set the neighbor direction to match curr direction
			c[2] = s[2]
			// check if we've seen this side with the same direction
			if _, f := processed[c]; f {
				found = true
			}
		}
		if !found {
			sideCount += 1
		}
		processed[s] = true
	}

	return sideCount
}

func Part2() {
	s := reader.GetScanner("day12/input.txt")

	input := [][]byte{}
	for s.Scan() != false {
		line := s.Text()
		input = append(input, []byte(line))
	}

	var dfs func([]int, [][]bool, *[][3]int) (int, int)
	dfs = func(pos []int, seen [][]bool, sides *[][3]int) (int, int) {
		key, y, x := input[pos[0]][pos[1]], pos[0], pos[1]

		seen[y][x] = true

		fences := 0
		area := 1

		neighbors := getNeighbors2(y, x)

		for _, n := range neighbors {
			ny, nx := n[0], n[1]

			if ny < 0 || ny >= len(input) || nx < 0 || nx >= len(input[ny]) || input[ny][nx] != key {
				if sides != nil {
					*sides = append(*sides, [3]int{ny, nx, n[2]})
				}
				fences += 1
			} else if !seen[ny][nx] {
				f, a := dfs([]int{ny, nx}, seen, sides)
				fences += f
				area += a
			}
		}

		return fences, area
	}

	seen := make([][]bool, len(input))
	for i := range seen {
		seen[i] = make([]bool, len(input[i]))
	}

	ans := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if !seen[i][j] {
				sides := make([][3]int, 0)
				_, area := dfs([]int{i, j}, seen, &sides)
				sideCount := countSides(sides)
				ans += sideCount * area
			}
		}
	}

	fmt.Println(ans)
}

func Part1() {
	s := reader.GetScanner("day12/input.txt")

	input := [][]byte{}
	for s.Scan() != false {
		line := s.Text()
		input = append(input, []byte(line))
	}

	var dfs func([]int, [][]bool) (int, int)
	dfs = func(pos []int, seen [][]bool) (int, int) {
		key, y, x := input[pos[0]][pos[1]], pos[0], pos[1]

		seen[y][x] = true

		fences := 0
		area := 1

		neighbors := getNeighbors2(y, x)

		for _, n := range neighbors {
			ny, nx := n[0], n[1]

			if ny < 0 || ny >= len(input) || nx < 0 || nx >= len(input[ny]) || input[ny][nx] != key {
				fences += 1
			} else if !seen[ny][nx] {
				f, a := dfs([]int{ny, nx}, seen)
				fences += f
				area += a
			}
		}

		return fences, area
	}

	seen := make([][]bool, len(input))
	for i := range seen {
		seen[i] = make([]bool, len(input[i]))
	}

	ans := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if !seen[i][j] {
				fences, area := dfs([]int{i, j}, seen)
				ans += fences * area
			}
		}
	}

	fmt.Println(ans)
}
