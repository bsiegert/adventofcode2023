package main

import (
	"fmt"
	"os"

	"github.com/bsiegert/adventofcode2023/03/grid"
)

func fileSum(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	g, err := grid.ReadGrid(f)
	if err != nil {
		return err
	}
	fmt.Println(grid.NumbersWithCC(g))
	fmt.Println(grid.Sum(grid.NumbersWithCC(g)))
	return nil
}

func main() {
	err := fileSum(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
