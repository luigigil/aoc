package day1

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	reader "github.com/luigigil/aoc2024/utils"
)

func Part1() {
	s := reader.GetScanner("day1/input.txt")
	left := []int{}
	right := []int{}

	for s.Scan() != false {
		text := s.Text()

		s0, s1 := splitAndConvert(text)

		left = append(left, s0)
		right = append(right, s1)
	}

	slices.Sort(left)
	slices.Sort(right)

	fmt.Println(left)
	fmt.Println(right)

	ans := 0
	for i := 0; i < len(left); i++ {
		diff := math.Abs(float64(left[i]) - float64(right[i]))
		fmt.Println(left[i], right[i], diff)
		ans = ans + int(diff)
	}

	fmt.Println(ans)
}

func Part2() {
	s := reader.GetScanner("day1/input.txt")
	left := []int{}
	right := make(map[int]int, 1)

	for s.Scan() != false {
		text := s.Text()

		s0, s1 := splitAndConvert(text)

		left = append(left, s0)

		v, ok := right[s1]
		if !ok {
			right[s1] = right[s1] + 1
		} else {
			right[s1] = v + 1
		}
	}

	ans := 0
	for i := 0; i < len(left); i++ {
		ans = ans + (left[i] * right[left[i]])
	}

	fmt.Println(ans)
}

func splitAndConvert(text string) (s0, s1 int) {
	split := strings.Split(text, "   ")

	fmt.Println(split)

	s0, _ = strconv.Atoi(split[0])
	s1, _ = strconv.Atoi(split[1])

	return s0, s1
}
