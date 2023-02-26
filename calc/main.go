package main

import (
	"calc/calc/parser"
	"fmt"
	"log"
	"strings"
)

func ParseNumber(str string, begin int) (string, int) {
	i := 0
	res := ""
	metDot := false
	for strings.Contains("123456789.", string(str[begin])) {
		token := string(str[begin])
		if token == "." {
			if metDot {
				panic("error")
			}
			if res == "" {
				res += "0"
			}
			metDot = true
		}
		res += token
		begin++
		i++
	}
	return res, i
}
func main() {
	var expression string
	_, err := fmt.Scan(&expression)
	if err != nil {
		log.Fatalln(err)
		return
	}
	pars := parser.NewPostfixParser(expression)
	err = pars.Parse()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("result: ", pars.Calculate())

}
