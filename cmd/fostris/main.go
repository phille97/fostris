package main

import (
	"fmt"
	"os"

	"github.com/phille97/fostris"
)

func main() {
	if os.Args[1] == "-d" {
		fmt.Println(fostris.Decode(os.Args[2]))
	} else {
		fmt.Println(fostris.Encode(os.Args[1]))
	}
}
