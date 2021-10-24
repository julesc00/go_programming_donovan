// Echo1 prints its command-line arguments, more efficient way of doing this.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1:])
}
