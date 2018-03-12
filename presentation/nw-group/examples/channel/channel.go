// Connect goroutines together using channels, similar to Unix pipes.
// Source:
// https://github.com/adonovan/gopl.io/blob/master/ch8/pipeline1/main.go
package main

import "fmt"

func main() {
	// Bidirectional channels
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x // Send a value to a channel.
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals // Receive a value from a channel.
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}
