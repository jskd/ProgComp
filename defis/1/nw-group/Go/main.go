package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"parse"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		main_program()
	} else {
		unit_test()
	}
}

func test_ocaml_call() {
	fmt.Println("Testing OCaml Call...")
	cmd := exec.Command("ocaml", "../OCaml/test.ml")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", out.String())
}

func unit_test() {
	test_ocaml_call()
	parse.ToFormula("=#(0,0,50,50,1)")
	parse.WriteFile()
}

func main_program() {
	fmt.Println("TODO: Main program")
}
