package parse

import (
	"share"
	"testing"
)

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
