package parse

import (
	"share"
	"testing"
)

func TestToFormula(t *testing.T) {
	f := ToFormula("=#(0,0,50,50,1)")
	share.AssertEqual(t, f.xDestination, 50, "xDestination is not 50 in the Formula instance.")
}

func TestReadCsvFile(t *testing.T) {
	s := ReadCsv("../../dataset/data.csv", ",")
	WriteFile("view0.csv", ";", s)
}

func TestReadUserTxt(t *testing.T) {
	s := ReadCsv("../../dataset/user.txt", " ")
	WriteFile("changes.txt", " ", s)
}
