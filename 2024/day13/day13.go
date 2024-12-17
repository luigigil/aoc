package day13

import (
	"fmt"
	"strconv"
	"strings"

	reader "github.com/luigigil/aoc2024/utils"
)

func getButtonBehavior(button string) [2]int {
	x, y := 0, 0

	button = strings.ReplaceAll(button, ",", "")
	split := strings.Split(button, " ")

	moveX := strings.Split(split[2], "+")
	x, _ = strconv.Atoi(moveX[1])

	moveY := strings.Split(split[3], "+")
	y, _ = strconv.Atoi(moveY[1])

	// y, x
	return [2]int{y, x}
}

func getPrizeCoordinates(prize string) [2]int {
	x, y := 0, 0
	prize = strings.ReplaceAll(prize, ",", "")
	split := strings.Split(prize, " ")

	moveX := strings.Split(split[1], "=")
	x, _ = strconv.Atoi(moveX[1])

	moveY := strings.Split(split[2], "=")
	y, _ = strconv.Atoi(moveY[1])

	return [2]int{y, x}
}

func cramer(a, b, p [2]int) [2]int {
	ax, ay := a[1], a[0]
	bx, by := b[1], b[0]
	px, py := p[1], p[0]

	divisor := (ax * by) - (bx * ay)

	A := (px*by - bx*py) / divisor
	B := (ax*py - px*ay) / divisor

	return [2]int{A, B}
}

func Part2() {
	s := reader.GetScanner("day13/input.txt")

	ans := 0
	var process func([]string)
	process = func(input []string) {
		buttonA := getButtonBehavior(input[0])
		buttonB := getButtonBehavior(input[1])
		prize := getPrizeCoordinates(input[2])
		prize = [2]int{prize[0] + 10000000000000, prize[1] + 10000000000000}

		presses := cramer(buttonA, buttonB, prize)

		if buttonA[0]*presses[0]+buttonB[0]*presses[1] == prize[0] &&
			buttonA[1]*presses[0]+buttonB[1]*presses[1] == prize[1] {
			ans += presses[0]*3 + presses[1]
		}

	}

	for {
		hasMore := s.Scan()
		if !hasMore {
			break
		}

		input := []string{}

		line := s.Text()
		if len(line) == 0 {
			continue
		}

		input = append(input, line)

		s.Scan()
		line = s.Text()
		input = append(input, line)

		s.Scan()
		line = s.Text()
		input = append(input, line)

		process(input)
	}

	fmt.Println(ans)
}

func Part1() {
	s := reader.GetScanner("day13/input.txt")

	ans := 0
	var process func([]string)
	process = func(input []string) {
		buttonA := getButtonBehavior(input[0])
		buttonB := getButtonBehavior(input[1])
		prize := getPrizeCoordinates(input[2])

		presses := cramer(buttonA, buttonB, prize)

		if buttonA[0]*presses[0]+buttonB[0]*presses[1] == prize[0] &&
			buttonA[1]*presses[0]+buttonB[1]*presses[1] == prize[1] {
			ans += presses[0]*3 + presses[1]
		}

	}

	for {
		hasMore := s.Scan()
		if !hasMore {
			break
		}

		input := []string{}

		line := s.Text()
		if len(line) == 0 {
			continue
		}

		input = append(input, line)

		s.Scan()
		line = s.Text()
		input = append(input, line)

		s.Scan()
		line = s.Text()
		input = append(input, line)

		process(input)
	}

	fmt.Println(ans)
}
