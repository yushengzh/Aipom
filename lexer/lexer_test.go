package lexer

import (
	"Aipom/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	//input := `=+(){},;-%*/`
	input := `
	let five = 5;
	let ten = 10;
	let add = fn(x, y){
		x + y;
	};
	let results = add(five, ten);
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
		{token.IDENT, "x"}, {token.PLUS, "+"}, {token.IDENT, "y"},
		{token.RBRACE, "}"}, {token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "results"}, {token.ASSIGN, "="}, {token.IDENT, "add"}, {token.LPAREN, "("}, {token.IDENT, "five"}, {token.COMMA, ","},
		{token.IDENT, "ten"}, {token.RPAREN, ")"}, {token.SEMICOLON, ";"},
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
