package digraph

import (
	"share"
	"testing"
)

func TestAddEdge(t *testing.T) {
	g := New()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	share.AssertEqual(t, g.NbNodes(), 4, "")
	share.AssertEqual(t, g.NbEdges(), 3, "")
}

func TestInitMark(t *testing.T) {
	g := New()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	h := g.initMark()
	share.AssertEqual(t, g.NbNodes(), h.digraph.NbNodes(), "")
	share.AssertEqual(t, g.NbEdges(), h.digraph.NbEdges(), "")
}

func TestTopologicalSort(t *testing.T) {
	g := New()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(4, 6)
	res, complete := g.TopologicalSort()
	expected := []int{1, 2, 3, 4, 5, 6}
	share.AssertEqual(t, complete, true, "")
	share.AssertEqual(t, len(res), len(expected), "")
	for i := 0; i < len(expected); i++ {
		share.AssertEqual(t, res[i], expected[i], "")
	}
}
