package parser

import (
	"github.com/stretchr/testify/require"
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

func TestOperators(t *testing.T) {
	type testCase struct {
		in  string
		out float64
	}

	cases := map[string][]testCase{
		"operator plus": {
			{"6+6", 6 + 6},
			{"55 + 55", 55 + 55},
		},
		"operator minus": {
			{"5-3", 2},
			{"17 - 23", 17 - 23},
		},
		"operator mul": {
			{"2*3", 6},
			{"45*13", 45 * 13},
		},
		"operator div": {
			{"1/3", 1.0 / 3},
			{"14221/1234", 14221. / 1234},
		},
		"different operators": {
			{"1/3 + 7 - 4 * 3", 1.0/3 + 7 - 4*3},
			{"12*3-4/2", 12*3 - 4/2},
		},
	}
	postfixParser := &PostfixParser{}
	t.Parallel()
	for name, testCases := range cases {
		for _, tC := range testCases {
			t.Run(name, func(t *testing.T) {
				result, err := postfixParser.ParseAndCalculate(tC.in)
				require.Equal(t, result, tC.out)
				require.NoError(t, err)
			})
		}
	}
}

func TestUnaryMinus(t *testing.T) {
	type testCase struct {
		in  string
		out float64
	}

	cases := map[string][]testCase{
		"double minus": {
			{"5--5", 5 - (-5)},
			{"12--13", 12 - (-13)},
		},
		"simple unary": {
			{"-3", -3},
			{"-143", -143},
		},
		"unary with operators": {
			{"2*-3", 2 * (-3)},
			{"45+-13", 45 + (-13)},
			{"13/-2", 13. / (-2)},
		},
	}
	postfixParser := &PostfixParser{}
	t.Parallel()
	for name, testCases := range cases {
		for _, tC := range testCases {
			t.Run(name, func(t *testing.T) {
				result, err := postfixParser.ParseAndCalculate(tC.in)
				require.Equal(t, result, tC.out)
				require.NoError(t, err)
			})
		}
	}
}

func TestWithBrackets(t *testing.T) {
	type testCase struct {
		in  string
		out float64
	}

	cases := map[string][]testCase{
		"simple brackets": {
			{"5--(-5)", 5 + (-5)},
			{"(12 + 13) * 3", (12 + 13) * 3},
		},
		"long expressions": {
			{"(1.9999999-1.9999999)/(0.3-0.29999)*999999999", 0.0},
			{"25*0/5.0", 0.0},
			{"345.2928 + (7374.4332-1.) - 847.958222*23.949032 - ((763.541 - .531) / 82.09093 * (522.00001 + 22.4373*88.146 - 115.7*2.323467))",
				345.2928 + (7374.4332 - 1.) - 847.958222*23.949032 - ((763.541 - .531) / 82.09093 * (522.00001 + 22.4373*88.146 - 115.7*2.323467)),
			},
		},
	}
	postfixParser := &PostfixParser{}
	t.Parallel()
	for name, testCases := range cases {
		for _, tC := range testCases {
			t.Run(name, func(t *testing.T) {
				result, err := postfixParser.ParseAndCalculate(tC.in)
				require.Equal(t, result, tC.out)
				require.NoError(t, err)
			})
		}
	}
}

func TestErrors(t *testing.T) {
	type testCase struct {
		in  string
		out float64
	}

	cases := map[string][]testCase{
		"errors": {
			{"5--(-5))", 5 + (-5)},
			{"(12 -+ 13) * 3", (12 + 13) * 3},
			{"adsf", 0},
		},
	}
	postfixParser := &PostfixParser{}
	t.Parallel()
	for name, testCases := range cases {
		for _, tC := range testCases {
			t.Run(name, func(t *testing.T) {
				_, err := postfixParser.ParseAndCalculate(tC.in)
				require.Error(t, err)
			})
		}
	}
}
