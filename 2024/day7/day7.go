package day7

import (
	"fmt"
	"strconv"
	"strings"

	reader "github.com/luigigil/aoc2024/utils"
)

func Part1() {
	s := reader.GetScanner("day7/input.txt")
	ans := 0

	for s.Scan() != false {
		line := string(s.Bytes())

		split := strings.Split(line, " ")

		var result int
		values := []int{}
		for idx, v := range split {
			if idx == 0 {
				result, _ = strconv.Atoi(strings.Replace(v, ":", "", 1))
				continue
			}
			value, _ := strconv.Atoi(v)
			values = append(values, value)
		}

		var process func(int, int) bool
		process = func(value, target int) bool {
			sum := value + values[target]
			mul := value * values[target]

			if target == len(values)-1 {
				return sum == result || mul == result
			}

			success := process(sum, target+1)
			if success {
				return success
			}
			success = process(mul, target+1)
			if success {
				return success
			}
			return false
		}

		found := process(values[0], 1)
		if found {
			ans += result
		}
	}

	fmt.Println(ans)
}

func Part2() {
	s := reader.GetScanner("day7/input.txt")
	ans := 0

	for s.Scan() != false {
		line := string(s.Bytes())

		split := strings.Split(line, " ")

		var result int
		values := []int{}
		for idx, v := range split {
			if idx == 0 {
				result, _ = strconv.Atoi(strings.Replace(v, ":", "", 1))
				continue
			}
			value, _ := strconv.Atoi(v)
			values = append(values, value)
		}

		var process func(int, int) bool
		process = func(value, target int) bool {
			sum := value + values[target]
			mul := value * values[target]
			concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", value, values[target]))

			if target == len(values)-1 {
				return sum == result || mul == result || concat == result
			}

			success := process(sum, target+1)
			if success {
				return success
			}
			success = process(mul, target+1)
			if success {
				return success
			}
			success = process(concat, target+1)
			if success {
				return success
			}
			return false
		}

		found := process(values[0], 1)
		if found {
			ans += result
		}
	}

	fmt.Println(ans)
}
