// Go can detect race conditions during the program exucution.  To do
// this one must add the flag '-race' to the go command :
//
//    go run -race race.go
//
// See https://golang.org/doc/articles/race_detector.html
package main

var global = 13

func writer() {
	global = 42
}

func reader() int {
	return global
}

func main() {
	go reader()
	go writer()
}
