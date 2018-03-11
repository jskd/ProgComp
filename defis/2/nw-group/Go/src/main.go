package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main_program(args []string) {
	log.Println("Build directory structure of binary files...")
	bin_repo := spreadsheet.FromFile(args[0], ';')
	log.Println("Directory structure of binary files built.")
	loop_count := spreadsheet.Evaluate(bin_repo, false)
	fmt.Printf("Looping formula: %d\n", loop_count)
	writeView(args[2], args[0], bin_repo)
	//commands := spreadsheet.CommandsFromFile(args[1])
	//changes := spreadsheet.Changes(commands, spreadSheet, [][]int{})
	//writeChanges(args[3], changes)
}

func print_usage() {
	fmt.Println("Usage: ")
	fmt.Println("ws <data-file> <user-file> <view-file> <changes-file>")
}

func writeChanges(filename string,
	changes map[*spreadsheet.Command][]spreadsheet.Change) {
	file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	for command, changes := range changes {
		msg := fmt.Sprintf("After \"%d %d %s\":\n", command.Row,
			command.Column, command.Command)
		file.WriteString(msg)
		for _, change := range changes {
			s := fmt.Sprintf("%d %d %d\n", change.Row,
				change.Column, change.Value)
			file.WriteString(s)
		}
	}
}

//TODO: To write view to file_output from original file_input CSV with formula values in bin_repo
func writeView(file_output string, file_input string, bin_repo string) {
	file_in, err := os.OpenFile(file_input, os.O_RDONLY, 0644)
	checkError(err)
	checkError(err)
	defer file_in.Close()
	reader := bufio.NewReader(file_in)
	reader2 := csv.NewReader(reader)
	reader2.Comma = ';'
	reader2.FieldsPerRecord = -1
	for {
		line, err := reader2.Read() // 0x0A separator = newline

		if err == io.EOF {
			println("fin de fichier")
			break
		} else if err != nil {
			checkError(err) // if you return error
		}
		parse.WriteOneLineCsv(file_output, line, ';')
	}
}
