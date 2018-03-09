package parse

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

//Singleton implementation BinFileManager in go with var
var (
	mgr  *BinFileMgr
	once sync.Once
)

func BinFileManager() *BinFileMgr {
	once.Do(func() {
		mgr = newBinFileMgr()
	})
	return mgr
}

//Note about how struct will be stored in binary
//https://stackoverflow.com/questions/2113751/sizeof-struct-in-go
type BinFile struct {
	data     []uint32
	file     string
	opened_f *os.File
	mut      sync.Mutex
}

type BinFileMgr struct {
	path2instance map[string]*BinFile
	max_open_file int
	count         int //Ensure to access this count with atomic.AddUint64(&var, 1) and atomic.LoadUint64(&var)
	mut           sync.Mutex
}

//This function should not be directly accessable. Call singleton BinFileManager() instead
func newBinFileMgr() *BinFileMgr {
	dir2instance_map := make(map[string]*BinFile)
	mgr := BinFileMgr{dir2instance_map, 100, 0, sync.Mutex{}}
	fmt.Printf("BinFileMgr initialed: count=%d max_open_file=%d \n", mgr.count, mgr.max_open_file)
	return &mgr
}

func (b *BinFileMgr) GetBinFile(file_path string) *BinFile {
	b.mut.Lock()
	defer b.mut.Unlock()
	bf := b.path2instance[file_path]
	if bf == nil {
		if b.count >= b.max_open_file {
			err := b.saveAndClose()
			if err != nil {
				panic(err)
			}
		}
		bf = newBinaryFile(file_path)
		b.path2instance[file_path] = bf
		b.count += 1
	}
	return bf
}

func (b *BinFileMgr) SaveAndCloseAll() error {
	b.mut.Lock()
	defer b.mut.Unlock()
	return b.saveAndClose()
}

//Note that this function is not thread safe! Internal use only
func (b *BinFileMgr) saveAndClose() error {
	var err error
	for key, value := range b.path2instance {
		err = value.Close()
		delete(b.path2instance, key)
		b.count -= 1
	}
	return err
}

func PurgeAndRecreateDir(path string) {
	var _, err = os.Stat(path)
	if os.IsExist(err) {
		err = os.RemoveAll(path)
		if err != nil {
			panic(err)
		}
		fmt.Println("Purge dir:", path)
	}

	err = os.MkdirAll(path, 0777)
	if err != nil {
		panic(err)
	}
	fmt.Println("New directory created:", path)
}

//Create file if not exists
func createFileIfNotexists(path string) {
	dirPath := filepath.Dir(path)
	var _, err = os.Stat(dirPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0777)
		if err != nil {
			panic(err)
		}
		fmt.Println("New directory created:", dirPath)
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fmt.Println("New file created:", path)
	}
}

func (b *BinFile) Close() error {
	b.mut.Lock()
	defer b.mut.Unlock()
	if b.opened_f != nil {
		err := b.opened_f.Close()
		b.opened_f = nil
		return err
	}
	return nil
}

func (b *BinFile) ReadAll() ([]uint32, error) {
	b.mut.Lock() //Ensure that no one is writing while reading
	defer b.mut.Unlock()
	f, err := os.OpenFile(b.file, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	br := bufio.NewReader(f)
	for {
		bytes := make([]byte, 4)
		_, err := br.Read(bytes)
		if err == io.EOF {
			return b.data, nil
		}
		if err != nil {
			return nil, err
		}

		d := binary.BigEndian.Uint32(bytes)
		//fmt.Printf("Read %d bytes from file: %x\n", nbRead, d)
		b.data = append(b.data, d)
	}
	return b.data, nil
}

/**
   Use NewBinFile() to create a new BinFile instance.
   The directory of the given file_path will be created automatically.
**/
func NewBinFile(file_path string) *BinFile {
	return BinFileManager().GetBinFile(file_path)
}

func newBinaryFile(file_path string) *BinFile {
	createFileIfNotexists(file_path)
	f, err_of := os.OpenFile(file_path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err_of != nil {
		log.Fatal(err_of)
		return nil
	}
	bin := BinFile{[]uint32{}, file_path, f, sync.Mutex{}}
	return &bin
}

func (b *BinFile) WritePair(x uint32, y uint32) error {
	b.mut.Lock() //Ensure that we write x first then y
	defer b.mut.Unlock()
	bsx := make([]byte, 4)
	binary.BigEndian.PutUint32(bsx, x)
	bsy := make([]byte, 4)
	binary.BigEndian.PutUint32(bsy, y)
	bw := bufio.NewWriter(b.opened_f)
	_, err := bw.Write(bsx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = bw.Write(bsy)
	if err != nil {
		log.Fatal(err)
		return err
	}

	//fmt.Printf("Write %d bytes to file: %x\n", nbWritten, bs)
	err = bw.Flush()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (b *BinFile) Write(v uint32) error {
	b.mut.Lock()
	defer b.mut.Unlock()
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, v)
	bw := bufio.NewWriter(b.opened_f)
	_, err := bw.Write(bs)
	if err != nil {
		log.Fatal(err)
		return err
	}

	//fmt.Printf("Write %d bytes to file: %x\n", nbWritten, bs)
	err = bw.Flush()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (b *BinFile) Delete() error {
	b.mut.Lock()
	defer b.mut.Unlock()
	err := os.Remove(b.file)
	if err != nil {
		return err
	}
	log.Printf("File %s deleted.\n", b.file)
	return nil
}

func (b *BinFile) Move(newPath string) error {
	b.mut.Lock()
	defer b.mut.Unlock()
	err := os.Rename(b.file, newPath)
	if err != nil {
		return err
	}
	log.Printf("File %s renamed to %s.\n", b.file, newPath)
	b.file = newPath
	return nil
}
