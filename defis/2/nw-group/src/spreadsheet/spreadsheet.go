package spreadsheet

import (
	"bufio"
	"digraph"
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

type immediate struct {
	value int
}

/*
 * Ensure that a data is either a formula or an immediate value and
 * nothing else.
 */
func (*formula) data()   {}
func (*immediate) data() {}

func (i *immediate) String() string {
	return fmt.Sprintf("%d", i.value)
}

func toData(s string) data {
	if data := toImmediate(s); data == nil {
		return nil // ToFormula(s)
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

/**
   TODO: Here the new Evaluate will process the formulas in the bin_repo
   It will return a number which represents how many formulas we are NOT aable to count (the looping formula)
   In the normal case, it should return 0 (no looping formula). In case of error, return -1.
**/
func Evaluate(bin_repo string, do_write_final_value bool, fm *FormulaMgr) []*formula {
	formula_list := setupChannels(fm.GetFormulaList())
	var wg sync.WaitGroup
	for _, f := range formula_list {
		wg.Add(1)
		// We *must* pass f to the following anonymous function, see:
		// https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		go func(f *formula) {
			defer wg.Done()
			f.evaluate(bin_repo, do_write_final_value)
		}(f)
	}
	wg.Wait()
	return formula_list
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
func FromFile(filename string, sep rune) (string, *FormulaMgr) {
	src_name := filepath.Base(filename)
	bin_dir := share.TempDir() + src_name + "/bin"
	//TODO: Skip if directory already exist
	parse.PurgeAndRecreateDir(bin_dir + "/FORMULAS")
	fm := GetAllFormulasInCsv(filename, bin_dir, sep)
	if fm.count > uint32(0) {
		log.Printf("Found total %d formulas in %s", fm.count, filename)
		m := make(map[uint32]int)
		ta_list := fm.GetFormulaList()
		for _, f := range ta_list {
			m[f.valueToCount]++
		}
		log.Printf("Pre-processing %d values for %d formulas in file %s", len(m), len(ta_list), filename)
		var wg sync.WaitGroup
		var count = 0
		for k, _ := range m {
			wg.Add(1)
			count += 1
			go PreprocessFileToBin(filename, bin_dir, sep, k, &wg)
			if count == len(m) || count%100 == 0 {
				wg.Wait() //Wait every 100 process to avoid too many opened files
			}
		}
	}
	parse.BinFileManager().SaveAndCloseAll()
	return bin_dir, fm
}

func GetAllFormulasInCsv(filename string, bin_dir string, sep rune) *FormulaMgr {
	file, err := os.Open(filename)
	defer file.Close()
	share.CheckError(err)
	fm := FormulaManager(bin_dir)
	x := uint32(0) //row
	csvReader := parse.NewCsvReader(file, sep)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		for col, str := range line {
			y := uint32(col) //column
			if strings.HasPrefix(str, "=") {
				f := fm.GetOrCreateFormula(str)
				f.positions = append(f.positions, position{x, y})
			}
		}
		x++
	}
	return fm
}

func PreprocessFileToBin(filename string, bin_dir string, sep rune, target uint32, wg *sync.WaitGroup) uint32 {
	if wg != nil {
		defer wg.Done()
	}
	count := uint32(0)
	x := uint32(0) // row
	file, err := os.Open(filename)
	defer file.Close()
	txtTarget := fmt.Sprint(target)
	bf := parse.NewBinFile(bin_dir + "/" + txtTarget)
	defer bf.Close()
	share.CheckError(err)
	log.Printf("Scaning value %s in %s\n", txtTarget, filename)
	csvReader := parse.NewCsvReader(file, sep)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		for col, str := range line {
			y := uint32(col)
			if !strings.HasPrefix(str, "=") {
				if str == txtTarget {
					bf.WritePair(x, y)
				}
			}
		}
		x++
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

func dependencyGraph(formulas []*formula) *digraph.Digraph {
	graph := digraph.New()
	for _, f1 := range formulas {
		graph.AddNode(f1)
		for _, f2 := range formulas {
			if f1.dependsOf(f2) {
				graph.AddEdge(f1, f2)
			}
		}
	}
	return graph
}

func splitFormulas(formulas []*formula) ([]*formula, []*formula) {
	asFormulas := func(array []interface{}) []*formula {
		fs := make([]*formula, len(array))
		for i := 0; i < len(fs); i++ {
			fs[i] = array[i].(*formula)
		}
		return fs
	}
	valids, invalids := dependencyGraph(formulas).TopologicalSort()
	return asFormulas(valids), asFormulas(invalids)
}

func setupChannels(formulas []*formula) []*formula {
	valids, _ := splitFormulas(formulas)
	for _, f1 := range formulas {
		for _, f2 := range formulas {
			if f1.dependsOf(f2) {
				f1.nbParents++
				f2.children = append(f2.children, f1)
			}
		}
		f1.channel = make(chan uint32, f1.nbParents)
	}
	return valids
}

func toFormulas(bin_repo string, formulas []string) []*formula {
	res := make([]*formula, len(formulas))
	for i := 0; i < len(res); i++ {
		res[i] = ToFormula(bin_repo, formulas[i])
	}
	return res
}

func Cleanup() {
	share.Purge(share.TempDir())
}
