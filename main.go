package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/AliiAhmadi/email_validator/validator"
)

var (
	numberOfCores = 16
	numberOfLines = 0
)

func main() {
	// Get input and output files with -i and -o flag
	inp, out := args()

	if err := run(inp, out); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, "Done!")
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

	c, err := lines(*input)
	if err != nil {
		return err
	}

	numberOfLines = c

	outf, err := os.Create(*output)
	if err != nil {
		return err
	}
	defer outf.Close()

	scanner := bufio.NewScanner(inpf)

	ch := make(chan bool, numberOfCores-1)
	var wg sync.WaitGroup
	wg.Add(numberOfLines)

	for i := 0; i < numberOfLines; i++ {
		go func(sc *bufio.Scanner) {
			defer func() {
				<-ch
				wg.Done()
			}()

			sc.Scan()
			email := sc.Text()
			v := validator.New()
			v.Email(email)

			err := v.Valid()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			ok := v.Status()

			if ok {
				_, err = outf.WriteString(email + "\n")
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
			}
		}(scanner)

		ch <- true
	}

	wg.Wait()
	close(ch)

	return nil
}

func lines(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	c := 0

	for sc.Scan() {
		c++
	}

	return c, nil
}
