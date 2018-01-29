package main

import (
	"fmt"
	"os"
	"parse"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		main_program(args)
	} else {
		print_usage()
	}
}

func main_program(args []string) {
	if args[0] == "-p" {
		s := parse.ReadCsv(args[1], ",")
		parse.WriteFile("view0.csv", ";", s)
	} else {
		print_usage()
	}
}

func print_usage() {
	fmt.Println("Usage: ")
	fmt.Println("       main.exe -p dataset/data.csv (to process dataset/data.csv)")
}
