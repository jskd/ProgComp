package spreadsheet

// import (
// 	"fmt"
// 	"os"
// )

type Change struct {
	Row    int
	Column int
	Value  int
}

//TODO: To reimplement this method with reference to orignal CSV file instead of double int array
// func Changes(commands []*Command, spreadSheetBefore [][]Cell,
// 	viewBefore [][]int) map[*Command][]Change {
// 	res := make(map[*Command][]Change)
// 	return res
// }

// func writeChanges(filename string,
// 	changes map[*Command][]Change) {
// 	file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
// 	defer file.Close()
// 	for command, changes := range changes {
// 		msg := fmt.Sprintf("After \"%d %d %s\":\n", command.Row,
// 			command.Column, command.Command)
// 		file.WriteString(msg)
// 		for _, change := range changes {
// 			s := fmt.Sprintf("%d %d %d\n", change.Row,
// 				change.Column, change.Value)
// 			file.WriteString(s)
// 		}
// 	}
// }
