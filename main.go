package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode/utf8"

	"github.com/akamensky/argparse"
	"github.com/shomali11/util/xstrings"
)

func isValid(gtin *string) bool {

	found, err := regexp.MatchString("^[0-9]{8}$|^[0-9]{12,14}$", *gtin)

	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	if !found {
		return false
	}

	if found {
		var lastChar string = (*gtin)[len(*gtin)-1:]
		var reversed string = xstrings.Reverse(*gtin)[1:]

		sum := 0
		for i, c := range reversed {
			buf := make([]byte, 1)
			_ = utf8.EncodeRune(buf, c)
			value, _ := strconv.Atoi(string(buf))
			var value2 int
			if i%2 == 0 {
				value2 = value * 3
			} else {
				value2 = value * 1
			}
			sum += value2
		}

		sum = sum % 10

		if sum == 0 {
			sum = 10
		}
		return lastChar == strconv.Itoa(10-sum)
	}

	return false
}

func main() {
	var parser *argparse.Parser = argparse.NewParser("print", "Prints provided string to stdout")
	var verbose *bool = parser.Flag("v", "verbose", &argparse.Options{Required: false, Help: "Print if it is a valid GTIN"})
	var quiet *bool = parser.Flag("q", "quiet", &argparse.Options{Required: false, Help: "Suppress all logging. Overrides verbose behaviour."})
	var gtin *string = parser.StringPositional(nil)

	var err error = parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(2)
	}
	isValidGtin := isValid(gtin)

	if *verbose && !*quiet {
		if isValidGtin {
			fmt.Printf("%s is a valid GTIN\n", *gtin)
			os.Exit(0)
		} else {
			fmt.Printf("%s is an invalid GTIN\n", *gtin)
			os.Exit(1)
		}
	}
	if isValidGtin {
		os.Exit(0)
	} else {
		os.Exit(1)
	}

}
