package spreadsheet

import (
	"fmt"
	"parse"
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
	exp_repo := share.TempDir() + "bigmama_10.csv/bin"
	share.AssertEqual(t, bin_repo, exp_repo, "Incorrect repository path")

	bfOut := parse.NewBinFile(bin_repo + "/3")
	out, errOut := bfOut.ReadAll()
	bfOut.Close()
	bfExpect := parse.NewBinFile("../../dataset/bin/3")
	exp, errExp := bfExpect.ReadAll()
	bfExpect.Close()

	share.AssertEqual(t, errOut, nil, "Something wrong reading bin "+bin_repo+"/3")
	share.AssertEqual(t, errExp, nil, "Something wrong reading bin ../../dataset/bin/3")
	share.AssertEqual(t, len(out), 985*2, "Incorrect length in bin value of '3'") //There are 985 value-'3' in bigmama_10.csv
	share.AssertEqual(t, out, exp, "Incorrect values in bin value '3'")
	parse.BinFileManager().SaveAndCloseAll()
}

func TestEvaluate(t *testing.T) {
	out := Evaluate("../../dataset/bin/")
	share.AssertEqual(t, out, 1, "")
}

func TestCounting(t *testing.T) {
	out := EvaluateFormula("../../dataset/bin/FORMULAS/", "1110_7572_3186_17282_3")
	fmt.Printf("%d\n", out)
}
