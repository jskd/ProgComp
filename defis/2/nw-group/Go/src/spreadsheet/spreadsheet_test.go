package spreadsheet

import (
	"math/rand"
	"parse"
	"share"
	"testing"
)

func TestToFormula(t *testing.T) {
	f := ToFormula("=#(0,1,30,50,1)")
	share.AssertEqual(t, f.area.src.col, uint32(0), "xSource is not uint32 0 from parsing formula string.")
	share.AssertEqual(t, f.area.src.row, uint32(1), "ySource is not uint32 0 from parsing formula string.")
	share.AssertEqual(t, f.area.dst.col, uint32(30), "xDestination is not uint32 50 from parsing formula string.")
	share.AssertEqual(t, f.area.dst.row, uint32(50), "yDestination is not uint32 50 from parsing formula string.")
	share.AssertEqual(t, f.value, uint32(1), "value is not uint32 1 from parsing formula string.")
}

func TestBinFileToFormula(t *testing.T) {
	f := ToFormula("0_1_30_50_1")
	share.AssertEqual(t, f.area.src.col, uint32(0), "xSource is not uint32 0 from parsing formula string.")
	share.AssertEqual(t, f.area.src.row, uint32(1), "ySource is not uint32 0 from parsing formula string.")
	share.AssertEqual(t, f.area.dst.col, uint32(30), "xDestination is not uint32 50 from parsing formula string.")
	share.AssertEqual(t, f.area.dst.row, uint32(50), "yDestination is not uint32 50 from parsing formula string.")
	formula_bin_path := FormulaToBinFileName("0_1_30_50_1")
	share.AssertEqual(t, formula_bin_path, "1/0_1_30_50_1", "Formula bin path is not correct.")
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

func formulas(n int) []*formula {
	src := position{0, 0}
	dst := src
	res := make([]*formula, n)
	for i := uint32(1); i <= uint32(len(res)); i++ {
		pos := position{i, i}
		res[i-1] = &formula{area: area{src, dst}, pos: pos, value: i}
		dst = pos
	}
	shuffle := func(fs []*formula) []*formula {
		perm := rand.Perm(len(fs))
		res := make([]*formula, len(fs))

		for i := 0; i < len(fs); i++ {
			res[i] = fs[perm[i]]
		}
		return res
	}
	return shuffle(res)
}

func TestDependencyGraph(t *testing.T) {
	n := 1000
	g := dependencyGraph(formulas(n))
	expectedNbEdges := func() int {
		s := 0
		for i := 0; i < n; i++ {
			s += i
		}
		return s
	}
	share.AssertEqual(t, g.NbNodes(), n, "")
	share.AssertEqual(t, g.NbEdges(), expectedNbEdges(), "")
}

func TestSplitFormulas(t *testing.T) {
	n := 1000
	fs, invalids := splitFormulas(formulas(n))
	share.AssertEqual(t, len(invalids), 0, "")
	for i, f := range fs {
		share.AssertEqual(t, f.value, uint32(i)+1, "")
	}
}
