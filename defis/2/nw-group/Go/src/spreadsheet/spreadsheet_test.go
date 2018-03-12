package spreadsheet

import (
	"parse"
	"share"
	"testing"
)

func TestToFormula(t *testing.T) {
	f := ToFormula("=#(0,0,50,50,1)")
	share.AssertEqual(t, f.xSource, uint32(0), "xSource is not uint32 0 from parsing formula string.")
	share.AssertEqual(t, f.ySource, uint32(0), "ySource is not uint32 0 from parsing formula string.")
	share.AssertEqual(t, f.xDestination, uint32(50), "xDestination is not uint32 50 from parsing formula string.")
	share.AssertEqual(t, f.yDestination, uint32(50), "yDestination is not uint32 50 from parsing formula string.")
	share.AssertEqual(t, f.value, uint32(1), "value is not uint32 1 from parsing formula string.")
}

func TestBinFileToFormula(t *testing.T) {
	f := ToFormula("0_0_50_50_1")
	share.AssertEqual(t, f.xSource, uint32(0), "xSource is not uint32 0 from parsing formula string.")
	share.AssertEqual(t, f.ySource, uint32(0), "ySource is not uint32 0 from parsing formula string.")
	share.AssertEqual(t, f.xDestination, uint32(50), "xDestination is not uint32 50 from parsing formula string.")
	share.AssertEqual(t, f.yDestination, uint32(50), "yDestination is not uint32 50 from parsing formula string.")
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
	out := Evaluate("../../dataset/bin/", false)
	share.AssertEqual(t, out, 2, "")
}

func TestEvaluateFormula(t *testing.T) {
	out := EvaluateFormula("../../dataset/bin/", "1110_7572_3186_17282_3", false, nil)
	share.AssertEqual(t, out, uint32(0), "Count formula 1110_7572_3186_17282_3 should be 0")
	out = EvaluateFormula("../../dataset/bin/", "236_0_236_611_3", false, nil)
	share.AssertEqual(t, out, uint32(2), "Count formula 236_0_236_611_3 should be 2")
}
