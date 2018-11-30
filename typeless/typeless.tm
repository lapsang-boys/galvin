language typeless(go);

lang = "typeless"
package = "github.com/lapsang-boys/galvin/typeless"
eventBased = true
eventFields = true
eventAST = true

:: lexer

identifier_tok: /[a-zA-Z_][a-zA-Z0-9_]*/
'(': /\(/
')': /\)/
'[': /\[/
']': /\]/
',': /,/
'\\': /\\/
'->': /->/
intLit: /[-]?([0-9]|[1-9][0-9]+)/

:: parser

%input File;

File -> File
   : Expression*
;

%interface Expression;

Expression -> Expression
   : FunctionAbstraction
   | FunctionApplication
   | Literal
   | Identifier
;

FunctionAbstraction -> FunctionAbstraction
   : '(' '\\' Parameters=(Identifier separator ',')* '->' Body ')'
;

Body -> Body
   : Expression
;

FunctionApplication -> FunctionApplication
   : FunctionName=FunctionAbstraction '[' Arguments=(Expression separator ',')* ']'
;

Literal -> Literal
   : intLit
;

Identifier -> Identifier
   : identifier_tok
;
