// Code generated from StatementParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // StatementParser

import "github.com/antlr4-go/antlr/v4"

type BaseStatementParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseStatementParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStatementParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStatementParserVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStatementParserVisitor) VisitObject(ctx *ObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStatementParserVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStatementParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
