package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

func first() int {
	inputBytes, err := os.ReadFile("2021/15/input.txt")
	if err != nil {
		panic(err)
	}

	nodes := make([][]*node, 0)

	lines := strings.Split(string(inputBytes), "\n")
	for i := 0; i < len(lines); i++ {
		nodes = append(nodes, make([]*node, 0))
		for j := 0; j < len(lines[i]); j++ {
			distance, _ := strconv.Atoi(string(lines[i][j]))
			nodes[i] = append(nodes[i], &node{distance: distance})
		}
	}

	graph := make(map[*node][]*node)
	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes[i]); j++ {
			n := nodes[i][j]
			if i-1 >= 0 {
				graph[n] = append(graph[n], nodes[i-1][j])
			}
			if j-1 >= 0 {
				graph[n] = append(graph[n], nodes[i][j-1])
			}
			if i+1 <= len(nodes)-1 {
				graph[n] = append(graph[n], nodes[i+1][j])
			}
			if j+1 <= len(nodes[i])-1 {
				graph[n] = append(graph[n], nodes[i][j+1])
			}
		}
	}

	start := nodes[0][0]
	end := nodes[len(nodes)-1][len(nodes[0])-1]

	return djikstra(graph, start, end)
}

func second() int {
	inputBytes, err := os.ReadFile("2021/15/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputBytes), "\n")
	l := len(lines)
	ll := len(lines[0])
	const scale = 5
	nodes := make([][]*node, scale*l)
	for i := 0; i < scale*l; i++ {
		nodes[i] = make([]*node, scale*ll)
	}
	for x := 0; x < scale; x++ {
		for y := 0; y < scale; y++ {
			for i := 0; i < l; i++ {
				for j := 0; j < ll; j++ {
					var distance int
					if x == 0 && y == 0 {
						distance, _ = strconv.Atoi(string(lines[i][j]))
					} else {
						if x > 0 {
							distance = nodes[i+(x-1)*l][j+y*ll].distance + 1
						} else if y > 0 {
							distance = nodes[i+x*l][j+(y-1)*ll].distance + 1
						}
						if distance > 9 {
							distance = 1
						}
					}
					nodes[i+x*l][j+y*ll] = &node{distance: distance}
				}
			}
		}
	}

	graph := make(map[*node][]*node)
	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes[i]); j++ {
			n := nodes[i][j]
			if i-1 >= 0 {
				graph[n] = append(graph[n], nodes[i-1][j])
			}
			if j-1 >= 0 {
				graph[n] = append(graph[n], nodes[i][j-1])
			}
			if i+1 <= len(nodes)-1 {
				graph[n] = append(graph[n], nodes[i+1][j])
			}
			if j+1 <= len(nodes[i])-1 {
				graph[n] = append(graph[n], nodes[i][j+1])
			}
		}
	}

	start := nodes[0][0]
	end := nodes[len(nodes)-1][len(nodes[0])-1]

	return djikstra(graph, start, end)
}

func djikstra(graph map[*node][]*node, start, end *node) int {
	distances := make(map[*node]int)
	distances[start] = 0

	unvisited := make(priorityQueue, len(graph))
	i := 0
	for n := range graph {
		if n != start {
			distances[n] = math.MaxInt
		}
		n.priority = distances[n]
		n.index = i
		unvisited[i] = n
		i++
	}
	heap.Init(&unvisited)

	for len(unvisited) > 0 {
		current := heap.Pop(&unvisited).(*node)

		if current == end {
			break
		}

		for _, neighbor := range graph[current] {
			distance := distances[current] + neighbor.distance
			if distance < distances[neighbor] {
				distances[neighbor] = distance
				unvisited.update(neighbor, neighbor.distance, distances[neighbor])
			}
		}
	}

	return distances[end]
}

type node struct {
	distance, priority, index int
}

type priorityQueue []*node

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i].priority < p[j].priority
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *priorityQueue) Push(x interface{}) {
	l := len(*p)
	n := x.(*node)
	n.index = l
	*p = append(*p, n)
}

func (p *priorityQueue) Pop() interface{} {
	old := *p
	l := len(old)
	n := old[l-1]
	old[l-1] = nil // avoid memory leak
	n.index = -1   // for safety
	*p = old[0 : l-1]
	return n
}

func (p *priorityQueue) update(n *node, distance int, priority int) {
	n.distance = distance
	n.priority = priority
	heap.Fix(p, n.index)
}
