package parse

import (
	"io"
	"share"
	"testing"
)

func TestCsvParser(t *testing.T) {
	cp := NewCsvParser("../../dataset/data.csv", ';', '"')
	expPos := []uint32{}
	expPos = append(expPos, 0)
	expPos = append(expPos, 1)
	expPos = append(expPos, 2)
	expPos = append(expPos, 3)
	expPos = append(expPos, 10)
	expPos = append(expPos, 11)
	expPos = append(expPos, 12)
	expPos = append(expPos, 13)
	outPos := []uint32{}
	exp := []string{}
	exp = append(exp, "1")
	exp = append(exp, " 0")
	exp = append(exp, " 1")
	exp = append(exp, "=#(0, 0, 0, 2, 1)")
	exp = append(exp, "0")
	exp = append(exp, " 1")
	exp = append(exp, " 1")
	exp = append(exp, "=#(0, 0, 1, 2, 1)")
	out := []string{}
	for {
		str, x, y, err := cp.ReadOneCell()
		if err == io.EOF {
			break
		}
		outPos = append(outPos, x*10+y)
		out = append(out, str)
	}
	share.AssertEqual(t, outPos, expPos, "Incorrect index output from reading data.csv")
	share.AssertEqual(t, out, exp, "Incorrect text output from reading data.csv")
}

func TestCsvParserWithQuote(t *testing.T) {
	cp2 := NewCsvParser("../../dataset/data2.csv", ',', '"')
	expPos := []uint32{}
	expPos = append(expPos, 0)
	expPos = append(expPos, 1)
	expPos = append(expPos, 2)
	expPos = append(expPos, 10)
	expPos = append(expPos, 11)
	expPos = append(expPos, 12)
	outPos := []uint32{}
	exp := []string{}
	exp = append(exp, "1")
	exp = append(exp, " 0")
	exp = append(exp, "=#(0, 0, 1, 1, 1)")
	exp = append(exp, "=#(1, 0, 1, 0, 0)")
	exp = append(exp, " 1")
	exp = append(exp, " 0")
	out := []string{}
	for {
		str, x, y, err := cp2.ReadOneCell()
		if err == io.EOF {
			break
		}
		//fmt.Printf("%d,%d: '%s'\n", x, y, str)
		outPos = append(outPos, x*10+y)
		out = append(out, str)
	}
	share.AssertEqual(t, outPos, expPos, "Incorrect index output from reading data2.csv")
	share.AssertEqual(t, out, exp, "Incorrect text output from reading data2.csv")
}

func TestReadOneLineCsv(t *testing.T) {
	in := "first_name;last_name;username\n#lines beginning with a # character are ignored\nKen;Thompson;ken\n\"Robert\";\"Griesemer\";\"gri\""
	out := ReadOneLineCsv(in, ';')
	exp := [][]string{}
	exp0 := []string{"first_name", "last_name", "username"}
	exp1 := []string{"Ken", "Thompson", "ken"}
	exp2 := []string{"Robert", "Griesemer", "gri"}
	exp = append(exp, exp0)
	exp = append(exp, exp1)
	exp = append(exp, exp2)
	share.AssertArrayEqual(t, out, exp, "Incorrect output from reading data3.csv")
}

func TestReadWriteCsvFile(t *testing.T) {
	s := ReadCsv("../../dataset/data.csv", ';')
	WriteCsv("../../dataset/view.csv", s, ',')
	exp := ReadCsv("../../dataset/view.csv", ',')
	share.AssertArrayEqual(t, s, exp, "Incorrect write file to view.csv")
}

func TestReadSimpleCsvFile(t *testing.T) {
	s := ReadCsv("../../dataset/data3.csv", ';')
	exp := [][]string{}
	exp0 := []string{"1", "1", "1", "0", "0", "0"}
	exp1 := []string{"2", "2", "2", "2", "2", "2"}
	exp2 := []string{"0", "0", "0", "3", "3", "3"}
	exp3 := []string{"1", "0", "-1", "0", "1", "-1"}
	exp = append(exp, exp0)
	exp = append(exp, exp1)
	exp = append(exp, exp2)
	exp = append(exp, exp3)
	share.AssertArrayEqual(t, s, exp, "Incorrect output from reading data3.csv")
}

func TestReadCsvFileWithFormula(t *testing.T) {
	s := ReadCsv("../../dataset/data2.csv", ',')
	exp := [][]string{}
	exp0 := []string{"1", "0", "=#(0, 0, 1, 1, 1)"}
	exp1 := []string{"=#(1, 0, 1, 0, 0)", "1", "0"}
	exp = append(exp, exp0)
	exp = append(exp, exp1)
	share.AssertArrayEqual(t, s, exp, "Incorrect output from reading data2.csv")
}
