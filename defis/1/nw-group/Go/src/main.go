package main

import (
	"fmt"
	"os"
	"parse"
	"spreadsheet"
)

/* TODO: Check that output files are different from input ones.  */
func main() {
	args := os.Args[1:]
	if len(args) == 4 {
		main_program(args)
	} else {
		print_usage()
		os.Exit(1)
	}
}

func main_program(args []string) {
	spreadSheet := spreadsheet.FromFile(args[0])
	view := spreadsheet.Evaluate(spreadSheet)
	writeView(args[2], view)
	commands := spreadsheet.CommandsFromFile(args[1])
	changes := spreadsheet.Changes(commands, spreadSheet, view)
	writeChanges(args[3], changes)
}

func print_usage() {
	fmt.Println("Usage: ")
	fmt.Println("ws <data-file> <user-file> <view-file> <changes-file>")
}

func writeChanges(filename string,
	changes map[*spreadsheet.Command][]spreadsheet.Change) {
	file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	for command, changes := range changes {
		msg := fmt.Sprintf("after \"%d %d %s\":\n", command.Row,
			command.Column, command.Command)
		file.WriteString(msg)
		for _, change := range changes {
			s := fmt.Sprintf("%d %d %d\n", change.Row,
				change.Column, change.Value)
			file.WriteString(s)
		}
	}
	file.Close()
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
