package eval

import (
	"ape/interpreter/data"
	"ape/interpreter/lexer"
	"ape/interpreter/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerData(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) data.Data {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	return Eval(program)
}

func testIntegerData(t *testing.T, obj data.Data, expected int64) bool {
	result, ok := obj.(*data.Integer)

	if !ok {
		t.Errorf("Expected Data to be Integer, got %T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("Expected Data to equal %d, got %d", expected, result.Value)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj data.Data, expected bool) bool {
	result, ok := obj.(*data.Boolean)
	if !ok {
		t.Errorf("Expected Data to be Boolean, got %T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("Expected Data to equal %t, got %t", expected, result.Value)
		return false
	}

	return true
}
