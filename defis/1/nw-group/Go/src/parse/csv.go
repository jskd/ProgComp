package parse

import (
	"encoding/csv"
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

func WriteCsv(filename string, data [][]string, sep rune) {
	file, err := os.Create(filename)
	checkError(err)
	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = sep
	err = csvWriter.WriteAll(data)
	checkError(err)
	//TODO: Close file?
}
