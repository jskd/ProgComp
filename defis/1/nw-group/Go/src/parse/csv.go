package parse

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func convertString2Int(str string) int {
	i, err := strconv.Atoi(str)
	checkError(err)
	return i
}

func convertInt2String(i int) string {
	s := strconv.Itoa(i)
	return s
}

/*read File and returns there content in a 2D slice of strings*/
func ReadCsv(fileToRead string, sep rune) [][]string {
	file, err := os.Open(fileToRead)
	checkError(err)
	csvReader := csv.NewReader(file)
	csvReader.Comma = sep
	csvReader.FieldsPerRecord = -1
	csvReader.TrimLeadingSpace = true
	res, err := csvReader.ReadAll()
	checkError(err)
	return res
}

/*write in File the 2D slice of strings*/
func WriteFile(fileToWrite string, sep string, s [][]string) {
	fmt.Printf("%q\n", s)
	f, err := os.OpenFile(fileToWrite, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			// formula, err := spreadsheet.ToData(s[i][j])
			// _, err = f.WriteString(convertInt2String(spreadsheet.Evaluate(formula, s)) + sep)
			// checkError(err)
		}
		_, err = f.WriteString("\n")
		checkError(err)
	}
}

func WriteCsv(filename string, data [][]string, sep rune) {
	file, err := os.Create(filename)
	checkError(err)
	csvReader := csv.NewWriter(file)
	csvReader.Comma = sep
	err = csvReader.WriteAll(data)
	checkError(err)
}
