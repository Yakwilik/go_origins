package main

import (
	"calc/calc/parser"
	"calc/calc/utils"
	"fmt"
	"log"
)

func main() {
	expression, err := utils.GetInput()

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
