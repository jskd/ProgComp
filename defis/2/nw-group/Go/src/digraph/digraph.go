package digraph

import (
	"container/list"
	"errors"
	"share"
)

type Digraph struct {
	neighbors map[interface{}][]interface{}
}

func New() *Digraph {
	return &Digraph{make(map[interface{}][]interface{})}
}

// Return the number of edges of the digraph.
func (g *Digraph) NbEdges() int {
	count := 0
	for _, neighbors := range g.neighbors {
		count += len(neighbors)
	}
	return count
}

// Return the number of nodes of the digraph.
func (g *Digraph) NbNodes() int {
	count := 0
	for range g.neighbors {
		count++
	}
	return count
}

func (g *Digraph) AddNode(u interface{}) {
	if _, ok := g.neighbors[u]; !ok {
		g.neighbors[u] = nil
	}
}

// TODO: Do not add an already existing edge.
func (g *Digraph) AddEdge(u interface{}, v interface{}) {
	g.AddNode(u)
	g.AddNode(v)
	g.neighbors[u] = append(g.neighbors[u], v)
}

type status int

const (
	notVisited status = iota
	visiting
	visited
)

type markedDigraph struct {
	digraph Digraph
	status  map[interface{}]status
}

func (g *Digraph) initMark() *markedDigraph {
	statuses := make(map[interface{}]status)
	for u, _ := range g.neighbors {
		statuses[u] = notVisited
	}
	return &markedDigraph{*g, statuses}
}

// Return a not yet visited node.  Return an error if there is no such
// node.
func (g *markedDigraph) pickNewNode() (interface{}, error) {
	for u, _ := range g.digraph.neighbors {
		if g.status[u] == notVisited {
			return u, nil
		}
	}
	return nil, errors.New("No such node")
}

// Perform a depth-first search on the given digraph from the given
// node.  [acc] accumulates the visited nodes in the topological order.
func (g *markedDigraph) visit(u interface{}, acc *list.List) error {
	switch g.status[u] {
	case visited:
		return nil
	case visiting: // We detect a cycle in the graph.
		return errors.New("Cycle detected")
	default:
		g.status[u] = visiting
	}
	for _, v := range g.digraph.neighbors[u] {
		err := g.visit(v, acc)
		if err != nil {
			return err
		}
	}
	g.status[u] = visited
	acc.PushFront(u)
	return nil
}

// If there are several topological sorts of the digraph, this function
// is non-deterministic because of the underlying structure of map.
func (g *Digraph) TopologicalSort() ([]interface{}, bool) {
	res := list.New()
	complete := true
	h := g.initMark()
	for u, err := h.pickNewNode(); err == nil; u, err = h.pickNewNode() {
		if h.visit(u, res) != nil {
			complete = false
			// NB: We don't clear the accumulator so that we
			// can obtain a partial topological sort of the
			// digraph.
		}
	}
	return share.ArrayFromList(res), complete
}

// TODO: Add a function to check if a given sort is a topological one
// for some graph.  This function for example useful is for testing
// purpose.
