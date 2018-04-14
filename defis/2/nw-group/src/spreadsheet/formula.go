package spreadsheet

import (
	"fmt"
	"log"
	"parse"
	"strings"
)

type formula struct {
	area         area
	valueToCount uint32
	positions    []position
	evaluation   uint32
	nbParents    int
	children     []*formula
	channel      chan uint32
}

func (f *formula) accumulateValue(val uint32) {
	if val == f.valueToCount {
		f.evaluation++
	}
}

func GetEvaluation(f *formula) string{
	return string(int(f.evaluation))
}

func (f *formula) accumulatePosition(pos position) {
	if f.area.contains(pos) {
		f.evaluation++
	}
}

func (f1 *formula) dependsOf(f2 *formula) bool {
	for _, pos := range f2.positions {
		if f1.area.contains(pos) {
			return true
		}
	}
	return false
}

func (f *formula) evaluate(bin_repo string, do_write_final_value bool) {
	binFile := parse.NewBinFile(bin_repo + "/" + fmt.Sprint(f.valueToCount))
	pos, err := readAllPos(binFile)
	if err != nil {
		panic(err)
	}

	//Count from position range
	for _, p := range pos {
		f.accumulatePosition(p)
	}

	for i := 0; i < f.nbParents; i++ {
		f.accumulateValue(<-f.channel)
	}

	log.Println(f, "evaluated into", f.evaluation)

	for _, child := range f.children {
		child.channel <- f.evaluation
	}

	//Write position to value bin file
	// valueBinFile := parse.NewBinFile(bin_repo + "/" + fmt.Sprint(f.evaluation))
	// err = valueBinFile.WriteAll(formulaPos)
	// if err != nil {
	//	panic(err)
	// }

	// if do_write_final_value {
	//	err = formulaBin.SetFinalValue(f.evaluation)
	//	if err != nil {
	//		panic(err)
	//	}
	// }
}

func (f *formula) String() string {
	return fmt.Sprintf("=#(%d,%d,%d,%d,%d)",
		f.area.src.row, f.area.src.col, f.area.dst.row, f.area.dst.col,
		f.valueToCount)
}

func ToFormula(bin_repo string, s string) *formula {
	if strings.HasPrefix(s, "=") {
		return toFormula(bin_repo, s, "=#(%d,%d,%d,%d,%d)")
	} else {
		return toFormula(bin_repo, s, "%d_%d_%d_%d_%d")
	}
}

//s: input string
//p: pattern string
func toFormula(bin_repo string, s string, p string) *formula {
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
		position{row: xSrc, col: ySrc},
		position{row: xDst, col: yDst},
	}
	// TODO: Replace position{} with the real position of the
	// formula.
	if bin_repo == "" {
		return &formula{area, val, make([]position, 0), uint32(0), 0, make([]*formula, 0), nil}
	} else {
		f := parse.NewBinFile(bin_repo + "/FORMULAS/" + FormulaToBinFileName(s))
		pos, _ := readAllPos(f)
		return &formula{area, val, pos, uint32(0), 0, make([]*formula, 0), nil}
	}
}

func (f *formula) formulaToBinFileName() string {
	return fmt.Sprintf("%d/%d_%d_%d_%d_%d", f.valueToCount, f.area.src.row,
		f.area.src.col, f.area.dst.row, f.area.dst.col, f.valueToCount)
}

func FormulaToBinFileName(s string) string {
	f := ToFormula("", s)
	return fmt.Sprintf("%d/%d_%d_%d_%d_%d", f.valueToCount, f.area.src.row,
		f.area.src.col, f.area.dst.row, f.area.dst.col, f.valueToCount)
}

type FormulaError struct {
	msg string
}

func (e *FormulaError) Error() string {
	return fmt.Sprintf("FormulaError: %v", e.msg)
}

func readAllPos(b *parse.BinFile) ([]position, error) {
	pos, err := b.ReadAll()
	if err != nil || len(pos)&1 != 0 {
		return nil, err
	}
	res := make([]position, 0)
	for i := uint32(1); i < uint32(len(pos)); i += 2 {
		res = append(res, position{pos[i-1], pos[i]})
	}
	return res, nil
}
