package lexer

import "Aipom/token"

type Lexer struct {
	input string

	//输入字符串的当前位置
	position int

	//输入字符串的当前读取位置(当前字符之后的一个字符)
	readPosition int

	//当前正在查看的字符
	curChar byte
}

//输入的源码(字符串)转换为Lexer结构体
func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readCharascii()
	return lexer
}

func (lexer *Lexer) readCharascii() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.curChar = 0 //NUL字符的ASCII编码
	} else {
		lexer.curChar = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

//TO DO
//实现lexer支持unicode字符
func (lexer *Lexer) readChar() {

}

func newToken(tokenType token.TokenType, curChar byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(curChar),
	}
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token
	switch lexer.curChar {
	case '=':
		tok = newToken(token.ASSIGN, lexer.curChar)
	case '+':
		tok = newToken(token.PLUS, lexer.curChar)
	case '-':
		tok = newToken(token.MINUS, lexer.curChar)
	case '*':
		tok = newToken(token.MUL, lexer.curChar)
	case '/':
		tok = newToken(token.DIV, lexer.curChar)
	case '%':
		tok = newToken(token.MOD, lexer.curChar)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.curChar)
	case ',':
		tok = newToken(token.COMMA, lexer.curChar)
	case '(':
		tok = newToken(token.LPAREN, lexer.curChar)
	case ')':
		tok = newToken(token.RPAREN, lexer.curChar)
	case '{':
		tok = newToken(token.LBRACE, lexer.curChar)
	case '}':
		tok = newToken(token.RBRACE, lexer.curChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	lexer.readCharascii()
	return tok
}
