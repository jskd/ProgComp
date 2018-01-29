package parse

import (
	"fmt"
	"testing"
)

func TestToFormula(t *testing.T) {
	fmt.Println("Testing ToFormula Call...")
	ToFormula("=#(0,0,50,50,1)")
}

func TestReadCsvFile(t *testing.T) {
	s := parse.ReadCsv("data.csv", ",")
	WriteFile("view0.csv", ";", s)
}

func TestReadUserTxt() {
	s := parse.ReadCsv("user.txt", " ")
	fmt.Println("%q", s)
	WriteFile("changes.txt", " ", s)
}
