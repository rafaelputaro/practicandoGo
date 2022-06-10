package parser

import (
	"fmt"
	"os"
	"encoding/csv"
)

type parserGenerico struct {
	rows [][]string
}

func CreateParserGenerico(pathArchivo string) (p *parserGenerico, err error) {
	p = new(parserGenerico)
	archivo, err := os.Open(pathArchivo)
	if( err == nil){
		defer archivo.Close()
		csvReader := csv.NewReader(archivo)
		p.rows, err = csvReader.ReadAll()	
	}
	return p, err
}

func Print(p *parserGenerico){
	for _, row := range p.rows{
		fmt.Println("Fila:", row)
	}
}
