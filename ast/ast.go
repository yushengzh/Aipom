/**
 * AST:抽象语法树
 */

package ast

import (
	"Aipom/token"
)

/*
 * interface
 * AST需要两种类型节点：语句和表达式
 */
type Node interface {
	TokenLiteral() string //该方法返回与其关联lex unit的字面量
}

/* 语句 */
type Statement interface {
	Node
	statementNode() //占位方法
}

/* 表达式 */
type Expression interface {
	Node
	expressionNode()
}

/* Program节点：每个AST的root node */
type Program struct {
	Statements []Statement //放置程序中的所有语句
}

func (pro *Program) TokenLiteral() string {
	if len(pro.Statements) > 0 {
		return pro.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

/* 标识符 */
type Identifier struct {
	Token token.Token
	Value string
}

/* let语句产生值的表达式 */
type LetStatement struct {
	Token token.Token
	Name  *Identifier //绑定的标识符
	Value Expression  //产生值的表达式
}

/* 实现Statement接口 */
func (ls *LetStatement) statementNode() {}

/* 实现Node接口 */
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (id *Identifier) expressionNode() {}

func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}
