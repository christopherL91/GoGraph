package main

import (
	"container/heap"
	"flag"
	"fmt"
	"github.com/christopherL91/Graph"
	"math"
	"runtime"
	"sort"
)

var (
	filename      string
	numofvertices int
)

func init() {
	flag.StringVar(&filename, "file", "hash.txt", "Which graph to use with the program")
	flag.IntVar(&numofvertices, "num", 10, "Number of vertices in graph")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	g := Graph.NewHash(numofvertices)
	g.Parse(filename)
	dijkstra(g, 1, 2)
}

func dijkstra(G *Graph.Hash, from, to int) {
	distance := make([]int, G.NumVertices())
	pq := &PriorityQueue{}
	for i, _ := range distance {
		if i == from {
			distance[i] = 0
		} else {
			distance[i] = math.MaxInt32
		}
		item := &Item{
			value:    i,
			priority: distance[i],
		}
		heap.Push(pq, item)
	}

	stack := make([]int, G.NumVertices())
	for i, _ := range stack {
		stack[i] = -1
	}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		v := current.value
		G.DoNeighbors(v, func(w, cost int) {
			dist := distance[v] + cost
			if dist < distance[w] {
				distance[w] = dist
				stack[w] = v
				item := pq.find(w)
				pq.update(item, dist)
			}
		})
	}
	printPath(G, from, to, stack)
}

func printPath(g *Graph.Hash, from, to int, stack []int) {
	if from > g.NumVertices() || to > g.NumVertices() {
		return
	}

	var path []int
	totalcost := 0
	found := false
	index := to
	path = append(path, index)
	for stack[index] != -1 {
		prev := stack[index]
		cost, _ := g.Cost(prev, index)
		totalcost += cost
		index = prev
		path = append(path, index)

		if index == from {
			found = true
			break
		}
	}
	if found {
		fmt.Printf("Path from %d to %d: %v cost(%d)\n", from, to, path, totalcost)
	} else {
		fmt.Printf("Dijkstra... No path found between %d and %d\n", from, to)
	}
}

func reverse(slice []int) []int {
	rev := sort.IntSlice(slice)
	sort.Reverse(rev)
	return rev
}
