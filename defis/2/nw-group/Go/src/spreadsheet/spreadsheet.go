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
	"sync"
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
		f.xSource, f.ySource, f.xDestination, f.yDestination, f.value)
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
	return &formula{xSrc, ySrc, xDst, yDst, val}
}

/**
   TODO: Here the new Evaluate will process the formulas in the bin_repo
   It will return a number which represents how many formulas we are NOT aable to count (the looping formula)
   In the normal case, it should return 0 (no looping formula). In case of error, return -1.
**/
func Evaluate(bin_repo string, do_write_final_value bool) int {
	formula_list, _, err := getBinFormulaList(bin_repo)
	if err == nil {
		formula_list = SortByDependency(formula_list)
		var wg sync.WaitGroup
		for _, f := range formula_list {
			wg.Add(1)
			go EvaluateFormula(bin_repo, f, do_write_final_value, &wg)
		}
		wg.Wait()
	}
	return len(formula_list)
}

//Return list of formula files and list of formula's target value, and error if occured
func getBinFormulaList(bin_repo string) ([]string, []string, error) {
	files, err := listAllFilesInDir(bin_repo + "/FORMULAS/")
	if err != nil {
		log.Fatal(err)
		return []string{}, []string{}, err
	}

	var formula_list []string
	f_value_list := make([]string, len(files))
	for idx, d := range files {
		f_value_list[idx] = d
		f_names, _ := listAllFilesInDir(bin_repo + "/FORMULAS/" + d)
		for _, f := range f_names {
			formula_list = append(formula_list, f)
		}
	}
	return formula_list, f_value_list, nil
}

func listAllFilesInDir(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return []string{}, err
	}

	file_names := make([]string, len(files))
	for idx, f := range files {
		file_names[idx] = f.Name()
	}
	return file_names, nil
}

//TODO: To sort fomulas with dependency algorithm
func SortByDependency(f_list []string) []string {
	return f_list
}

//Pre-process Csv file to value-position binary files.
//Return a directory where value-position binary files are stored
func FromFile(filename string, sep rune) string {
	src_name := filepath.Base(filename)
	bin_dir := share.TempDir() + src_name + "/bin"
	//TODO: Skip if directory already exist
	parse.PurgeAndRecreateDir(bin_dir)
	formula_count := PreprocessFileToBin(filename, bin_dir, sep, nil)
	if formula_count > uint32(0) {
		log.Printf("Found %d formulas in %s", formula_count, filename)
		m := make(map[string]int)
		_, ta_list, err := getBinFormulaList(bin_dir)
		if err == nil {
			for _, v := range ta_list {
				m[v]++
			}
		}
		PreprocessFileToBin(filename, bin_dir, sep, m)
	}
	parse.BinFileManager().SaveAndCloseAll()
	return bin_dir
}

func PreprocessFileToBin(filename string, bin_dir string, sep rune, target map[string]int) uint32 { //When target is nil, extract only formulas
	count := uint32(0)
	csvParser := parse.NewCsvParser(filename, sep, '"')
	for {
		str, x, y, err := csvParser.ReadOneCell()
		//fmt.Printf("%d,%d: '%s'\n", x, y, str)
		if err == io.EOF {
			break
		}

		if target == nil {
			if saveOneFormulaToBin(bin_dir, str, x, y) {
				count++
			}
		} else {
			if target[str] > 0 {
				if saveOneCellToBin(bin_dir, str, x, y) {
					count++
				}
			}
		}
	}
	return count
}

func saveOneFormulaToBin(bin_dir string, txt string, pos_x uint32, pos_y uint32) bool {
	if strings.HasPrefix(txt, "=") {
		bf := parse.NewBinFile(bin_dir + "/FORMULAS/" + FormulaToBinFileName(txt))
		bf.WritePair(pos_x, pos_y)
		return true
	}
	return false
}

func saveOneCellToBin(bin_dir string, txt string, pos_x uint32, pos_y uint32) bool {
	if !strings.HasPrefix(txt, "=") {
		bf := parse.NewBinFile(bin_dir + "/" + txt)
		bf.WritePair(pos_x, pos_y)
		return true
	}
	return false
}

func FormulaToBinFileName(s string) string {
	f := ToFormula(s)
	return fmt.Sprintf("%d/%d_%d_%d_%d_%d", f.value, f.xSource, f.ySource, f.xDestination, f.yDestination, f.value)
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

func EvaluateFormula(bin_repo string, formula_name string, do_write_final_value bool, wg *sync.WaitGroup) uint32 {
	if wg != nil { //When WaitGroup is not provided, nothing to notify
		defer wg.Done()
	}
	formulaBin := parse.NewBinFile(bin_repo + "/FORMULAS/" + FormulaToBinFileName(formula_name))
	formulaPos, err := formulaBin.ReadAll()
	if err != nil {
		panic(err)
	}
	if len(formulaPos) == 1 {
		//If this formula has been evaluated, return its value directly
		return formulaPos[0]
	}
	f := ToFormula(formula_name)
	binFile := parse.NewBinFile(bin_repo + "/" + fmt.Sprint(f.value))
	pos, err := binFile.ReadAll()
	if err != nil {
		panic(err)
	}

	//Count from position range
	count := uint32(0)
	var x uint32
	var y uint32
	for indx, val := range pos {
		if indx/2 == 1 {
			y = val
			//fmt.Printf("(%d, %d) vs (%d, %d, %d, %d)\n", x, y, f.xSource, f.ySource, f.xDestination, f.yDestination)
			if x <= f.xDestination && x >= f.xSource &&
				y <= f.yDestination && y >= f.ySource {
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

	if do_write_final_value {
		err = formulaBin.SetFinalValue(count)
		if err != nil {
			panic(err)
		}
	}
	return count
}
