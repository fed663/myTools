package graph

import (
	"github.com/fed663/myTools/deque"
)

type Graph[T comparable] struct {
	adj map[T][]T
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{adj: make(map[T][]T)}
}

func (g *Graph[T]) AddDirEdge(u, w T) {
	g.adj[u] = append(g.adj[u], w)
}

func (g *Graph[T]) AddUnDirEdge(u, w T) {
	g.adj[u] = append(g.adj[u], w)
	g.adj[w] = append(g.adj[w], u)
}

func (g *Graph[T]) DFSRec(start T) {
	marked := make(map[T]bool)
	g.dfs(start, marked)
}

func (g *Graph[T]) dfs(u T, marked map[T]bool) {
	marked[u] = true
	for _, w := range g.adj[u] {
		if !marked[w] {
			g.dfs(w, marked)
		}
	}
}
func (g *Graph[T]) DFSStack(start T) {
	marked := make(map[T]bool)
	stack := deque.NewDeque[T]()
	stack.PushBack(start)

	for stack.Len() > 0 {
		u, _ := stack.PopBack()
		if !marked[u] {
			marked[u] = true
			for _, w := range g.adj[u] {
				if !marked[w] {
					stack.PushBack(w)
				}
			}
		}
	}
}

func (g *Graph[T]) BFS(start T) {
	marked := make(map[T]bool)
	g.bfs(start, marked)
}

func (g *Graph[T]) bfs(start T, marked map[T]bool) {
	deq := deque.NewDeque[T]()
	deq.PushBack(start)
	for deq.Len() > 0 {
		u, _ := deq.PopFront()
		if !marked[u] {
			marked[u] = true
			for _, w := range g.adj[u] {
				if !marked[w] {
					deq.PushBack(w)
				}
			}
		}
	}
}
