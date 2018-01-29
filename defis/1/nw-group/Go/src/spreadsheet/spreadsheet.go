package spreadsheet

import (
	"fmt"
	"strconv"

	"parse"
)

type Cell struct {
	data       data
	evaluating bool
}

/*
 * Data may be an immediate value (an integer literal) or formula to
 * evaluate.
 */
type data interface {
	data()
}

/* Strucure of a formula */
type formula struct {
	xSource      int
	ySource      int
	xDestination int
	yDestination int
	value        int
}

type immediate struct {
	value int
}

/*
 * Ensure that a data is either a formula or an immediate value and
 * nothing else.
 */
func (*formula) data()   {}
func (*immediate) data() {}

func (f *formula) String() string {
	return fmt.Sprintf("=#(%d,%d,%d,%d,%d)",
		f.ySource, f.xSource, f.yDestination, f.xDestination, f.value)
}

func (i *immediate) String() string {
	return fmt.Sprintf("%d", i.value)
}

func toData(s string) data {
	if data := toImmediate(s); data == nil {
		return toFormula(s)
	} else {
		return data
	}
}

func toImmediate(s string) *immediate {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &immediate{i}
}

func toFormula(s string) *formula {
	var xSrc, ySrc, xDst, yDst, val int

	count, _ :=
		fmt.Sscanf(s, "=#(%d,%d,%d,%d,%d)",
			&ySrc, &xSrc, &yDst, &xDst, &val)
	if count != 5 {
		return nil
	}
	return &formula{xSrc, ySrc, xDst, yDst, val}
}

/* Evaluates a formula.  On success, a formula evaluates into an
   integer.  Returns an error when the dependency graph of the formula
   has a cycle.  */
func (formula *formula) evaluate(spreadSheet [][]Cell) int {
	count := 0
	if formula == nil {
		return -1
	}
	for r := formula.ySource; r <= formula.yDestination; r++ {
		for c := formula.xSource; c <= formula.xDestination; c++ {
			if spreadSheet[r][c].evaluating {
				return -1
			}
			switch evaluate(r, c, spreadSheet) {
			case -1:
				return -1
			case formula.value:
				count++
			default:
				continue
			}
		}
	}
	return count
}

func evaluate(row int, col int, spreadSheet [][]Cell) int {
	var res int

	cell := &spreadSheet[row][col]
	switch data := cell.data.(type) {
	case *formula:
		cell.evaluating = true
		res = data.evaluate(spreadSheet)
		cell.evaluating = false
	case *immediate:
		res = data.value
	default:
		res = -1
	}
	return res
}

func cells(spreadSheet [][]string) [][]Cell {
	res := make([][]Cell, len(spreadSheet))
	for r, row := range spreadSheet {
		for _, s := range row {
			res[r] = append(res[r], Cell{toData(s), false})
		}
	}
	return res
}

func Evaluate(spreadSheet [][]Cell) [][]int {
	res := make([][]int, len(spreadSheet))
	for r, row := range spreadSheet {
		for c := range row {
			res[r] = append(res[r], evaluate(r, c, spreadSheet))
		}
	}
	return res
}

func FromFile(filename string) [][]Cell {
	spreadSheet := parse.ReadCsv(filename, ';')
	return cells(spreadSheet)
}
