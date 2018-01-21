package main

import ("fmt"
		 "io/ioutil"
		 "os"
		 "strings")

func checkError(err error){
	if err != nil {
		 panic(err)
	}
}

/*read File and returns there content in a 2D slice of strings*/
func readFile() [][]string{
	file, err := ioutil.ReadFile("data.csv")
	checkError(err)
   lines := strings.Split(string(file), "\n")
	car:=make([][]string, 0, 0)
	for i:=0;i<len(lines);i++ {
		car=append(car,strings.Split(lines[i], ","))
	}
	return car
}
/*write in File the 2D slice of strings*/
func writeFile(){
	taille := len(readFile())
	s:=make([][]string, taille,cap(readFile()))
	s=readFile()
	fmt.Printf("%q\n",s)
	f, err := os.OpenFile("view0.csv", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	sep := "\n"
	for i:=0;i<len(s);i++ {
		for j:=0;j<len(s[i]);j++ {
			fmt.Println("---%d---",j)
			_, err = f.WriteString(s[i][j]+",")
			checkError(err)
		}
		_, err = f.WriteString(sep)
		checkError(err)
	}
}

func main() {
	writeFile()
}
