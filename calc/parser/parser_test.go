package parser

import (
	"testing"
)

func TestNewPostfixParser(t *testing.T) {
	parser := NewPostfixParser("6+3")
	if parser == nil {
		t.Fatalf("parser must not be nil")
	}
}

func TestPostfixParser_Calculate(t *testing.T) {
	parser := NewPostfixParser("6+3")

	if parser.Calculate() != 0 {
		t.Errorf("parser must return 0, if Parse function haven`t been called")
	}
}

func TestPostfixParser_GetInfixExpression(t *testing.T) {
	parser := NewPostfixParser("6+3")
	if parser.GetInfixExpression() != "6+3" {
		t.Fatalf("wrong storing")
	}
}

func TestPostfixParser_GetParsedExpression(t *testing.T) {
	parser := NewPostfixParser("6+3")

	if parser.GetParsedExpression() != "" {
		t.Errorf("parser must return empty string, if Parse function haven`t been called")
	}
}

func TestPostfixParser_SetInfixExpression(t *testing.T) {
	parser := NewPostfixParser("6+3")
	parser.SetInfixExpression("5+5")
	if parser.GetInfixExpression() != "5+5" {
		t.Errorf("Parser must change infix expression after setting a new one")
	}
}

func TestPostfixParser_Parse(t *testing.T) {
	parser := NewPostfixParser("6+3")
	_ = parser.Parse()
	ans := parser.Calculate()
	if ans != 9 {
		t.Errorf("wrong answer")
	}
	parser.SetInfixExpression(")5+3")
	err := parser.Parse()
	if err == nil {
		t.Errorf("Parsing must have ended with error")
	}

	parser.SetInfixExpression("..5+3")
	err = parser.Parse()
	if err == nil {
		t.Errorf("Parsing must have ended with error")
	}

	parser.SetInfixExpression("allo")
	err = parser.Parse()
	if err == nil {
		t.Errorf("Parsing must have ended with error")
	}

	parser.SetInfixExpression("5*(3+6)")
	_ = parser.Parse()
	ans = parser.Calculate()
	if ans != 45 {
		t.Errorf("wrong answer")
	}
	parser.SetInfixExpression("-5*(3+6)")
	_ = parser.Parse()
	ans = parser.Calculate()
	if ans != -45 {
		t.Errorf("wrong answer")
	}
	parser.SetInfixExpression("5*(3+6)/2")
	_ = parser.Parse()
	ans = parser.Calculate()
	if ans != 45/2. {
		t.Errorf("wrong answer")
	}
}
func TestExecuteOperation(t *testing.T) {
	res := executeOperation("q", 1, 1)
	if res != 0 {
		t.Errorf("invalid operator must return 0")
	}

}
