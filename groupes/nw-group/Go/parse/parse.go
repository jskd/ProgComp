package main

import ("fmt"
		 "io/ioutil"
		 "strings")

func main() {
	file, err := ioutil.ReadFile("data.csv")
	if err != nil {
		 panic(err)
	}
   lines := strings.Split(string(file), "\n")
	fmt.Printf("%q\n", lines)
	for i:=0;i<len(lines);i++ {
		fmt.Printf("%q\n",strings.Split(lines[i], ","))
		//car:=strings.Split(lines[i], ",")
	}
}
