package parse

import (
	"share"
	"testing"
)

func TestBinaryReadWriteDelete(t *testing.T) {
	bf := NewBinFile("../../dataset/bin/test")
	bf.WritePair(0, 1)
	bf.WritePair(2, 3)

	bf = NewBinFile("../../dataset/bin/test")
	arr, err := bf.ReadAll()
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
