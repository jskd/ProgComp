package main

import (
	"fmt"
	"os"
	"parse"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		main_program(args)
	} else {
		print_usage()
	}
}

func main_program(args []string) {
	s := parse.ReadCsv(args[0], ",")
	parse.WriteFile("view0.csv", ";", s)
}

func print_usage() {
	fmt.Println("Usage: main.exe dataset/data.csv")
}
