package spreadsheet

import (
	"bufio"
	"fmt"
	"os"
)

type Command struct {
	Row     int
	Column  int
	Command data
}

func toCommand(s string) *Command {
	var row, col int
	var command string

	count, _ :=
		fmt.Sscanf(s, "%d %d %s",
			&row, &col, &command)
	if count != 3 {
		return nil
	}
	return &Command{row, col, toData(command)}
}

func CommandsFromFile(filename string) []*Command {
	var commands []*Command
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commands = append(commands, toCommand(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return commands
}

func applyCommand(command *Command, spreadSheetBefore [][]Cell,
	view [][]int) [][]Cell {
	spreadSheet := make([][]Cell, len(view))

	for r, row := range view {
		spreadSheet[r] = make([]Cell, len(row))
	}
	for r, row := range view {
		for c, value := range row {
			if r == command.Row && c == command.Column {
				spreadSheet[r][c] = Cell{command.Command, false}
			} else {
				switch spreadSheetBefore[r][c].data.(type) {
				case *formula:
					spreadSheet[r][c] =
						spreadSheetBefore[r][c]
				default:
					spreadSheet[r][c] =
						Cell{&immediate{value}, false}
				}
			}
		}
	}
	return spreadSheet
}
