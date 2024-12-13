package day11

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	reader "github.com/luigigil/aoc2024/utils"
)

func Part2() {
	s := reader.GetScanner("day11/input.txt")

	input := []int{}
	for s.Scan() != false {
		line := s.Text()
		split := strings.Split(line, " ")
		for _, spl := range split {
			val, _ := strconv.Atoi(string(spl))
			input = append(input, val)
		}
	}
	fmt.Println(input)

	cache := make(map[string]int)

	count := 0
	blinks := 75

	var process func(int, int) int
	process = func(num, blink int) int {
		if blink == blinks {
			cache[fmt.Sprintf("%d,%d", num, blink)] = 1
			return 1
		}

		if cached, ok := cache[fmt.Sprintf("%d,%d", num, blink)]; ok {
			return cached
		}

		blink += 1

		if num == 0 {
			result := process(1, blink)
			cache[fmt.Sprintf("%d,%d", 0, blink-1)] = result
			return result
		}

		str := fmt.Sprintf("%d", num)
		if len(str)%2 == 0 {
			d1, _ := strconv.Atoi(str[0 : len(str)/2])
			d2, _ := strconv.Atoi(str[len(str)/2:])

			result1 := process(d1, blink)
			result2 := process(d2, blink)

			cache[fmt.Sprintf("%d,%d", num, blink-1)] = result1 + result2

			return result1 + result2
		}

		mul := num * 2024
		result := process(mul, blink)
		cache[fmt.Sprintf("%d,%d", num, blink-1)] = result

		return result
	}

	for i := 0; i < len(input); i++ {
		count += process(input[i], 0)
	}

	fmt.Println(count)
}

func Part1() {
	s := reader.GetScanner("day11/input.txt")

	input := []int{}
	for s.Scan() != false {
		line := s.Text()
		split := strings.Split(line, " ")
		for _, spl := range split {
			val, _ := strconv.Atoi(string(spl))
			input = append(input, val)
		}
	}
	fmt.Println(input)

	blinks := 25
	for i := 0; i < blinks; i++ {
		for j := 0; j < len(input); j++ {
			if input[j] == 0 {
				input[j] = 1
				continue
			}
			str := fmt.Sprintf("%d", input[j])
			if len(str)%2 == 0 {
				d1, _ := strconv.Atoi(str[0 : len(str)/2])
				d2, _ := strconv.Atoi(str[len(str)/2:])
				input[j] = d1
				input = slices.Insert(input, j+1, d2)
				j += 1
				continue
			}
			input[j] *= 2024
		}
	}
	fmt.Println(len(input))
}
