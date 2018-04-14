package spreadsheet

import (
	"digraph"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"parse"
	"path/filepath"
	"share"
	"strings"
	"sync"
)

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
