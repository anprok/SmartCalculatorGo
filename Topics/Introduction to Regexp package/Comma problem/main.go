package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// DO NOT modify the code block below:
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	// Create the regexps and write the rest of the required code below:
	reDot := regexp.MustCompile(`  *\.`)         // whitespace + dot
	reComma := regexp.MustCompile(`  *\,`)       // whitespace + comma
	reExclamation := regexp.MustCompile(`  *\!`) // whitespace + exclamation

	onDot := reDot.MatchString(line)
	onComma := reComma.MatchString(line)
	onExclamation := reExclamation.MatchString(line)

	if onDot || onComma || onExclamation {
		fmt.Print("error")
	} else {
		fmt.Print("ok")
	}
}
