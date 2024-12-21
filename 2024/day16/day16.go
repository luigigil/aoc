package day16

import (
	"fmt"
	reader "github.com/luigigil/aoc2024/utils"
)

func Part1() {
	s := reader.GetScanner("day16/input.in")

	ans := 0
	input := make([][]byte, 0)
	moves := []byte{}
	gotMap := false
	for s.Scan() != false {
		line := s.Text()

		if len(line) == 0 {
			gotMap = true
			continue
		}

		if gotMap {
			for i := range line {
				moves = append(moves, line[i])
			}
		} else {
			input = append(input, []byte(line))
		}
	}

	pos := [2]int{0, 0}
	found := false
	for i := range input {
		for j := range input[i] {
			if input[i][j] == '@' {
				found = true
				pos = [2]int{i, j}
				break
			}
		}
		if found {
			break
		}
	}

	for _, m := range moves {
		nextY, nextX := pos[0], pos[1]

		switch m {
		case '^':
			nextY -= 1
		case '>':
			nextX += 1
		case '<':
			nextX -= 1
		case 'v':
			nextY += 1
		}

		if nextY < 0 || nextX < 0 || nextY >= len(input) || nextX >= len(input[0]) || input[nextY][nextX] == '#' {
			continue
		}

		if input[nextY][nextX] == '.' {
			input[nextY][nextX] = '@'
			input[pos[0]][pos[1]] = '.'
			pos[0], pos[1] = nextY, nextX

			continue
		}

		var dfs func(int, int) bool
		dfs = func(nextY, nextX int) bool {
			nY, nX := nextY, nextX

			switch m {
			case '^':
				nY -= 1
			case '>':
				nX += 1
			case '<':
				nX -= 1
			case 'v':
				nY += 1
			}

			if nY < 0 || nX < 0 || nY >= len(input) || nX >= len(input[0]) || input[nY][nX] == '#' {
				return false
			}

			if input[nY][nX] == '.' {
				input[nY][nX] = 'O'
				return true
			}

			return dfs(nY, nX)
		}

		canMove := dfs(nextY, nextX)

		if !canMove {
			continue
		}

		input[nextY][nextX] = '@'
		input[pos[0]][pos[1]] = '.'
		pos[0], pos[1] = nextY, nextX
	}

	for i := range input {
		fmt.Println(string(input[i]))
		for j := range input[i] {
			if input[i][j] != 'O' {
				continue
			}
			ans += 100*i + j
		}
	}

	fmt.Println(ans)
}
