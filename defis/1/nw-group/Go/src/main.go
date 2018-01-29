package main

import (
	"fmt"
	"os"
	"parse"
	"spreadsheet"
)

func main() {
	args := os.Args[1:]
	if len(args) == 4 {
		main_program(args)
	} else {
		print_usage()
	}
}

func main_program(args []string) {
	spreadSheet := spreadsheet.FromFile(args[0])
	writeView(args[2], spreadsheet.Evaluate(spreadSheet))
}

func print_usage() {
	fmt.Println("Usage: ")
	fmt.Println("ws <data-file> <user-file> <view-file> <changes-file>")
}

func writeView(filename string, values [][]int) {
	res := make([][]string, len(values))
	for r, row := range values {
		for _, value := range row {
			var s string

			if value >= 0 {
				s = fmt.Sprintf("%d", value)
			} else {
				s = "P"
			}
			res[r] = append(res[r], s)
		}
	}
	parse.WriteCsv(filename, res, ';')
}
