package spreadsheet

import (
	"share"
	"testing"
)

func TestToFormula(t *testing.T) {
	f := ToFormula("=#(0,0,50,50,1)")
	share.AssertEqual(t, f.xDestination, 50, "xDestination is not 50 in the Formula instance.")
}

func TestFromFile(t *testing.T) {
	bin_repo := FromFile("../../dataset/data2.csv", ',')
	exp := share.TempDir() + "data2.csv/bin"
	share.AssertEqual(t, bin_repo, exp, "Incorrect repository path")
}

func Test10LinesBigmamaFile(t *testing.T) {
	bin_repo := FromFile("../../dataset/bigmama_10.csv", ';')
	exp := share.TempDir() + "bigmama_10.csv/bin"
	share.AssertEqual(t, bin_repo, exp, "Incorrect repository path")
}
