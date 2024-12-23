package day16

import (
	"container/heap"
	"fmt"
	"math"

	reader "github.com/luigigil/aoc2024/utils"
)

var directions = []string{"N", "E", "S", "W"}
var deltas = map[string][2]int{
	"N": {-1, 0},
	"S": {1, 0},
	"E": {0, 1},
	"W": {0, -1},
}

type State struct {
	x, y      int
	facing    string
	cost      int
	visited   map[[2]int]bool
	heapIndex int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].heapIndex = i
	pq[j].heapIndex = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*State)
	item.heapIndex = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func Part2() {
	s := reader.GetScanner("day16/example1.txt")

	input := [][]byte{}
	for s.Scan() != false {
		line := s.Text()
		input = append(input, []byte(line))
	}

	start := [2]int{0, 0}
	end := [2]int{0, 0}
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'S' {
				start[0], start[1] = i, j
			}

			if input[i][j] == 'E' {
				end[0], end[1] = i, j
			}
		}
	}

	startState := &State{
		x:       start[1],
		y:       start[0],
		facing:  "E",
		cost:    0,
		visited: map[[2]int]bool{{start[0], start[1]}: true},
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, startState)

	// y, x, facing
	visited := make(map[[3]interface{}]int)

	var dijkstra func() (int, map[[2]int]bool)
	dijkstra = func() (int, map[[2]int]bool) {
		foundMinCost := false
		minCost := math.MaxInt
		bestPathTiles := make(map[[2]int]bool)
		for pq.Len() > 0 {
			current := heap.Pop(pq).(*State)
			if current.x == end[1] && current.y == end[0] {
				if foundMinCost && current.cost > minCost {
					continue
				}
				minCost = current.cost
				for tile := range current.visited {
					bestPathTiles[tile] = true
				}
				foundMinCost = true
				continue
			}

			stateKey := [3]interface{}{current.y, current.x, current.facing}
			if visitedCost, ok := visited[stateKey]; ok && visitedCost <= current.cost {
				continue
			}
			visited[stateKey] = current.cost

			// move forward
			delta := deltas[current.facing]
			newY, newX := current.y+delta[0], current.x+delta[1]
			if input[newY][newX] != '#' {
				newVisited := copyVisited(current.visited)
				newVisited[[2]int{newY, newX}] = true
				heap.Push(pq, &State{y: newY, x: newX, facing: current.facing, cost: current.cost + 1, visited: newVisited})
			}

			// rotate clockwise
			newFacingCW := directions[(indexOf(current.facing)+1)%4]
			heap.Push(pq, &State{x: current.x, y: current.y, facing: newFacingCW, cost: current.cost + 1000, visited: copyVisited(current.visited)})

			// rotate counterclockwise
			newFacingCCW := directions[(indexOf(current.facing)+3)%4] // +3 is equivalent to -1 mod 4
			heap.Push(pq, &State{x: current.x, y: current.y, facing: newFacingCCW, cost: current.cost + 1000, visited: copyVisited(current.visited)})
		}
		return minCost, bestPathTiles
	}

	cost, bestPathTiles := dijkstra()

	fmt.Println(cost)
	fmt.Println(len(bestPathTiles))

	markedGrid := markBestPath(input, bestPathTiles)

	for _, row := range markedGrid {
		fmt.Println(string(row))
	}
}

func Part1() {
	s := reader.GetScanner("day16/example1.txt")

	input := [][]byte{}
	for s.Scan() != false {
		line := s.Text()
		input = append(input, []byte(line))
	}

	start := [2]int{0, 0}
	end := [2]int{0, 0}
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'S' {
				start[0], start[1] = i, j
			}

			if input[i][j] == 'E' {
				end[0], end[1] = i, j
			}
		}
	}

	startState := &State{
		x:      start[1],
		y:      start[0],
		facing: "E",
		cost:   0,
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, startState)

	// y, x, facing
	visited := make(map[[3]interface{}]int)

	var dijkstra func() int
	dijkstra = func() int {
		for pq.Len() > 0 {
			current := heap.Pop(pq).(*State)
			if current.x == end[1] && current.y == end[0] {
				return current.cost
			}

			stateKey := [3]interface{}{current.y, current.x, current.facing}
			if visitedCost, ok := visited[stateKey]; ok && visitedCost <= current.cost {
				continue
			}
			visited[stateKey] = current.cost

			// move forward
			delta := deltas[current.facing]
			newY, newX := current.y+delta[0], current.x+delta[1]
			if input[newY][newX] != '#' {
				heap.Push(pq, &State{y: newY, x: newX, facing: current.facing, cost: current.cost + 1})
			}

			// rotate clockwise
			newFacingCW := directions[(indexOf(current.facing)+1)%4]
			heap.Push(pq, &State{x: current.x, y: current.y, facing: newFacingCW, cost: current.cost + 1000})

			// rotate counterclockwise
			newFacingCCW := directions[(indexOf(current.facing)+3)%4] // +3 is equivalent to -1 mod 4
			heap.Push(pq, &State{x: current.x, y: current.y, facing: newFacingCCW, cost: current.cost + 1000})
		}
		return -1
	}

	result := dijkstra()
	fmt.Println(result)
}

func indexOf(direction string) int {
	for i, d := range directions {
		if d == direction {
			return i
		}
	}
	return -1
}

func copyVisited(visited map[[2]int]bool) map[[2]int]bool {
	newVisited := make(map[[2]int]bool)
	for k, v := range visited {
		newVisited[k] = v
	}
	return newVisited
}

func markBestPath(grid [][]byte, bestPathTiles map[[2]int]bool) [][]byte {
	markedGrid := make([][]byte, len(grid))
	for y := range grid {
		markedGrid[y] = make([]byte, len(grid[y]))
		copy(markedGrid[y], grid[y])
	}

	for tile := range bestPathTiles {
		x, y := tile[1], tile[0]
		if markedGrid[y][x] == '.' {
			markedGrid[y][x] = 'O'
		}
	}
	return markedGrid
}
