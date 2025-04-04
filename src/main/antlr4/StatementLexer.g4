// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar StatementLexer;


COMMA  : ',';
COLON  : ':';
LSQUAR : '[';
RSQUAR : ']';
LCURLY : '{';
RCURLY : '}';

BOOLEAN : 'true' | 'false';

fragment SIGN      : [+-];
fragment BIN       : [01];
fragment DEC       : [0-9];
fragment HEX_UPPER : [0-9A-F];
fragment HEX_LOWER : [0-9a-f];
fragment HEX       : HEX_UPPER | HEX_LOWER;

INTEGER     : SIGN? DEC+;
FLOAT       : SIGN? DEC+ '.' DEC+;
BINARY      : SIGN? '0b' BIN+;
HEXADECIMAL : SIGN? '0x' HEX+;

STRING: '"' .*? '"';
// todo: escape quotes
// todo: multiline strings

fragment IP_OCTET  : DEC | [1-9] DEC | '1' DEC DEC | '2' [0-4] DEC | '25' [0-5];
fragment MAC_OCTET : HEX HEX;
fragment H16       : HEX HEX? HEX? HEX?;
fragment LS32      : (H16 ':' H16) | IPV4;

MAC  : MAC_OCTET ':' MAC_OCTET ':' MAC_OCTET ':' MAC_OCTET ':' MAC_OCTET ':' MAC_OCTET;
IPV4 : IP_OCTET '.' IP_OCTET '.' IP_OCTET '.' IP_OCTET | 'localhost';
IPV6 : H16 ':' H16 ':' H16 ':' H16 ':' H16 ':' H16 ':' H16 ':' H16 | '::1';
// todo: allow for shortened ipv6

fragment COLOR_SHORT       : HEX HEX HEX;
fragment COLOR_SHORT_ALPHA : COLOR_SHORT HEX;
fragment COLOR_FULL        : COLOR_SHORT COLOR_SHORT;
fragment COLOR_FULL_ALPHA  : COLOR_FULL HEX HEX;

COLOR: '#' ( COLOR_SHORT | COLOR_SHORT_ALPHA | COLOR_FULL | COLOR_FULL_ALPHA);

SEMVER: DEC '.' DEC '.' DEC;
// todo: handle all valid semver

DURATION: (INTEGER | FLOAT) ('s' | 'm' | 'h' | 'ms');
// todo: add more types of duration and composablility (1m30s)

ID: [a-zA-Z_][a-zA-Z_0-9]*;

WS                : [ \t\n\r\f]+  -> skip;
COMMENT           : '//' ~[\n\r]* -> skip;
MULTILINE_COMMENT : '/*' .*? '*/' -> skip;