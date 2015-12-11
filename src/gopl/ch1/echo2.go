package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for idx, arg := range os.Args[1:] {
		s += sep + strconv.Itoa(idx) + "->" + arg
		sep = " "
	}
	fmt.Println(s)
}
