// Code generated from StatementParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // StatementParser

import "github.com/antlr4-go/antlr/v4"

// BaseStatementParserListener is a complete listener for a parse tree produced by StatementParser.
type BaseStatementParserListener struct{}

var _ StatementParserListener = &BaseStatementParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStatementParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStatementParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStatementParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStatementParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseStatementParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseStatementParserListener) ExitStatement(ctx *StatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStatementParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStatementParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterArray is called when production array is entered.
func (s *BaseStatementParserListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseStatementParserListener) ExitArray(ctx *ArrayContext) {}

// EnterObject is called when production object is entered.
func (s *BaseStatementParserListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseStatementParserListener) ExitObject(ctx *ObjectContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseStatementParserListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseStatementParserListener) ExitPair(ctx *PairContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseStatementParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseStatementParserListener) ExitLiteral(ctx *LiteralContext) {}
