package main

import (
	"flag"
	"fmt"
)

func main() {
	inp, out := args()
	fmt.Println(*inp, *out)
}

func args() (*string, *string) {
	inp := flag.String("i", "input.txt", "The name of input file that contains the emails you want to validate")
	out := flag.String("o", "output.txt", "The name of output file that want to see result")

	flag.Parse()
	return inp, out
}
