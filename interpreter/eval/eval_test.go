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
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
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
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanData(t, evaluated, tt.expected)
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
		testBooleanData(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)

		if ok {
			testIntegerData(t, evaluated, int64(integer))
		} else {
			testNullData(t, evaluated)
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{`
			if (10 > 1) {
				if (10 > 1) {
					return 10;
				}
				return 1;
			}
		`,
			10,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerData(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"Type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"Type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"Unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"Unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"Unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"Unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
				if (10 > 1) {
					if (10 > 1) {
						return true + false;
					}
					return 1;
				}
			`,
			"Unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"Identifier not found: foobar",
		},
		{
			`"Hello" - "World"`,
			"Unknown operator: STRING - STRING",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		err, ok := evaluated.(*data.Error)

		if !ok {
			t.Errorf("Expected Error, got %T(%+v)",
				evaluated, evaluated)
			continue
		}

		if err.Message != tt.expectedMessage {
			t.Errorf("Expected Message to be %q, got %q",
				tt.expectedMessage, err.Message)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
	}

	for _, tt := range tests {
		testIntegerData(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionData(t *testing.T) {
	input := "fn(x) { x + 2; };"
	evaluated := testEval(input)
	fn, ok := evaluated.(*data.Function)

	if !ok {
		t.Fatalf("Expected Data to be Function, got %T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("Expected 1 Parameter, got %+v", fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("Expected first Parameter to be 'x', got %q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"
	if fn.Body.String() != expectedBody {
		t.Fatalf("Expected Body to be %q, got %q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let identity = fn(x) { x; }; identity(5);", 5},
		{"let identity = fn(x) { return x; }; identity(5);", 5},
		{"let double = fn(x) { x * 2; }; double(5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5, 5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn(x) { x; }(5)", 5},
	}

	for _, tt := range tests {
		testIntegerData(t, testEval(tt.input), tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
		let newAdder = fn(x) {
			fn(y) { x + y };
		};

		let addTwo = newAdder(2);
		addTwo(2);
	`

	testIntegerData(t, testEval(input), 4)
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`
	evaluated := testEval(input)
	str, ok := evaluated.(*data.String)

	if !ok {
		t.Fatalf("Expected Data to be String, got %T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("Expected String to be 'Hello World!', got %q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`
	evaluated := testEval(input)
	str, ok := evaluated.(*data.String)

	if !ok {
		t.Fatalf("Expected Data to be String, got %T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("Expected String to be 'Hello World!', got %q", str.Value)
	}
}

func testEval(input string) data.Data {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := data.NewEnvironment()

	return Eval(program, env)
}

func testIntegerData(t *testing.T, d data.Data, expected int64) bool {
	result, ok := d.(*data.Integer)

	if !ok {
		t.Errorf("Expected Data to be Integer, got %T (%+v)", d, d)
		return false
	}

	if result.Value != expected {
		t.Errorf("Expected Data to equal %d, got %d", expected, result.Value)
		return false
	}

	return true
}

func testBooleanData(t *testing.T, d data.Data, expected bool) bool {
	result, ok := d.(*data.Boolean)
	if !ok {
		t.Errorf("Expected Data to be Boolean, got %T (%+v)", d, d)
		return false
	}

	if result.Value != expected {
		t.Errorf("Expected Data to equal %t, got %t", expected, result.Value)
		return false
	}

	return true
}

func testNullData(t *testing.T, d data.Data) bool {
	if d != NULL {
		t.Errorf("Expected Data to be Null, got %T (%+v)", d, d)
		return false
	}
	return true
}
