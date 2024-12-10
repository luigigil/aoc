package day2

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	reader "github.com/luigigil/aoc2024/utils"
)

func Part2() {
	s := reader.GetScanner("day2/input.txt")

	ans := 0
	for s.Scan() != false {
		text := s.Text()

		split := strings.Split(text, " ")

		// lets check if whole input is safe
		safe := safeChecker(split)

		// lets remove 1 level
		if !safe {
			fmt.Println("not safe let's try to remove one level")
			for i := 0; i < len(split); i++ {
				cloned := slices.Clone(split)
				cloned = slices.Delete(cloned, i, i+1)
				fmt.Println(i, cloned, split)
				safe = safeChecker(cloned)
				if safe {
					break
				}
			}
		}

		if safe {
			ans += 1
			fmt.Printf("safe, ans=%d\n", ans)
		}

		fmt.Println()
		fmt.Println()
	}
	fmt.Println(ans)
}

func safeChecker(split []string) bool {
	safe := true

	fmt.Println(split)

	first, _ := strconv.Atoi(split[0])
	second, _ := strconv.Atoi(split[1])
	asc := first < second
	fmt.Printf("asc=%t\n", asc)
	for i := 1; i < len(split); i++ {
		s0, _ := strconv.Atoi(split[i-1])
		s1, _ := strconv.Atoi(split[i])
		fmt.Printf("s0=%d; s1=%d\n", s0, s1)

		// check ascending or descending
		if asc && s0 > s1 {
			fmt.Println("ascends but s0 > s1")
			safe = false
			break
		}

		if !asc && s0 < s1 {
			fmt.Println("descends but s0 < s1")
			safe = false
			break
		}

		// check differ at least 1 or at most 3
		diff := float64(s0) - float64(s1)
		fmt.Printf("diff=%f\n", diff)
		diff = math.Abs(diff)
		fmt.Printf("abs(diff)=%f\n", diff)
		if diff > 3 || diff < 1 {
			fmt.Printf("diff greater than 3 or lower than 1\n")
			safe = false
			break
		}
	}

	return safe
}

func Part1() {
	s := reader.GetScanner("day2/input.txt")

	ans := 0
	for s.Scan() != false {
		text := s.Text()

		split := strings.Split(text, " ")

		safe := safeChecker(split)

		if safe {
			ans += 1
			fmt.Printf("safe, ans=%d\n", ans)
		}
		fmt.Println()
		fmt.Println()
	}

	fmt.Println(ans)
}
