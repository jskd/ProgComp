package main

import "fmt"

func newFreshIdGenerator(prefix string) func() string {
	c := 0
	return func() string {
		c++
		return fmt.Sprintf("%s%d", prefix, c)
	}
}

func main() {
	freshId := newFreshIdGenerator("x")
	fmt.Println(freshId())
	fmt.Println(freshId())
	freshId = newFreshIdGenerator("y")
	fmt.Println(freshId())
	fmt.Println(freshId())
}
