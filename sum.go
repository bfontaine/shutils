package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	sum      float64
	floatSum bool

	scanner *bufio.Scanner
)

func main() {
	flag.BoolVar(&floatSum, "F", false, "compute the sum as a float")
	flag.Parse()

	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	if floatSum {
		for scanner.Scan() {
			incrFloat(scanner.Text())
		}
		// we use FormatFloat instead of %f here to print only the smallest
		// number of digits
		fmt.Println(strconv.FormatFloat(sum, 'f', -1, 64))
	} else {
		for scanner.Scan() {
			incrInt(scanner.Text())
		}
		fmt.Printf("%d\n", int64(sum))
	}
}

func incrFloat(s string) {
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		sum += f
	}
}

func incrInt(s string) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		sum += float64(i)
	}
}
