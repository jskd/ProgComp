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
	file *os.File
}

func createFileIfNotexists(path string) {
	// detect if file exists
	var _, err = os.Stat(path)
	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()
	}

	fmt.Println("New file created: ", path)
}

func (b *BinFile) ReadAll() ([]uint32, error) {
	for {
		record, err := b.Read()
		fmt.Println(record)
		if err == io.EOF {
			return b.data, nil
		}
		if err != nil {
			return nil, err
		}
		b.data = append(b.data, record)
	}
	return b.data, nil
}

//TODO: to ensure that new directory is create if not exist
func NewBinFile(file_path string) *BinFile {
	createFileIfNotexists(file_path)
	file, err := os.OpenFile(file_path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &BinFile{[]uint32{}, file}
}

func (b *BinFile) Close() {
	b.file.Close()
}

func (b *BinFile) ReadPair() (uint32, uint32, error) {
	x, ex := b.Read()
	if ex != nil {
		log.Fatal(ex)
		return 0, 0, ex
	}
	y, ey := b.Read()
	if ey != nil {
		log.Fatal(ey)
		return 0, 0, ey
	}
	return x, y, nil
}

func (b *BinFile) Read() (uint32, error) {
	bytes := make([]byte, 4)
	_, err := b.file.Read(bytes)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	data := binary.BigEndian.Uint32(bytes)
	return data, err
}

//TODO: Synchronize write
func (b *BinFile) Write(x uint32, y uint32) {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, x)
	_, err := b.file.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
}
