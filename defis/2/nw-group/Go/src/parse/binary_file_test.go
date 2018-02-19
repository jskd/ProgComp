package parse

import (
	"share"
	"testing"
)

func TestReadWrite(t *testing.T) {
	bf := NewBinFile("../../dataset/bin_test")
	bf.Write(0, 1)
	bf.Write(2, 3)
	bf.Close()

	bf = NewBinFile("../../dataset/bin_test")
	arr, err := bf.ReadAll()
	bf.Close()
	exp := []uint32{}
	exp = append(exp, 0)
	exp = append(exp, 1)
	exp = append(exp, 2)
	exp = append(exp, 3)
	share.AssertEqual(t, err, nil, "Some error happened while reading file")
	share.AssertEqual(t, arr, exp, "Incorrect uint32 value")
}
