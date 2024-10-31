package main

import (
	"os"
	"strings"

	lemIn "lemIn/funcs"
)

func main() {
	if len(os.Args) != 2 {
		lemIn.Error("wrong number of arguments")
		return
	}
	file := os.Args[1]
	if !strings.HasSuffix(file, ".txt") {
		lemIn.Error("not a valid file extenstion")
		return
	}
	err := lemIn.ReadData(file)
	if err != nil {
		lemIn.Error(err.Error())
		return
	}
	// lemIn.PrintAll()
	err = lemIn.Lem()
	if err != nil {
		lemIn.Error(err.Error())
		return
	}
}
