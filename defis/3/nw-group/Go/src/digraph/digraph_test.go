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

func TestTopologicalSortOnDag(t *testing.T) {
	g := New()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(4, 6)
	res1, res2 := g.TopologicalSort()
	expected1 := []int{1, 2, 3, 4, 5, 6}
	share.AssertEqual(t, len(res1)+len(res2), g.NbNodes(), "")
	share.AssertEqual(t, len(res1), len(expected1), "")
	for i := 0; i < len(res1); i++ {
		share.AssertEqual(t, res1[i], expected1[i], "")
	}
}

func TestTopologicalSortWithCycle(t *testing.T) {
	g := New()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 2)
	g.AddEdge(5, 6)
	g.AddEdge(6, 7)
	res1, res2 := g.TopologicalSort()
	expected1 := []int{5, 6, 7}
	share.AssertEqual(t, len(res1)+len(res2), g.NbNodes(), "")
	share.AssertEqual(t, len(res1), len(expected1), "")
	for i := 0; i < len(res1); i++ {
		share.AssertEqual(t, res1[i], expected1[i], "")
	}
}
