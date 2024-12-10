package day5

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	reader "github.com/luigigil/aoc2024/utils"
)

func Part2() {
	s := reader.GetScanner("day5/input.txt")
	rules := make(map[int][]int, 0)
	ans := 0

	updates := [][]int{}
	// input := []string{}
	for s.Scan() != false {
		text := s.Text()
		if len(text) == 0 {
			continue
		}
		if strings.Contains(text, "|") {
			split := strings.Split(text, "|")
			n0, _ := strconv.Atoi(split[0])
			n1, _ := strconv.Atoi(split[1])
			if _, ok := rules[n0]; ok {
				rules[n0] = append(rules[n0], n1)
			} else {
				rules[n0] = []int{n1}
			}
		} else {
			split := strings.Split(text, ",")
			list := []int{}
			for _, n := range split {
				nn, _ := strconv.Atoi(n)
				list = append(list, nn)
			}
			updates = append(updates, list)
		}
	}

	incorrects := [][]int{}
	for _, u := range updates {
		isCorrect := true
		for i := 0; i < len(u)-1; i++ {
			rule := rules[u[i]]
			for j := 0; j < len(u[i+1:]); j++ {
				if !slices.Contains(rule, u[i+1:][j]) {
					isCorrect = false
					break
				}
			}
		}
		if !isCorrect {
			incorrects = append(incorrects, u)
		}
	}

	// fix them
	for _, incorrect := range incorrects {
		fixed := make([]int, len(incorrect))
		for i := 0; i < len(incorrect); i++ {
			rule := rules[incorrect[i]]
			counter := 0
			for j := 0; j < len(incorrect); j++ {
				if slices.Contains(rule, incorrect[j]) {
					counter += 1
				}
			}
			fixed[len(incorrect)-1-counter] = incorrect[i]
		}
		ans += fixed[int(math.Ceil(float64((len(fixed)-1)/2)))]
	}

	fmt.Println(rules)
	fmt.Println(incorrects)
	fmt.Println(ans)
}

func Part1() {
	s := reader.GetScanner("day5/input.txt")
	rules := make(map[int][]int, 0)
	ans := 0

	updates := [][]int{}
	// input := []string{}
	for s.Scan() != false {
		text := s.Text()
		if len(text) == 0 {
			continue
		}
		if strings.Contains(text, "|") {
			split := strings.Split(text, "|")
			n0, _ := strconv.Atoi(split[0])
			n1, _ := strconv.Atoi(split[1])
			if _, ok := rules[n0]; ok {
				rules[n0] = append(rules[n0], n1)
			} else {
				rules[n0] = []int{n1}
			}
		} else {
			split := strings.Split(text, ",")
			list := []int{}
			for _, n := range split {
				nn, _ := strconv.Atoi(n)
				list = append(list, nn)
			}
			updates = append(updates, list)
		}
	}

	for _, u := range updates {
		isCorrect := true
		for i := 0; i < len(u)-1; i++ {
			rule := rules[u[i]]
			for j := 0; j < len(u[i+1:]); j++ {
				if !slices.Contains(rule, u[i+1:][j]) {
					isCorrect = false
					break
				}
			}
		}
		if isCorrect {
			ans += u[int(math.Ceil(float64((len(u)-1)/2)))]
		}
	}

	fmt.Println(rules)
	fmt.Println(updates)
	fmt.Println(ans)
}
