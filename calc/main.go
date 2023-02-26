package main

import (
	"fmt"
	"log"

	"calc/calc/parser"
)

func main() {
	var expression string
	_, err := fmt.Scanln(&expression)
	if err != nil {
		log.Fatalln(err)
	}

	pars := parser.NewPostfixParser(expression)
	err = pars.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("result: ", pars.Calculate())
}
