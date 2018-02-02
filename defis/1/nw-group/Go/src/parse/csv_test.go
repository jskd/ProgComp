package parse

import (
	"testing"
)

func TestReadCsvFile(t *testing.T) {
	s := ReadCsv("../../dataset/data.csv", ',')
	WriteFile("view0.csv", ";", s)
}

func TestReadUserTxt(t *testing.T) {
	s := ReadCsv("../../dataset/user.txt", ' ')
	WriteFile("changes.txt", " ", s)
}
