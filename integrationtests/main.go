package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/a-h/round"
)

var in = flag.String("in", "python3.txt", "The name of the file to use.")
var algorithm = flag.String("round", "ToEven", "The rounding method to use, ToEven or AwayFromZero.")
var places = flag.Int("places", 2, "The number of decimal places to round to.")

func main() {
	flag.Parse()

	var rounder func(float64, int) float64
	alg := "ToEven"

	if strings.EqualFold(*algorithm, "ToEven") {
		rounder = round.ToEven
		fmt.Println("Rounding to even.")
	} else {
		rounder = round.AwayFromZero
		fmt.Println("Rounding away from zero.")
		alg = "AwayFromZero"
	}

	var err error

	file, err := os.Open(*in)
	if err != nil {
		fmt.Println(fmt.Errorf("Failed to read the input file with err %v", err))
		os.Exit(-1)
	}

	reader := bufio.NewReader(file)

	testsRan := 0

	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		parts := strings.Split(strings.TrimSpace(line), " ")
		input, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Failed to parse input float string.")
			break
		}

		expected, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			fmt.Printf("Failed to parse expected float string %v with err %v.\n", expected, err)
			break
		}

		actual := rounder(input, *places)

		if actual != expected {
			fmt.Println(strings.TrimSpace(line))
			fmt.Printf("Input %s, parsed as %v. Rounded to %d dp with %s, expected %v, but got %v\n", parts[0], input, *places, alg, expected, actual)
			return
		}

		testsRan++

		if testsRan%100 == 0 {
			fmt.Printf("%d tests executed\n", testsRan)
		}
	}

	if err != io.EOF {
		fmt.Printf("Failed!: %v\n", err)
	}
}
