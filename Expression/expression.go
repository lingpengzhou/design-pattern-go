package Expression

import "strings"

type Expression interface {
	interpret(context string) bool
}

type TerminalExpression struct {
	data string
}

func (terminalExpression *TerminalExpression) interpret(context string) bool {
	if strings.Contains(context, terminalExpression.data) {
		return true
	}
	return false
}

type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func SetOrExpression(expr1 Expression, expr2 Expression) *OrExpression {
	return &OrExpression{
		expr1: expr1, expr2: expr2,
	}
}

func (orExpression *OrExpression) interpret(context string) bool {
	return orExpression.expr1.interpret(context) || orExpression.expr2.interpret(context)
}

type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func SetAndExpression(expr1 Expression, expr2 Expression) *AndExpression {
	return &AndExpression{
		expr1: expr1, expr2: expr2,
	}
}

func (andExpression *AndExpression) interpret(context string) bool {
	return andExpression.expr1.interpret(context) && andExpression.expr2.interpret(context)
}
