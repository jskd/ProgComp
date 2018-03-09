package spreadsheet

import (
	"bufio"
	//"encoding/csv"
	"fmt"
	"io"
	"os"
	"parse"
	"path/filepath"
	"share"
	"strconv"
	"strings"
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

type FormulaError struct {
	msg string
}

func (e *FormulaError) Error() string {
	return fmt.Sprintf("FormulaError: %v", e.msg)
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
		return ToFormula(s)
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

func ToFormula(s string) *formula {
	var xSrc, ySrc, xDst, yDst, val int

	count, _ :=
		fmt.Sscanf(s, "=#(%d,%d,%d,%d,%d)",
			&ySrc, &xSrc, &yDst, &xDst, &val)
	if count != 5 {
		panic(FormulaError{"Incorrect formula format:" + s})
	}
	if xSrc > xDst || ySrc > yDst {
		panic(FormulaError{"Source x or y must greater than destination x or y."})
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

/**
   TODO: Here the new Evaluate will process the formulas in the bin_repo
   It will return a number which represents how many formulas we are NOT aable to count (the looping formula)
   In the normal case, it should return 0 (no looping formula). In case of error, return -1.
**/
func Evaluate(bin_repo string) int {
	return 0
}

//Pre-process Csv file to value-position binary files.
//Return a directory where value-position binary files are stored
func FromFile(filename string, sep rune) string {
	src_name := filepath.Base(filename)
	bin_dir := share.TempDir() + src_name + "/bin"
	//TODO: Skip if directory already exist
	parse.PurgeAndRecreateDir(bin_dir)
	//1st implementation to process CSV with Go CSV lib.
	//	file, err := os.Open(filename)
	//	if err != nil {
	//		panic(err)
	//	}
	//	csvReader := csv.NewReader(file)
	//	csvReader.Comma = sep
	//	csvReader.FieldsPerRecord = -1
	//	csvReader.TrimLeadingSpace = true
	//	err = nil
	//	for pos_x := 0; err != io.EOF; pos_x += 1 {
	//		var l []string
	//		l, err = csvReader.Read()
	//		processOneLineOfCsv(bin_dir, l, sep, pos_x)
	//		fmt.Printf("Processing %d\n", len(l))
	//	}
	//2nd implementation to process CSV with customized parser.
	csvParser := parse.NewCsvParser(filename, sep, '"')

	for {
		str, x, y, err := csvParser.ReadOneCell()
		//fmt.Printf("%d,%d: '%s'\n", x, y, str)
		if err == io.EOF {
			break
		}

		saveOneCellToBin(bin_dir, str, x, y)
	}

	parse.BinFileManager().SaveAndCloseAll()
	return bin_dir
}

func saveOneCellToBin(bin_dir string, txt string, pos_x uint32, pos_y uint32) {
	if strings.HasPrefix(txt, "=") {
		bf := parse.NewBinFile(bin_dir + "/FORMULAS/" + FormulaToFileName(txt))
		bf.WritePair(pos_x, pos_y)
	} else {
		bf := parse.NewBinFile(bin_dir + "/" + txt)
		bf.WritePair(pos_x, pos_y)
	}
}

func processOneLineOfCsv(bin_dir string, line []string, sep rune, pos_x int) {
	for pos_y, element := range line {
		txt := strings.TrimSpace(element)
		saveOneCellToBin(bin_dir, txt, uint32(pos_x), uint32(pos_y))
		pos_y += 1
	}
}

func FormulaToFileName(s string) string {
	f := ToFormula(s)
	return fmt.Sprintf("%d_%d_%d_%d_%d", f.ySource, f.xSource, f.yDestination, f.xDestination, f.value)
}

type Command struct {
	Row     int
	Column  int
	Command data
}

type Change struct {
	Row    int
	Column int
	Value  int
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

//TODO: To reimplement this method with reference to orignal CSV file instead of double int array
func Changes(commands []*Command, spreadSheetBefore [][]Cell,
	viewBefore [][]int) map[*Command][]Change {
	res := make(map[*Command][]Change)
	return res
}

func Intersect(path string, formulaName string) int {
	form := strings.Split(formulaName, "_")
	x1, err := strconv.Atoi(form[0])
	x2, err := strconv.Atoi(form[1])
	y1, err := strconv.Atoi(form[2])
	y2, err := strconv.Atoi(form[3])
	valToCount := form[4]
	binFile := parse.NewBinFile(string("../../dataset/bin/" + valToCount))
	tab, err := binFile.ReadAll()
	if err != nil {
		panic(err)
	}
	count := 0
	for i := 0; i < len(tab); i++ {
		x := int(tab[i])
		i++
		y := int(tab[i])
		if x < x1 && x > x2 {
			if y < y1 && y < y2 {
				count++
			}
		}
	}
	formule := parse.NewBinFile(path + "" + formulaName)
	tab, err = formule.ReadAll()
	if err != nil {
		panic(err)
	}
	parse.NewBinFile(string("../../dataset/bin/" + string(count)))
	for i := 0; i < len(tab); i++ {
		x := tab[i]
		i++
		y := tab[i]
		binFile.WritePair(x, y)
	}
	return count
}
