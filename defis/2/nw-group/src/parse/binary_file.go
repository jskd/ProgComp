package parse

import (
	"bufio"
	"encoding/binary"
	"io"
	"log"
	"os"
	"path/filepath"
	"share"
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
	count         int
	mut           sync.Mutex
}

//This function should not be directly accessable. Call singleton BinFileManager() instead
func newBinFileMgr() *BinFileMgr {
	dir2instance_map := make(map[string]*BinFile)
	mgr := BinFileMgr{dir2instance_map, 100, 0, sync.Mutex{}}
	log.Printf("BinFileMgr initialed: count=%d max_open_file=%d \n", mgr.count, mgr.max_open_file)
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

func (b *BinFileMgr) Remove(file_path string) {
	b.mut.Lock()
	defer b.mut.Unlock()
	bf := b.path2instance[file_path]
	if bf != nil {
		delete(b.path2instance, file_path)
		b.count -= 1
	}
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
	if err == nil {
		share.Purge(path)
		log.Println("Purge directory:", path)
	}

	err = os.MkdirAll(path, 0777)
	if err != nil {
		panic(err)
	}
	log.Println("New directory created:", path)
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
		log.Println("New directory created:", dirPath)
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		//log.Println("New file created:", path)
	}
}

func (b *BinFile) ReadAll() ([]uint32, error) {
	b.mut.Lock() //Ensure that no one is writing while reading
	defer b.mut.Unlock()
	if len(b.data) == 0 {
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
				//log.Printf("Loaded positions from %s %v:", b.file, b.data)
				return b.data, nil
			}
			if err != nil {
				return nil, err
			}

			d := binary.BigEndian.Uint32(bytes)
			b.data = append(b.data, d)
		}
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

func (b *BinFile) SetFinalValue(final uint32) error {
	b.mut.Lock()
	defer b.mut.Unlock()
	b.close()
	f, err_of := os.OpenFile(b.file, os.O_CREATE|os.O_RDWR, 0644)
	if err_of != nil {
		log.Fatal(err_of)
		return err_of
	}
	b.opened_f = f
	bs := uint32ToBytes(final)
	bw := bufio.NewWriter(b.opened_f)
	_, err := bw.Write(bs)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = bw.Flush()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func uint32ToBytes(val uint32) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, val)
	return bs
}

func (b *BinFile) WriteAll(list []uint32) error {
	b.mut.Lock()
	defer b.mut.Unlock()
	bw := bufio.NewWriter(b.opened_f)
	for _, val := range list {
		bs := uint32ToBytes(val)
		_, err := bw.Write(bs)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	err := bw.Flush()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (b *BinFile) WritePair(x uint32, y uint32) error {
	b.mut.Lock() //Ensure that we write x first then y
	defer b.mut.Unlock()
	bsx := uint32ToBytes(x)
	bsy := uint32ToBytes(y)
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
	bs := uint32ToBytes(v)
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

//TODO: Should notify BinFileManager()
func (b *BinFile) Delete() error {
	b.mut.Lock()
	defer b.mut.Unlock()
	b.close()
	err := os.Remove(b.file)
	if err != nil {
		return err
	}
	log.Printf("File %s deleted.\n", b.file)
	return nil
}

func (b *BinFile) Close() error {
	b.mut.Lock()
	defer b.mut.Unlock()
	b.close()
	return nil
}

func (b *BinFile) close() {
	if b.opened_f != nil {
		err := b.opened_f.Close()
		if err != nil {
			panic(err)
		}
		b.opened_f = nil
	}
}

func (b *BinFile) Move(newPath string) error {
	b.mut.Lock()
	defer b.mut.Unlock()
	b.close()
	err := os.Rename(b.file, newPath)
	if err != nil {
		return err
	}
	log.Printf("File %s renamed to %s.\n", b.file, newPath)
	b.file = newPath
	return nil
}
