package parser

import (
	"fmt"
	"strconv"
	"strings"

	"calc/calc/utils"
)

type PostfixParser struct {
	infixExpression   string
	postfixExpression string
	operationStack    utils.Stack[string]
	valuesStack       utils.Stack[float64]
}

func NewPostfixParser(infixExpression string) *PostfixParser {
	return &PostfixParser{infixExpression: infixExpression}
}

func (p *PostfixParser) GetInfixExpression() string {
	return p.infixExpression
}

func (p *PostfixParser) SetInfixExpression(expr string) {
	p.infixExpression = expr
}

func isOperator(value string) bool {
	operators := "+-*/"
	return strings.Contains(operators, value)
}

func getOperatorPriority(value string) int {
	switch {
	case strings.Contains("+-", value):
		return 1
	case strings.Contains("*/", value):
		return 2
	default:
		return -1
	}
}

func (p *PostfixParser) Parse() (err error) {
	p.infixExpression = strings.ReplaceAll(p.infixExpression, " ", "")
	p.infixExpression = strings.ReplaceAll(p.infixExpression, ",", ".")
	last := " "
	for i := 0; i < len(p.infixExpression); i++ {
		currentToken := string(p.infixExpression[i])
		switch {
		case utils.IsPartOfNumber(currentToken):
			{
				res, parsedRunes, err := utils.ParseNumber(p.infixExpression, i)
				if err != nil {
					return err
				}
				p.postfixExpression += res + " "
				i += parsedRunes
				last = currentToken
			}
		case utils.IsOpeningBracket(currentToken):
			{
				p.operationStack.PushBack(currentToken)
				last = currentToken
			}
		case utils.IsClosingBracket(currentToken):
			{
				if !p.operationStack.Has("(") {
					return fmt.Errorf("ошибка при парсинге выражения,"+
						"не удалось найти открывающую скобоку для закрывающей на позиции %d", i)
				}
				for p.operationStack.Top() != "(" {
					p.postfixExpression += p.operationStack.Top()
					p.operationStack.Pop()
				}
				p.operationStack.Pop()
				last = currentToken
			}
		case isOperator(currentToken):
			if isOperator(last) && strings.Contains("+~*/", currentToken) || currentToken == "-" && last == "~" {
				return fmt.Errorf("unexpected token: %s in %d position", currentToken, i)
			}
			// учитываем унарный минус
			if currentToken == "-" && strings.Contains("(+-*/ ", last) {
				p.postfixExpression += "0 "
				p.operationStack.PushBack(currentToken)
				last = "~"
				continue
			}
			for !p.operationStack.Empty() &&
				getOperatorPriority(currentToken) <=
					getOperatorPriority(p.operationStack.Top()) {
				p.postfixExpression += p.operationStack.Top()
				p.operationStack.Pop()
			}
			p.operationStack.PushBack(currentToken)
			last = currentToken
		default:
			return fmt.Errorf("unexpected token: %s", currentToken)
		}
	}
	for !p.operationStack.Empty() {
		p.postfixExpression += p.operationStack.Top()
		p.operationStack.Pop()
	}
	return err
}

func (p *PostfixParser) Calculate() (float64, error) {
	for i := 0; i < len(p.postfixExpression); i++ {
		currentToken := string(p.postfixExpression[i])
		switch {
		case utils.IsPartOfNumber(currentToken):
			{
				res, parsedRunes, err := utils.ParseNumber(p.postfixExpression, i)
				if err != nil {
					return 0, err
				}
				number, err := strconv.ParseFloat(res, 64)
				if err != nil {
					return 0, err
				}
				p.valuesStack.PushBack(number)
				i += parsedRunes
			}
		case isOperator(currentToken):
			{
				second := p.valuesStack.GetTopOrDefault()
				p.valuesStack.Pop()
				first := p.valuesStack.GetTopOrDefault()
				p.valuesStack.Pop()
				p.valuesStack.PushBack(executeOperation(currentToken, first, second))
			}
		}
	}
	return p.valuesStack.GetTopOrDefault(), nil
}

func executeOperation(operator string, lhs, rhs float64) float64 {
	switch operator {
	case "+":
		return lhs + rhs
	case "-":
		return lhs - rhs
	case "*":
		return lhs * rhs
	case "/":
		return lhs / rhs
	}
	return 0
}

func (p *PostfixParser) GetParsedExpression() string {
	return p.postfixExpression
}

func (p *PostfixParser) ParseAndCalculate(expression string) (float64, error) {
	p.SetInfixExpression(expression)
	err := p.Parse()
	if err != nil {
		return 0, err
	}
	return p.Calculate()
}
