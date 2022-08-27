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
	ch      rune
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
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0 //NUL字符的unicode编码
	}
}

// 初始化词法单元
func newToken(tokenType token.TokenType, curChar byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(curChar),
	}
}

/*
 * 输出下一个词法单元
 * NextToken()的基本结构: 检查curchar 根据具体字符返回相应的词法单元 返回之前指针前移 调用newToken()初始化这些词法单元
 */
func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token
	lexer.skipWhitespace() //
	switch lexer.curChar {
	case '=':
		tok = newToken(token.ASSIGN, lexer.curChar)
	case '+':
		tok = newToken(token.PLUS, lexer.curChar)
	//case "++":
	//	tok = newToken(token.DOUBLE_PLUS, lexer.curChar)
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
	default:
		if isLetter(lexer.curChar) {
			// 标识符和关键字的词法分析
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupId(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.curChar) //非字母字符 --> 返回token.ILLEGAL
		}
	}
	lexer.readCharascii()
	return tok
}

// 判断读入字符是否为字母 (下划线_允许在标识符和关键字中使用)
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '-'
}

/*
 * 读入一个标识符并前移lexer的扫描位置
 * 当遇到非字母字符后停止
 */
func (lexer *Lexer) readIdentifier() string {
	pos := lexer.position
	for isLetter(lexer.curChar) {
		lexer.readCharascii()
	}
	return lexer.input[pos:lexer.position]
}

/*
 * 跳过空白字符' '
 */
func (lexer *Lexer) skipWhitespace() {
	for lexer.curChar == ' ' || lexer.curChar == '\r' || lexer.curChar == '\t' || lexer.curChar == '\n' {
		lexer.readChar()
	}
}
