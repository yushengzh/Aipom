package token

//类型
type TokenType string

//结构
type Token struct {
	Type    TokenType
	Literal string
}

/*
 * 标识符
 * 单字符byte 多字符rune
 */
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	//运算符
	ASSIGN       = "="
	PLUS         = "+"
	DOUBLE_PLUS  = "++"
	EQUAL_PLUS   = "+="
	MINUS        = "-"
	DOUBLE_MINUS = "--"
	EQUAL_minus  = "-="
	MOD          = "%"
	MUL          = "*"
	DIV          = "/"
	AND          = "&&"
	OR           = "||"
	NOT          = "!"
	LEG          = "<="
	LT           = "<"
	GEQ          = ">="
	GT           = ">"

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
	FOR      = "FOR"
	IF       = "IF"
	ELSE     = "ELSE"
)

// 区分关键字和用户定义的标识符
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

/*
 * 检查关键字表来判断给定标识符是否是关键字
 * 是-->返回TokenType常量 不是-->返回token.IDENT
 */
func LookupId(id string) TokenType {
	if tok, ok := keywords[id]; ok {
		return tok
	}
	return IDENT
}
