package token

// Token type
type TokenType int
type Token struct {
	Type    TokenType
	Literal string
}

// metacharacter
//
//	A character that, when unquoted,
//	separates words.  One of the
//	following:
//	|  & ; ( ) < > space tab
const METACHARS = "|&;()<> \t"

const (
	ILLEGAL TokenType = iota
	NAME
	WORD
	REDIRECT
	NUMBER
	ASSIGN
	EOF

	/* Reserved Words */
	FUNCTION
	IF
	ELSE
	NOT
	DO
	DONE
	ELIF
	ESAC
	FI
	FOR
	IN
	SELECT
	THEN
	UNTIL
	WHILE
	LBRACE
	RBRACE
	TIME
	LBRACKET
	RBRACKET

	/* Control Operator */
	OR
	ASYNC
	PIPE
	AND
	LPAREN
	RPAREN
	SEMICOLON
	DOUBLE_SEMICOLON
	NEWLINE
)

// control operator
//
//	A token that performs a control
//	function.  It is one of the
//	following symbols:
//	|| & && ; ;; ( ) | <newline>
var controlOperator = map[string]TokenType{
	"||": OR,
	"&":  ASYNC,
	"|":  PIPE,
	"&&": AND,
	"(":  LPAREN,
	")":  RPAREN,
	";":  SEMICOLON,
	"\n": NEWLINE,
	";;": DOUBLE_SEMICOLON,
	"<":  REDIRECT,
	">":  REDIRECT,
}

// ! case  do done elif else esac fi for
//
//	function if in select then until while {
//	} time [[ ]]
var reservedWords = map[string]TokenType{
	"!":        NOT,
	"case":     FUNCTION,
	"do":       DO,
	"done":     DONE,
	"elif":     ELIF,
	"else":     ELSE,
	"esac":     ESAC,
	"fi":       FI,
	"for":      FOR,
	"function": FUNCTION,
	"if":       IF,
	"in":       IN,
	"select":   SELECT,
	"then":     THEN,
	"until":    UNTIL,
	"while":    WHILE,
	"{":        LBRACE,
	"}":        RBRACE,
	"time":     TIME,
	"[[":       LBRACKET,
	"]]":       RBRACKET,
}
