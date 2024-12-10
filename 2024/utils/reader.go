package reader

import (
	"bufio"
	"os"
)

func GetScanner(path string) *bufio.Scanner {
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	return s
}
