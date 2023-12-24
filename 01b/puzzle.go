package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func findDigits(line string) []rune {
	var rv []rune
	for i := range line {
		if c, ok := expandedDigit(line[i:]); ok {
			rv = append(rv, c)
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if c, ok := expandedDigit(line[i:]); ok {
			rv = append(rv, c)
			break
		}
	}
	return rv
}

func expandedDigit(fragment string) (rune, bool) {
	switch {
	case unicode.IsDigit(rune(fragment[0])):
		return rune(fragment[0]), true
	case strings.HasPrefix(fragment, "one"):
		return '1', true
	case strings.HasPrefix(fragment, "two"):
		return '2', true
	case strings.HasPrefix(fragment, "three"):
		return '3', true
	case strings.HasPrefix(fragment, "four"):
		return '4', true
	case strings.HasPrefix(fragment, "five"):
		return '5', true
	case strings.HasPrefix(fragment, "six"):
		return '6', true
	case strings.HasPrefix(fragment, "seven"):
		return '7', true
	case strings.HasPrefix(fragment, "eight"):
		return '8', true
	case strings.HasPrefix(fragment, "nine"):
		return '9', true
	}
	return 0, false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum int
	for scanner.Scan() {
		digits := string(findDigits(scanner.Text()))
		num, _ := strconv.Atoi(digits)
		fmt.Println(num)
		sum += num
	}
	fmt.Printf("----\n%d", sum)
}
