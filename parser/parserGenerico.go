package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type parserGenerico struct {
	rows [][]string
}

func CreateParserGenerico(pathArchivo string) (p *parserGenerico, err error) {
	p = new(parserGenerico)
	archivo, err := os.Open(pathArchivo)
	fileScanner := bufio.NewScanner(archivo)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		p.rows = append(p.rows, strings.Split(fileScanner.Text(), ","))
	}
	defer archivo.Close()
	return p, err
}

func Print(p *parserGenerico){
	for _, row := range p.rows{
		fmt.Println("Fila:", row)
	}
}
