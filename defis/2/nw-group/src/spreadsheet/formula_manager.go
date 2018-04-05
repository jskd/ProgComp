package spreadsheet

import (
	"log"
	"sync"
)

type FormulaMgr struct {
	path2instance map[string]*formula
	count         uint32
	bin_dir       string
	mut           sync.Mutex
}

func FormulaManager(bin_dir string) *FormulaMgr {
	dir2instance_map := make(map[string]*formula)
	mgr := FormulaMgr{dir2instance_map, uint32(0), bin_dir, sync.Mutex{}}
	log.Printf("FormulaMgr initialed for: %s \n", mgr.bin_dir)
	return &mgr
}

func (f *FormulaMgr) GetOrCreateFormula(key string) *formula {
	f.mut.Lock()
	defer f.mut.Unlock()
	bf := f.path2instance[key]
	if bf == nil {
		bf = ToFormula("", key)
		f.path2instance[key] = bf
		f.count += 1
	}
	return bf
}

func (f *FormulaMgr) GetFormulaList() []*formula {
	f.mut.Lock()
	defer f.mut.Unlock()
	var flist = make([]*formula, 0)
	for _, v := range f.path2instance {
		flist = append(flist, v)
	}
	return flist
}

func (b *FormulaMgr) Remove(key string) {
	b.mut.Lock()
	defer b.mut.Unlock()
	bf := b.path2instance[key]
	if bf != nil {
		delete(b.path2instance, key)
		b.count -= 1
	}
}

//Return list of formula files and list of formula's target value, and error if occured
func (b *FormulaMgr) getBinFormulaList() ([]string, []string, error) {
	files, err := listAllFilesInDir(b.bin_dir + "/FORMULAS/")
	if err != nil {
		log.Fatal(err)
		return []string{}, []string{}, err
	}

	var formula_list []string
	f_value_list := make([]string, len(files))
	for idx, d := range files {
		f_value_list[idx] = d
		f_names, _ := listAllFilesInDir(b.bin_dir + "/FORMULAS/" + d)
		for _, f := range f_names {
			formula_list = append(formula_list, f)
		}
	}
	return formula_list, f_value_list, nil
}
