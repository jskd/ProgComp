package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"parse"
	"share"
	"spreadsheet"
	"strings"
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
	spreadsheet.Cleanup()
	log.Println("Build directory structure of binary files...")
	bin_repo, fm := spreadsheet.FromFile(args[0], ';')
	log.Println("Directory structure of binary files built.")
	spreadsheet.Evaluate(bin_repo, false, fm)
	writeView(args[2], args[0], bin_repo)
	//commands := spreadsheet.CommandsFromFile(args[1])
	//changes := spreadsheet.Changes(commands, spreadSheet, [][]int{})
	//writeChanges(args[3], changes)
}

func print_usage() {
	fmt.Println("Usage:", "ws", "<data-file>", "<user-file>", "<view-file>",
		"<changes-file>")
}

//TODO: To write view to file_output from original file_input CSV with formula values in bin_repo
func writeView(file_output string, file_input string, bin_repo string) {
	file_in, err := os.OpenFile(file_input, os.O_RDONLY, 0644)
	share.CheckError(err)
	share.CheckError(err)
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
			share.CheckError(err) // if you return error
		}
		for i, val := range line {
			if strings.HasPrefix(val, "=") {
				fm := spreadsheet.FormulaManager("../../dataset/bin")
				f := fm.GetOrCreateFormula(val)
				line[i] = spreadsheet.GetEvaluation(f)
			}
		}
		parse.WriteOneLineCsv(file_output, line, ';')
	}
}
