package main

import ("fmt"
		 "io/ioutil"
		 "os"
		 "strings")

/*strucure of formula*/
type Formula struct{
	xSource int;
	ySource int;
	xDestination int;
	yDestination int;
	value int;
}

/*data may be value or formula to evaluate*/
type Data struct {
	valeur string;
	form formula;
}

/*location of cell and content*/
type Cell struct {
   x int;
   y int;
   value data;
}

func checkError(err error){
	if err != nil {
		 panic(err)
	}
}

func toFormula(formuleLue string) formula{
	formule := new(Formula)
	if(formuleLue[0]="=" && formuleLue[1]="#")
	
	formule.xSource=formuleLue[3]
	formule.ySource=formuleLue[5]
	formule.xDestination=formuleLue[7]
	formule.yDestination=formuleLue[9]
	formule.value = formuleLue[11]
}

/*read File and returns there content in a 2D slice of strings*/
func readFile(fileToRead string, sep string) [][]string{
	file, err := ioutil.ReadFile(fileToRead)
	checkError(err)
   lines := strings.Split(string(file), "\n")
	car:=make([][]string, 0, 0)
	for i:=0;i<len(lines);i++ {
		if(lines[i]!=""){
			car=append(car,strings.Split(lines[i], sep))
		}
	}
	
	return car
}

/* Evaluate content of data.csv TODO*/
func evaluate(element formula) string{
	
	return element
}

/*write in File the 2D slice of strings*/
func writeFile(){
	taille := len(readFile("data.csv",","))
	s:=make([][]string, taille,taille)
	s=readFile("data.csv",",")
	fmt.Printf("%q\n",s)
	f, err := os.OpenFile("view0.csv", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	sep := "\n"
	for i:=0;i<len(s);i++ {
		for j:=0;j<len(s[i]);j++ {
			_, err = f.WriteString(evaluate(s[i][j])+",")
			checkError(err)
		}
		_, err = f.WriteString(sep)
		checkError(err)
	}
}

/* user actions */
func userActions(){
	s:=make([][]string, 0,0)
	s=readFile("user.txt"," ")
	fmt.Println("%q",s)
}

func main() {
	writeFile()
	userActions()
}
