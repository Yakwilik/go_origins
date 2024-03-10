package utils

import (
	"testing"
)

func TestIsPartOfNumber(t *testing.T) {
	strs := []string{".", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for _, str := range strs {
		if !IsPartOfNumber(str) {
			t.Errorf("Ошибка, %s – часть числа", str)
		}
	}
	strs = []string{",", "q", "w", "e", "r", "t", "u", "v", "w", "x", "w"}
	for _, str := range strs {
		if IsPartOfNumber(str) {
			t.Errorf("Ошибка, %s – не часть числа", str)
		}
	}
}

func TestIsClosingBracket(t *testing.T) {
	if !IsClosingBracket(")") {
		t.Fatalf("Ошибка, это закрывающая скобка")
	}
	if IsClosingBracket("(") {
		t.Fatalf("Ошибка, это открывающая скобка")
	}
}

func TestIsOpeningBracketBracket(t *testing.T) {
	if IsOpeningBracket(")") {
		t.Fatalf("Ошибка, это закрывающая скобка")
	}
	if !IsOpeningBracket("(") {
		t.Fatalf("Ошибка, это открывающая скобка")
	}
}

func TestStack_Empty(t *testing.T) {
	stack := Stack[int]{}
	if !stack.Empty() {
		t.Errorf("stack is empty!")
	}
	stack.PushBack(1)

	if stack.Empty() {
		t.Errorf("stack is not empty!")
	}
}

func TestStack_Has(t *testing.T) {
	stack := Stack[int]{}

	if stack.Has(1) {
		t.Errorf("stack is empty!")
	}
	stack.PushBack(1)

	if !stack.Has(1) {
		t.Errorf("stack has 1!")
	}
}

func TestStack_Pop(t *testing.T) {
	stack := Stack[int]{}
	stack.Pop()
	if !stack.Empty() {
		t.Fatalf("stack is empty")
	}
	stack.PushBack(1)
	stack.Pop()
	if !stack.Empty() {
		t.Fatalf("stack is empty")
	}
}

func TestStack_Size(t *testing.T) {
	stack := Stack[int]{}
	if !(stack.Size() == 0) {
		t.Errorf("stack size is 0")
	}
	stack.PushBack(1)
	if !(stack.Size() == 1) {
		t.Errorf("stack size is 1")
	}
	stack.PushBack(2)
	if !(stack.Size() == 2) {
		t.Errorf("stack size is 2")
	}
}

func TestStack_GetTopOrDefault(t *testing.T) {
	stack := Stack[string]{}

	def := stack.GetTopOrDefault()

	if def != "" {
		t.Fatalf("empty stack must return default value")
	}
	stack.PushBack("hello")

	def = stack.GetTopOrDefault()

	if def != "hello" {
		t.Fatalf("not empty stack must return true value")
	}
}

func TestStack_Top(t *testing.T) {
	stack := Stack[string]{}
	stack.PushBack("hello")
	if stack.Top() != "hello" {
		t.Fatalf("wrong value from top of stack")
	}
}

func TestStack_PushBack(t *testing.T) {
	stack := Stack[string]{}
	stack.PushBack("hello")
	if stack.Top() != "hello" {
		t.Fatalf("wrong value from top of stack")
	}
	stack.PushBack("goodbye")
	if stack.Top() != "goodbye" {
		t.Fatalf("wrong value from top of stack")
	}
}

func TestParseNumber(t *testing.T) {
	expression := "..525+36"
	_, _, err := ParseNumber(expression, 0)
	if err == nil {
		t.Fatalf("Parsing %s from %d index must return error", expression, 0)
	}
	res, num, err := ParseNumber(expression, 1)
	if err != nil {
		t.Fatalf("Parsing %s from %d index must not return error", expression, 1)
	}
	if res != "0.525" {
		t.Fatalf("function parsed the number wrong")
	}
	if num != 3 {
		t.Errorf("wrong number of chars in parsed number")
	}
	res, num, err = ParseNumber(expression, 6)
	if err != nil {
		t.Fatalf("Parsing %s from %d index must not return error", expression, 5)
	}
	if res != "36" {
		t.Fatalf("function parsed the number wrong")
	}
	if num != 1 {
		t.Errorf("wrong number of chars in parsed number")
	}
}
