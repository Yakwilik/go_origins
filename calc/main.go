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

	pars := parser.PostfixParser{}
	result, err := pars.ParseAndCalculate(expression)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("result: ", result)
	fmt.Println("parsed: ", pars.GetParsedExpression())
}
