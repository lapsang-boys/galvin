language typeless(go);

lang = "typeless"
package = "github.com/llir/ll"
eventBased = true
eventFields = true
eventAST = true

:: lexer

identifier: /[a-zA-Z_][a-zA-Z0-9_]+/
'(': /\(/
')': /\)/
'[': /\[/
']': /\]/
',': /,/
'\\': /\\/
'->': /->/
'_': /_/
intLit: /[-]?([0-9]|[1-9][0-9]+)/


# (\x, y -> x+y)(x, y)

:: parser

%input File;

File: Expression*;

Expression: FunctionAbstraction | FunctionApplication | Literal;

FunctionAbstraction: '(' '\\' ( (identifier separator ',')* | '_' ) '->' Expression ')';

FunctionApplication: FunctionAbstraction '[' (Expression separator ',')* ']';

Literal: intLit;
