package lexer

import (
	"Aipom/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	//input := `=+(){},;-%*/`
	input := `let five = 5;
let ten = 10;
let add = fn(x, y){
	x + y;
};

let results = add(five, ten);
!-*/%^
1 < 11 > 1

if (1 < 11){
	return true;
} else {
	return false;
}

1 == 1;
1 != 11;
++ += >= < <= > # & || && |
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		/*
			{token.ASSIGN, "="},
			{token.PLUS, "+"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
			{token.COMMA, ","},
			{token.SEMICOLON, ";"},
			{token.MINUS, "-"},
			{token.MOD, "%"},
			{token.MUL, "*"},
			{token.DIV, "/"},
			{token.EOF, ""},
		*/
		{token.LET, "let"}, {token.IDENT, "five"}, {token.ASSIGN, "="}, {token.INT, "5"}, {token.SEMICOLON, ";"},
		{token.LET, "let"}, {token.IDENT, "ten"}, {token.ASSIGN, "="}, {token.INT, "10"}, {token.SEMICOLON, ";"},
		{token.LET, "let"}, {token.IDENT, "add"}, {token.ASSIGN, "="}, {token.FUNCTION, "fn"}, {token.LPAREN, "("}, {token.IDENT, "x"},
		{token.COMMA, ","}, {token.IDENT, "y"}, {token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"}, {token.PLUS, "+"}, {token.IDENT, "y"}, {token.SEMICOLON, ";"},
		{token.RBRACE, "}"}, {token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "results"}, {token.ASSIGN, "="}, {token.IDENT, "add"}, {token.LPAREN, "("}, {token.IDENT, "five"}, {token.COMMA, ","},
		{token.IDENT, "ten"}, {token.RPAREN, ")"}, {token.SEMICOLON, ";"},

		/*
			! - * / % ^
				1 < 11 > 1
		*/
		{token.BANG, "!"}, {token.MINUS, "-"}, {token.MUL, "*"},
		{token.DIV, "/"}, {token.MOD, "%"}, {token.POW, "^"},
		{token.INT, "1"}, {token.LT, "<"}, {token.INT, "11"}, {token.GT, ">"}, {token.INT, "1"},

		{token.IF, "if"}, {token.LPAREN, "("}, {token.INT, "1"}, {token.LT, "<"}, {token.INT, "11"},
		{token.RPAREN, ")"}, {token.LBRACE, "{"}, {token.RETURN, "return"}, {token.TRUE, "true"}, {token.SEMICOLON, ";"},
		{token.RBRACE, "}"}, {token.ELSE, "else"}, {token.LBRACE, "{"}, {token.RETURN, "return"}, {token.FALSE, "false"}, {token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.INT, "1"}, {token.EQ, "=="}, {token.INT, "1"}, {token.SEMICOLON, ";"},
		{token.INT, "1"}, {token.NOT_EQ, "!="}, {token.INT, "11"}, {token.SEMICOLON, ";"},

		//++ += >= < <= > # & || && |
		{token.DOUBLE_PLUS, "++"}, {token.EQUAL_PLUS, "+="}, {token.GEQ, ">="}, {token.LT, "<"},
		{token.LEG, "<="}, {token.GT, ">"}, {token.ANNOTATION, "#"},
		{token.BITWISE_AND, "&"}, {token.OR, "||"}, {token.AND, "&&"},
		{token.BITWISE_OR, "|"},

		{token.EOF, ""},
	}

	lexer := New(input)
	for i, tt := range tests {
		tok := lexer.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
