package main

import (
	"capudo/parser"
)

func main() {
	path := string("trips_2021.csv")
	p, _ := parser.Parser(path)
	if p != nil {
		parser.Print(p)
	}
}
