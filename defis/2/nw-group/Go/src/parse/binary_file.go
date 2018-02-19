package parse

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

type BinFile struct {
	data []uint32
	file string
}

//Create file if not exists
func createFileIfNotexists(path string) {
	//TODO: to ensure that new directory is create if not exist
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		//fmt.Println("New file created: ", path)
	}
}

func (b *BinFile) ReadAll() ([]uint32, error) {
	f, err := os.OpenFile(b.file, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for {
		bytes := make([]byte, 4)
		_, err := f.Read(bytes)
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

func NewBinFile(file_path string) *BinFile {
	createFileIfNotexists(file_path)
	return &BinFile{[]uint32{}, file_path}
}

func (b *BinFile) WritePair(x uint32, y uint32) {
	b.Write(x)
	b.Write(y)
}

//TODO: Synchronize write
func (b *BinFile) Write(v uint32) error {
	f, err_of := os.OpenFile(b.file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err_of != nil {
		log.Fatal(err_of)
		return nil
	}
	defer f.Close()
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, v)
	_, err := f.Write(bs)
	if err != nil {
		log.Fatal(err)
		return err
	}

	//fmt.Printf("Write %d bytes to file: %x\n", nbWritten, bs)
	err = f.Sync()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (b *BinFile) Delete() error {
	err := os.Remove(b.file)
	if err != nil {
		return err
	}
	fmt.Printf("File %s deleted.\n", b.file)
	return nil
}

func (b *BinFile) Move(newPath string) error {
	err := os.Rename(b.file, newPath)
	if err != nil {
		return err
	}
	fmt.Printf("File %s renamed to %s.\n", b.file, newPath)
	b.file = newPath
	return nil
}
