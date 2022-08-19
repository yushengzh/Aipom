package token

//类型
type TokenType string

//结构
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	//运算符
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	MOD    = "%"
	MUL    = "*"
	DIV    = "/"

	//分隔符
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
