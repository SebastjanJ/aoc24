package main

import (
	"container/heap"
	"fmt"
	"math"
)

type PriorityQueue []Deer

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Deer))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type Deer struct {
	pos   Point
	dir   rune
	score int
	path  []Point
}

func (d Deer) move() Deer {
	switch d.dir {
	case '<':
		d.pos.j--
	case '>':
		d.pos.j++
	case '^':
		d.pos.i--
	case 'v':
		d.pos.i++
	}
	return d
}

func (d Deer) turnRight() Deer {
	switch d.dir {
	case '<':
		d.dir = '^'
	case '>':
		d.dir = 'v'
	case '^':
		d.dir = '>'
	case 'v':
		d.dir = '<'
	}
	return d.move()
}

func (d Deer) turnLeft() Deer {
	switch d.dir {
	case '<':
		d.dir = 'v'
	case '>':
		d.dir = '^'
	case '^':
		d.dir = '<'
	case 'v':
		d.dir = '>'
	}
	return d.move()
}

func Day16() {
	// lines, _ := GetTest[string]()
	lines := getInputStrings(16)
	start := Point{len(lines) - 2, 1}
	end := Point{1, len(lines[0]) - 2}
	bestscore, bestpaths := dijkstra(lines, start, end)
	seats := getSeatsCount(bestpaths)
	fmt.Println(bestscore, seats)

}

func getSeatsCount(paths [][]Point) int {
	seats := make(map[Point]bool)
	for _, path := range paths {
		for _, p := range path {
			seats[p] = true
		}
	}
	return len(seats)
}

func dijkstra(lines []string, start, end Point) (int, [][]Point) {
	minscore := math.MaxInt32
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Deer{pos: start, dir: '>', score: 0, path: []Point{start}})
	fromstart := make(map[string]int)
	bestpaths := make([][]Point, 0)

	for pq.Len() > 0 {
		deer := heap.Pop(pq).(Deer)

		if deer.score > minscore {
			continue
		}
		key := fmt.Sprintf("%d,%d,%d", deer.pos.i, deer.pos.j, deer.dir)
		if prevScore, exists := fromstart[key]; exists && prevScore < deer.score {
			continue
		}
		fromstart[key] = deer.score

		if deer.pos == end {
			bestpaths = append(bestpaths, deer.path)
			minscore = deer.score
			continue
		}

		tmpdeer := deer.move()
		if lines[tmpdeer.pos.i][tmpdeer.pos.j] != '#' {
			tmpdeer.score++
			newPath := make([]Point, len(deer.path))
			copy(newPath, deer.path)
			tmpdeer.path = append(newPath, tmpdeer.pos)
			heap.Push(pq, tmpdeer)
		}
		tmpdeer = deer.turnRight()
		if lines[tmpdeer.pos.i][tmpdeer.pos.j] != '#' {
			tmpdeer.score += 1001
			newPath := make([]Point, len(deer.path))
			copy(newPath, deer.path)
			tmpdeer.path = append(newPath, tmpdeer.pos)
			heap.Push(pq, tmpdeer)
		}
		tmpdeer = deer.turnLeft()
		if lines[tmpdeer.pos.i][tmpdeer.pos.j] != '#' {
			tmpdeer.score += 1001
			newPath := make([]Point, len(deer.path))
			copy(newPath, deer.path)
			tmpdeer.path = append(newPath, tmpdeer.pos)
			heap.Push(pq, tmpdeer)
		}

	}
	return minscore, bestpaths
}
