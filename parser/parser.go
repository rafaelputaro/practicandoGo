package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type parser struct {
	rows [][]string
}

func Parser(pathArchivo string) (p *parser, err error) {
	p = new(parser)
	archivo, err := os.Open(pathArchivo)
	fileScanner := bufio.NewScanner(archivo)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		p.rows = append(p.rows, strings.Split(fileScanner.Text(), ","))
	}
	defer archivo.Close()
	return p, err
}

func Print(p *parser){
	for _, row := range p.rows{
		fmt.Println("Fila:", row)
	}
}
