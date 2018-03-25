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
	digraph  Digraph
	status   map[interface{}]status
	worklist *list.List
}

func (g *Digraph) initMarking() *markedDigraph {
	status := make(map[interface{}]status)
	worklist := list.New()
	for u, _ := range g.neighbors {
		status[u] = notVisited
		worklist.PushFront(u)
	}
	return &markedDigraph{*g, status, worklist}
}

// Return a not yet visited node.  Return an error if there is no such
// node.
func (g *markedDigraph) pickNewNode() (interface{}, error) {
	if u := g.worklist.Front(); u != nil {
		return g.worklist.Remove(u), nil
	}
	return nil, errors.New("No such node")
}

// Perform a depth-first search on the given digraph from the given
// node.  [done] accumulates the visited nodes in the topological order.
// [doing] contains the nodes that are being visited.
func (g *markedDigraph) visit(u interface{}, done *list.List,
	doing *list.List) error {
	switch g.status[u] {
	case visited:
		return nil
	case visiting: // We detect a cycle in the graph.
		return errors.New("Cycle detected")
	default:
		g.status[u] = visiting
		doing.PushFront(u)
	}
	for _, v := range g.digraph.neighbors[u] {
		if err := g.visit(v, done, doing); err != nil {
			return err
		}
	}
	g.status[u] = visited
	if doing.Remove(doing.Front()) != u {
		panic("Assertion failed")
	}
	done.PushFront(u)
	return nil
}

// If there are several topological sorts of the digraph, this function
// is non-deterministic because of the underlying structure of map.
func (graph *Digraph) TopologicalSort() ([]interface{}, []interface{}) {
	visited := list.New()
	visiting := list.New()
	g := graph.initMarking()
	for u, err := g.pickNewNode(); err == nil; u, err = g.pickNewNode() {
		g.visit(u, visited, visiting)
	}

	// Here, [visiting] contains the list of the nodes of [g] that
	// have a path to a cycle of [g].
	return share.ArrayFromList(visited), share.ArrayFromList(visiting)
}

// TODO: Add a function to check if a given sort is a topological one
// for some graph.  This function for example useful is for testing
// purpose.
