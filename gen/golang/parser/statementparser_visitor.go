// Code generated from StatementParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // StatementParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by StatementParser.
type StatementParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by StatementParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by StatementParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by StatementParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by StatementParser#object.
	VisitObject(ctx *ObjectContext) interface{}

	// Visit a parse tree produced by StatementParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by StatementParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
