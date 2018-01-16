package main

import (
  "fmt"
  "os/exec"
  "log"
  "bytes"
)

func main() {
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

