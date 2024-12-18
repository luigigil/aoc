package day14

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	reader "github.com/luigigil/aoc2024/utils"
)

func getCoordinates(input string) [2]int {
	i := strings.Split(input, "=")
	i = strings.Split(i[1], ",")

	y, _ := strconv.Atoi(i[1])
	x, _ := strconv.Atoi(i[0])

	return [2]int{y, x}
}

func getQuadrant(pos [2]int, y, x int) int {
	halfX := math.Floor(float64(x) / 2.0)
	halfY := math.Floor(float64(y) / 2.0)

	if pos[0] < int(halfY) && pos[1] < int(halfX) {
		return 0
	}

	if pos[0] < int(halfY) && pos[1] > int(halfX) {
		return 1
	}

	if pos[0] > int(halfY) && pos[1] < int(halfX) {
		return 2
	}

	if pos[0] > int(halfY) && pos[1] > int(halfX) {
		return 3
	}

	return -1
}

func Part2() {
	s := reader.GetScanner("day14/input.txt")

	// real input
	yMax := 103
	xMax := 101

	var process func([2]int, [2]int, int) [2]int
	process = func(pos, vel [2]int, seconds int) [2]int {
		for i := 0; i < seconds; i++ {
			pos[0], pos[1] = pos[0]+vel[0], pos[1]+vel[1]
			if pos[0] < 0 {
				pos[0] += yMax
			}

			if pos[0] >= yMax {
				pos[0] -= yMax
			}

			if pos[1] < 0 {
				pos[1] += xMax
			}

			if pos[1] >= xMax {
				pos[1] -= xMax
			}
		}
		return pos
	}

	var area [103][101]byte

	for i := 0; i < yMax; i++ {
		for j := 0; j < xMax; j++ {
			area[i][j] = '.'
		}
	}

	ans := 1
	robots := [][2][2]int{}
	for s.Scan() != false {
		line := s.Text()

		split := strings.Split(line, " ")
		pos := getCoordinates(split[0])
		vel := getCoordinates(split[1])

		robots = append(robots, [2][2]int{pos, vel})
		area[pos[0]][pos[1]] = '#'
	}

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(robots); j++ {
			pos := robots[j][0]
			vel := robots[j][1]

			area[pos[0]][pos[1]] = '.'

			robots[j][0] = process(pos, vel, 1)

			newPos := robots[j][0]

			area[newPos[0]][newPos[1]] = '#'
		}

		fmt.Println(i)
		for y := 0; y < yMax; y++ {
			for x := 0; x < xMax; x++ {
				fmt.Printf("%s ", string(area[y][x]))
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}

	fmt.Println(ans)
}

func Part1() {
	s := reader.GetScanner("day14/input.txt")

	// real input
	yMax := 103
	xMax := 101

	// test input
	// yMax := 7
	// xMax := 11
	seconds := 100

	var process func([2]int, [2]int) [2]int
	process = func(pos, vel [2]int) [2]int {
		for i := 0; i < seconds; i++ {
			pos[0], pos[1] = pos[0]+vel[0], pos[1]+vel[1]
			if pos[0] < 0 {
				pos[0] += yMax
			}

			if pos[0] >= yMax {
				pos[0] -= yMax
			}

			if pos[1] < 0 {
				pos[1] += xMax
			}

			if pos[1] >= xMax {
				pos[1] -= xMax
			}
		}
		return pos
	}

	ans := 1
	quads := [4]int{0, 0, 0, 0}
	for s.Scan() != false {
		line := s.Text()

		split := strings.Split(line, " ")
		pos := getCoordinates(split[0])
		vel := getCoordinates(split[1])

		lastPos := process(pos, vel)
		quad := getQuadrant(lastPos, yMax, xMax)
		if quad != -1 {
			quads[quad] += 1
		}
	}

	for _, v := range quads {
		ans *= v
	}

	fmt.Println(ans)
}
