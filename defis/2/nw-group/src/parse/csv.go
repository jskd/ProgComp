package parse

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"share"
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
	reader    *bufio.Reader
}

func (r *CsvParser) ReadOneCell() (string, uint32, uint32, error) {
	str_length := 0
	x := uint32(0)
	y := uint32(0)
	isQuoteMode := false
	var str []byte
	for {
		var ru, _, err = r.reader.ReadRune()
		if err == io.EOF {
			return "", r.current_x, r.current_y, err
		}

		if err != nil && err != iotest.ErrTimeout {
			panic(err)
		}
		x = r.current_x
		y = r.current_y
		str_length += 1
		if ru == r.quote {
			isQuoteMode = !isQuoteMode
		} else if ru == r.delimiter && !isQuoteMode {
			r.current_y += 1
			break
		} else if ru == '\n' {
			r.current_x += 1
			r.current_y = 0
			break
		} else {
			str = AppendRune(str, ru)
		}
	}
	out := string(str[:])
	txt := strings.TrimSpace(out)
	return txt, x, y, nil
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

func convertString2Int(str string) int {
	i, err := strconv.Atoi(str)
	share.CheckError(err)
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
	share.CheckError(err)
	return records
}

/*read File and returns there content in a 2D slice of strings*/
func ReadCsv(fileToRead string, sep rune) [][]string {
	file, err := os.Open(fileToRead)
	share.CheckError(err)
	csvReader := csv.NewReader(file)
	csvReader.Comma = sep
	csvReader.FieldsPerRecord = -1
	csvReader.TrimLeadingSpace = true
	res, err := csvReader.ReadAll()
	share.CheckError(err)
	return res
}

func WriteCsv(filename string, data [][]string, sep rune) {
	file, err := os.Create(filename)
	share.CheckError(err)
	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = sep
	err = csvWriter.WriteAll(data)
	share.CheckError(err)
	//TODO: Close file?
}

func WriteOneLineCsv(filename string, data []string, sep rune) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if os.IsNotExist(err) {
		file, err = os.Create(filename)
	}
	share.CheckError(err)
	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = sep
	err = csvWriter.Write(data)
	share.CheckError(err)
	csvWriter.Flush()
}

func NewCsvReader(r io.Reader, sep rune) *csv.Reader {
	csvReader := csv.NewReader(r)
	csvReader.Comma = sep
	csvReader.FieldsPerRecord = -1
	csvReader.TrimLeadingSpace = true
	return csvReader
}
