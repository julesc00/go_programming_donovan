// Echo2 prints its command-line arguments, inefficient way of doing it.
package main

import (
	"fmt"
	"os"
)

func main() {
	// Ways to declare a variable
	// p := ""  This is used within a function only.
	// var p string
	// var p = ""
	// var p string = ""

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println()
}
