package main

import (
	"capudo/parser"
	"fmt"
)

func main() {
	
	path := string("test_trips_2021.csv")

	p, _ := parser.CreateParserGenerico(path)
	if p != nil {
		fmt.Println(p)
		//parser.Print(p)
	}

}
