package spreadsheet

import (
	"log"
	"sync"
)

type FormulaMgr struct {
	path2instance map[string]*formula
	count         int
	bin_dir       string
	mut           sync.Mutex
}

func FormulaManager(bin_dir string) *FormulaMgr {
	dir2instance_map := make(map[string]*formula)
	mgr := FormulaMgr{dir2instance_map, 0, bin_dir, sync.Mutex{}}
	log.Printf("FormulaMgr initialed for: %s \n", mgr.bin_dir)
	return &mgr
}

func (f *FormulaMgr) GetFormula(key string) *formula {
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
