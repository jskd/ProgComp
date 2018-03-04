package parse

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing/iotest"
	"unicode/utf8"
)

type CsvParser struct {
	delimiter rune
	quote     rune
	current_x uint32
	current_y uint32
	reader    bufio.Reader
}

func (r *CsvParser) ReadOneCell() (string, uint32, uint32, error) {
	delim := uint8('\n')
	str_length := 0
	x := uint32(0)
	y := uint32(0)
	isQuoteMode := false
	var str []byte
	for {
		var ru, size, err = r.reader.ReadRune()
		fmt.Printf("%s, %d, %s", string(ru), size, err)
		if err == io.EOF {
			return "", r.current_x, r.current_y, err
		}

		if err != nil && err != iotest.ErrTimeout {
			panic(err)
		}
		x = r.current_x
		y = r.current_y
		str_length += 1
		if ru == r.delimiter {
			r.current_y += 1
			break
		} else if ru == '\n' {
			r.current_x += 1
			r.current_y = 0
			break
		} else {
			AppendRune(str, ru)
		}

	}
	return string(str[:]), x, y, nil
}

func AppendRune(dest []byte, uc rune) []byte {
	var buff [utf8.UTFMax]byte
	n := utf8.EncodeRune(buff[:], uc)
	return append(dest, buff[:n]...)
}

func NewCsvParser(filename string, delimiter rune, quote rune) *CsvParser {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	return &CsvParser{delimiter, quote, uint32(0), uint32(0), reader}
}

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

func ReadOneLineCsv(line string, sep rune) [][]string {
	r := csv.NewReader(strings.NewReader(line))
	r.Comma = sep
	r.Comment = '#'
	records, err := r.ReadAll()
	checkError(err)
	return records
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

func WriteCsv(filename string, data []string, sep rune) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if os.IsNotExist(err) {
		file, err = os.Create(filename)
	}
	checkError(err)
	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = sep
	csvWriter.UseCRLF = false
	err = csvWriter.Write(data)
	b := []string{"\n"}
	err = csvWriter.Write(b)
	checkError(err)
}
