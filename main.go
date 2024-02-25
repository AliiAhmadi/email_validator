package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/AliiAhmadi/email_validator/validate"
)

func main() {
	// Get input and output files with -i and -o flag
	inp, out := args()

	if err := run(inp, out); err != nil {
		log.Fatal(err)
	}
}

func args() (*string, *string) {
	inp := flag.String("i", "input.txt", "The name of input file that contains the emails you want to validate")
	out := flag.String("o", "output.txt", "The name of output file that want to see result")

	flag.Parse()
	return inp, out
}

func run(input *string, output *string) error {
	inpf, err := os.Open(*input)
	if err != nil {
		return err
	}
	defer inpf.Close()

	outf, err := os.Create(*output)
	if err != nil {
		return err
	}
	defer outf.Close()

	scanner := bufio.NewScanner(inpf)

	for scanner.Scan() {
		email := scanner.Text()
		ok, err := validate.Validate(email)
		if err != nil {
			return err
		}

		var message string = "invalid"
		if ok {
			message = "valid"
		}

		_, err = outf.WriteString(email + " --> " + message + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
