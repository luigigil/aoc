package day3

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	reader "github.com/luigigil/aoc2024/utils"
)

func Part2() {
	s := reader.GetScanner("day3/input.txt")
	allowed := []byte{
		'1',
		'2',
		'3',
		'4',
		'5',
		'6',
		'7',
		'8',
		'9',
		'0',
		',',
	}

	var ans int64
	ans = 0
	enabled := true
	for s.Scan() != false {
		text := s.Text()

		for i := 0; i < len(text); i++ {
			if text[i] == 'm' {
				if text[i:i+4] != "mul(" {
					continue
				}

				i = i + 4
				start := i
				for j := i; j < len(text); j, i = j+1, i+1 {
					if slices.Contains(allowed, text[j]) {
						continue
					}

					if text[j] != ')' {
						break
					}

					if !enabled {
						break
					}

					nums := text[start:j]
					split := strings.Split(nums, ",")
					n0, _ := strconv.ParseInt(split[0], 10, 0)
					n1, _ := strconv.ParseInt(split[1], 10, 0)
					ans = ans + (n0 * n1)
					break
				}
			}
			if text[i] == 'd' {
				fmt.Println(text[i : i+6])
				if text[i:i+4] == "do()" {
					enabled = true
				}
				if text[i:i+7] == "don't()" {
					enabled = false
				}
			}
		}
	}

	fmt.Println(ans)
}

func Part1() {
	s := reader.GetScanner("day3/input.txt")
	allowed := []byte{
		'1',
		'2',
		'3',
		'4',
		'5',
		'6',
		'7',
		'8',
		'9',
		'0',
		',',
	}

	var ans int64
	ans = 0
	enabled := true
	for s.Scan() != false {
		text := s.Text()

		for i := 0; i < len(text); i++ {
			if text[i] == 'm' {
				if i+3 >= len(text) {
					continue
				}

				if text[i+1] != 'u' || text[i+2] != 'l' || text[i+3] != '(' {
					continue
				}

				i = i + 4
				start := i
				for j := i; j < len(text); j, i = j+1, i+1 {
					if slices.Contains(allowed, text[j]) {
						continue
					}

					if text[j] != ')' {
						break
					}

					if !enabled {
						break
					}

					nums := text[start:j]
					split := strings.Split(nums, ",")
					n0, _ := strconv.ParseInt(split[0], 10, 0)
					n1, _ := strconv.ParseInt(split[1], 10, 0)
					ans = ans + (n0 * n1)
					break
				}
			}
		}
	}

	fmt.Println(ans)
}
