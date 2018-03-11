// Sum type and pattern matching simulation in Go using interfaces and
// dynamic types.
package main

import "fmt"

type expr interface {
	isExpr()
}

type literal struct {
	val int
}

type add struct {
	left  expr
	right expr
}

type mult struct {
	left  expr
	right expr
}

func (l literal) isExpr() {}
func (e add) isExpr()     {}
func (e mult) isExpr()    {}

func toString(e expr) string {
	switch e := e.(type) { // Analyze the type of e at runtime.
	case literal: // Here e has type literal.
		return fmt.Sprintf("%d", e.val)
	case add: // Here e has type add.
		return toString(e.left) + " + " + toString(e.right)
	case mult: // Here e has type mult.
		return toString(e.left) + " * " + toString(e.right)
	default:
		panic("Unexpected expression.")
	}
}

func eval(e expr) int {
	switch e := e.(type) {
	case literal:
		return e.val
	case add:
		return eval(e.left) + eval(e.right)
	case mult:
		return eval(e.left) * eval(e.right)
	default:
		panic("Unexpected expression.")
	}
}

func main() {
	one := literal{1}
	two := literal{2}
	three := literal{3}
	expr := add{one, mult{two, three}}
	fmt.Printf("%s = %d\n", toString(expr), eval(expr))
}
