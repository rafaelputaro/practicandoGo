package main

import (
	"capudo/parser"
)

//"capudo/parser"

func main() {
	
	path := string("trips_2021.csv")

	p, _ := parser.CreateParserGenerico(path)
	if p == nil {
		parser.Print(p)
	}

}
