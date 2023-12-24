package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func findDigits(line string) []rune {
	l := []rune(line)
	var rv []rune
	for _, c := range l {
		if unicode.IsDigit(c) {
			rv = append(rv, c)
			break
		}
	}
	for i := len(l) - 1; i >= 0; i-- {
		c := l[i]
		if unicode.IsDigit(c) {
			rv = append(rv, c)
			break
		}
	}
	return rv
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
