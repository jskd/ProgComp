package spreadsheet

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
	xSource      uint32
	ySource      uint32
	xDestination uint32
	yDestination uint32
	value        uint32
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

func BinFileToFormula(s string) *formula {
	return toFormula(s, "%d_%d_%d_%d_%d")
}

func ToFormula(s string) *formula {
	return toFormula(s, "=#(%d,%d,%d,%d,%d)")
}

//s: input string
//p: pattern string
func toFormula(s string, p string) *formula {
	var xSrc, ySrc, xDst, yDst, val uint32
	count, _ :=
		fmt.Sscanf(s, p,
			&ySrc, &xSrc, &yDst, &xDst, &val)
	if count != 5 {
		panic(FormulaError{"Incorrect formula format:" + s + "(" + p + ")"})
	}
	if xSrc > xDst || ySrc > yDst {
		panic(FormulaError{"Source x or y must greater than destination x or y."})
	}
	return &formula{xSrc, ySrc, xDst, yDst, val}
}

/**
   TODO: Here the new Evaluate will process the formulas in the bin_repo
   It will return a number which represents how many formulas we are NOT aable to count (the looping formula)
   In the normal case, it should return 0 (no looping formula). In case of error, return -1.
**/
func Evaluate(bin_repo string) int {
	files, err := ioutil.ReadDir(bin_repo + "/FORMULAS/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	return len(files)
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

func EvaluateFormula(bin_repo string, formulaName string) uint32 {
	formulaBin := parse.NewBinFile(bin_repo + "/FORMULAS/" + formulaName)
	formulaPos, err := formulaBin.ReadAll()
	if err != nil {
		panic(err)
	}
	if len(formulaPos) == 1 {
		//If this formula has been evaluated, return its value directly
		return formulaPos[0]
	}
	f := BinFileToFormula(formulaName)
	binFile := parse.NewBinFile(bin_repo + "/" + fmt.Sprint(f.value))
	pos, err := binFile.ReadAll()
	if err != nil {
		panic(err)
	}
	count := uint32(0)
	var x uint32
	var y uint32
	for indx, val := range pos {
		if indx/2 == 1 {
			y = val
			if x <= f.xDestination && x >= f.xSource &&
				y <= f.yDestination && y <= f.ySource {
				count++
			}
		} else {
			x = val
		}
	}

	//Write position to value bin file
	valueBinFile := parse.NewBinFile(bin_repo + "/" + fmt.Sprint(count))
	err = valueBinFile.WriteAll(formulaPos)
	if err != nil {
		panic(err)
	}

	//TODO: Write value to formula file

	return count
}
