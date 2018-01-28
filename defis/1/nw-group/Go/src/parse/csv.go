package parse

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/*strucure of formula*/
type Formula struct {
	xSource      int
	ySource      int
	xDestination int
	yDestination int
	value        int
}

/*data may be value or formula to evaluate*/
type Data struct {
	valeur string
	form   Formula
}

/*location of cell and content*/
type Cell struct {
	x     int
	y     int
	value Data
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

func ToFormula(formuleLue string) Formula {
	formule := new(Formula)
	trimedFormule := strings.TrimPrefix(formuleLue, "=#(")
	if trimedFormule != formuleLue {
		arg_list := strings.Split(trimedFormule, ",")
		arg_list[4] = strings.Replace(arg_list[4], ")", "", -1)
		formule.xSource = convertString2Int(arg_list[0])
		formule.ySource = convertString2Int(arg_list[1])
		formule.xDestination = convertString2Int(arg_list[2])
		formule.yDestination = convertString2Int(arg_list[3])
		formule.value = convertString2Int(arg_list[4])
	}
	return *formule
}

/*read File and returns there content in a 2D slice of strings*/
func ReadCsv(fileToRead string, sep string) [][]string {
	file, err := ioutil.ReadFile(fileToRead)
	checkError(err)
	fmt.Println("hellooooo")
	lines := strings.Split(string(file), "\n")
	car := make([][]string, 0, 0)
	for i := 0; i < len(lines); i++ {
		if lines[i] != "" {
			car = append(car, strings.Split(lines[i], sep))
		}
	}
	return car
}

/* Evaluates a formula.  On success, a formula evaluates into an
   integer.  Returns an error when the dependency graph of the formula
   has a cycle.  */
func evaluate(formula Formula, spreadSheet [][]string) int {
	count := 0
	for r := formula.ySource; r <= formula.yDestination; r++ {
		for c := formula.xSource; c <= formula.xDestination; c++ {
			if spreadSheet[r][c] == string(formula.value) {
				count++
			}
		}
	}
	return count
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
			formula := ToFormula(s[i][j])
			_, err = f.WriteString(convertInt2String(evaluate(formula, s)) + sep)
			checkError(err)
		}
		_, err = f.WriteString("/n")
		checkError(err)
	}
}

/* user actions */
func userActions() {
	s := make([][]string, 0, 0)
	s = ReadCsv("user.txt", " ")
	fmt.Println("%q", s)
	WriteFile("changes.txt", " ", s)
}