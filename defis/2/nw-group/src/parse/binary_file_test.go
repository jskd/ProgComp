package parse

import (
	"share"
	"testing"
)

func TestBinarySetFinalValue(t *testing.T) {
	//TODO
}

func TestBinaryReadWriteDelete(t *testing.T) {
	bf := newBinaryFile("../../dataset/bin/test")
	bf.WritePair(0, 1)
	bf.WritePair(2, 3)
	share.AssertEqual(t, bf.Close(), nil, "Unable to close file.")

	bf = NewBinFile("../../dataset/bin/test")
	arr, err := bf.ReadAll()
	share.AssertEqual(t, bf.Close(), nil, "Unable to close file.")
	BinFileManager().SaveAndCloseAll()
	exp := []uint32{}
	exp = append(exp, 0)
	exp = append(exp, 1)
	exp = append(exp, 2)
	exp = append(exp, 3)

	share.AssertEqual(t, bf.Move("../../dataset/bin/1"), nil, "Unable to move file.")
	share.AssertEqual(t, bf.Delete(), nil, "Unable to delete file.")
	share.AssertEqual(t, err, nil, "Some error happened while reading file.")
	share.AssertEqual(t, arr, exp, "Incorrect uint32 value.")
}

func TestBinFileMgr(t *testing.T) {
	bf1 := NewBinFile("../../dataset/bin/test1")
	bf2 := NewBinFile("../../dataset/bin/test2")
	bf1.WritePair(4, 5)
	bf1.WritePair(6, 7)
	bf2.WritePair(4, 5)
	bf2.WritePair(6, 7)
	share.AssertEqual(t, BinFileManager().count, 2, "BinFileManager contains incorrect open files.")
	bf1.Delete()
	BinFileManager().SaveAndCloseAll()
	bf2.Delete()
	share.AssertEqual(t, BinFileManager().count, 0, "BinFileManager contains incorrect open files after close")
}
