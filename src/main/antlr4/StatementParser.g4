// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

parser grammar StatementParser;

options {
    tokenVocab = StatementLexer;
}

statement: value EOF | EOF;

value:
    array
    | object
    | BOOLEAN
    | INTEGER
    | FLOAT
    | BINARY
    | HEXADECIMAL
    | STRING
    | IPV4
    | IPV6
    | MAC
    | COLOR
    | SEMVER
    | DURATION
    | PERCENTAGE
;

array: '[' (value (',' value)* ','?)? ']';

object: '{' (ID ':' value (',' ID ':' value)* ','?)? '}';