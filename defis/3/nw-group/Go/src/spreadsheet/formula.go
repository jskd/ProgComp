package spreadsheet

import (
	"fmt"
	"strings"
)

// Remark: We will certainly add some fields to this structure to deal
// with dependencies.
type formula struct {
	area  area
	value uint32
	pos   position
}

func (f1 *formula) dependsOf(f2 *formula) bool {
	return f1.area.contains(f2.pos)
}

func (f *formula) String() string {
	return fmt.Sprintf("=#(%d,%d,%d,%d,%d)",
		f.area.src.col, f.area.src.row, f.area.dst.col, f.area.dst.row,
		f.value)
}

func ToFormula(s string) *formula {
	if strings.HasPrefix(s, "=") {
		return toFormula(s, "=#(%d,%d,%d,%d,%d)")
	} else {
		return toFormula(s, "%d_%d_%d_%d_%d")
	}
}

//s: input string
//p: pattern string
func toFormula(s string, p string) *formula {
	var xSrc, ySrc, xDst, yDst, val uint32
	count, _ :=
		fmt.Sscanf(s, p,
			&xSrc, &ySrc, &xDst, &yDst, &val)
	if count != 5 {
		panic(FormulaError{"Incorrect formula format:" + s + "(" + p + ")"})
	}
	if xSrc > xDst || ySrc > yDst {
		panic(FormulaError{"Source x or y must greater than destination x or y."})
	}
	area := area{
		position{col: xSrc, row: ySrc},
		position{col: xDst, row: yDst},
	}
	// TODO: Replace position{} with the real position of the
	// formula.
	return &formula{area, val, position{}}
}

func FormulaToBinFileName(s string) string {
	f := ToFormula(s)
	return fmt.Sprintf("%d/%d_%d_%d_%d_%d", f.value, f.area.src.col,
		f.area.src.row, f.area.dst.col, f.area.dst.row, f.value)
}

type FormulaError struct {
	msg string
}

func (e *FormulaError) Error() string {
	return fmt.Sprintf("FormulaError: %v", e.msg)
}
