package day15

import (
	"fmt"
	reader "github.com/luigigil/aoc2024/utils"
)

func Part2() {
	s := reader.GetScanner("day15/input.txt")

	ans := 0
	input := make([][]byte, 0)
	moves := []byte{}
	gotMap := false
	pos := [2]int{0, 0}
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
			l := []byte{}
			for i := range line {
				if line[i] == '@' {
					l = append(l, line[i], '.')
					continue
				}
				if line[i] == 'O' {
					l = append(l, '[', ']')
				} else {
					l = append(l, line[i], line[i])
				}
			}
			input = append(input, l)
		}
	}

	found := false
	for i := range input {
		for j := range input[i] {
			if input[i][j] == '@' {
				found = true
				pos[0], pos[1] = i, j
				break
			}
		}
		if found {
			break
		}
	}

	var moveHorizontal func(int)
	moveHorizontal = func(deltaX int) {
		nextX := pos[1]

		// search for the next first free space
		for {
			nextX += deltaX

			nextSpot := input[pos[0]][nextX]

			if nextSpot == '#' {
				return
			}

			if nextSpot == '.' {
				break
			}
		}

		for {
			prevX := nextX - deltaX

			input[pos[0]][nextX] = input[pos[0]][prevX]

			nextX = prevX

			if nextX == pos[1] {
				break
			}
		}
		input[pos[0]][pos[1]] = '.'
		pos[1] += deltaX
	}

	var moveBoxVertical func(boxY, boxX, deltaY int)
	moveBoxVertical = func(boxY, boxX, deltaY int) {
		if input[boxY][boxX] == ']' {
			boxX -= 1
		}

		nextY := boxY + deltaY

		nextSpotLeft := input[nextY][boxX]
		nextSpotRight := input[nextY][boxX+1]

		if nextSpotLeft == '[' {
			moveBoxVertical(nextY, boxX, deltaY)
		}

		if nextSpotLeft == ']' {
			moveBoxVertical(nextY, boxX, deltaY)
		}

		if nextSpotRight == '[' {
			moveBoxVertical(nextY, boxX+1, deltaY)
		}

		input[boxY][boxX] = '.'
		input[boxY][boxX+1] = '.'

		input[nextY][boxX] = '['
		input[nextY][boxX+1] = ']'
	}

	var canMoveBoxVertical func(int, int, int) bool
	canMoveBoxVertical = func(boxY, boxX, deltaY int) bool {
		if input[boxY][boxX] == ']' {
			boxX -= 1
		}

		nextX := boxY + deltaY

		nextSpotLeft := input[nextX][boxX]
		nextSpotRight := input[nextX][boxX+1]

		if nextSpotLeft == '#' {
			return false
		}

		if nextSpotRight == '#' {
			return false
		}

		if nextSpotLeft == '[' {
			if !canMoveBoxVertical(nextX, boxX, deltaY) {
				return false
			}
		}

		if nextSpotLeft == ']' {
			if !canMoveBoxVertical(nextX, boxX, deltaY) {
				return false
			}
		}

		if nextSpotRight == '[' {
			if !canMoveBoxVertical(nextX, boxX+1, deltaY) {
				return false
			}
		}

		return true
	}

	var moveVertical func(int)
	moveVertical = func(deltaY int) {
		nextY := pos[0] + deltaY

		nextSpot := input[nextY][pos[1]]

		if nextSpot == '#' {
			return
		}

		if nextSpot == '.' {
			input[pos[0]][pos[1]] = '.'
			pos[0] = nextY
			input[pos[0]][pos[1]] = '@'
			return
		}

		if !canMoveBoxVertical(nextY, pos[1], deltaY) {
			return
		}

		moveBoxVertical(nextY, pos[1], deltaY)

		input[pos[0]][pos[1]] = '.'
		pos[0] = nextY
		input[pos[0]][pos[1]] = '@'
	}

	// for i := range input {
	// 	fmt.Println(string(input[i]))
	// }

	for _, m := range moves {
		switch m {
		case '^':
			moveVertical(-1)
		case 'v':
			moveVertical(1)
		case '>':
			moveHorizontal(1)
		case '<':
			moveHorizontal(-1)
		}
		// for i := range input {
		// 	fmt.Println(string(input[i]))
		// }
	}

	for i := range input {
		for j := range input[i] {
			if input[i][j] != '[' {
				continue
			}
			ans += 100*i + j
		}
	}

	fmt.Println(ans)
}

func Part1() {
	s := reader.GetScanner("day15/input.txt")

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
