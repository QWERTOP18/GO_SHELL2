package lexer

import (
	"testing"

	"shell/token"
)

func TestNextToken(t *testing.T) {
	input := `grep -v go < out && echo 'Hello Kotaro' >> /dev/null`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.WORD, "grep"},
		{token.WORD, "-v"},
		{token.WORD, "go"},
		{token.REDIRECT, "<"},
		{token.WORD, "out"},
		{token.AND, "&&"},
		{token.WORD, "echo"},
		{token.WORD, "'Hello Kotaro'"},
		{token.REDIRECT, ">>"},
		{token.WORD, "/dev/null"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		//fmt.Println(i, tok.Literal)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
