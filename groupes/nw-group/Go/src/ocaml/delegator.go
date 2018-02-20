package ocaml

import (
	"bytes"
	"os/exec"
)

func Execute(mlFilePath string) string {
	cmd := exec.Command("ocaml", mlFilePath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return out.String()
}
