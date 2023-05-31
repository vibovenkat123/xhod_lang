package lexer

import (
	"testing"
	"xhod/pkg/token"
)

func TestNextToken(t *testing.T) {
    input := `
    let five = 5;
    let eight = 8;
    let add = fun(x, y) {
        x + y;
    };
    let result = add(five, eight);
    !-/*5;
    5 < 12 > 5;
    fr (5 < 10) {
        hereis fax;
    } otherwise {
        hereis cap;
    }
`
    tests := []struct {
        expectedType token.TokenType
        expectedLiteral string
       }{
        {token.LET, "let"},
        {token.IDENT, "five"},
        {token.ASSIGN, "="},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENT, "eight"},
        {token.ASSIGN, "="},
        {token.INT, "8"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENT, "add"},
        {token.ASSIGN, "="},
        {token.FUNCTION, "fun"},
        {token.LPAREN, "("},
        {token.IDENT, "x"},
        {token.COMMA, ","},
        {token.IDENT, "y"},
        {token.RPAREN, ")"},
        {token.LBRACE, "{"},
        {token.IDENT, "x"},
        {token.PLUS, "+"},
        {token.IDENT, "y"},
        {token.SEMICOLON, ";"},
        {token.RBRACE, "}"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENT, "result"},
        {token.ASSIGN, "="},
        {token.IDENT, "add"},
        {token.LPAREN, "("},
        {token.IDENT, "five"},
        {token.COMMA, ","},
        {token.IDENT, "eight"},
        {token.RPAREN, ")"},
        {token.SEMICOLON, ";"},
        {token.BANG, "!"},
        {token.MINUS, "-"},
        {token.SLASH, "/"},
        {token.ASTERISK, "*"},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.INT, "5"},
        {token.LESS_THAN, "<"},
        {token.INT, "12"},
        {token.GREATER_THAN, ">"},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.IF, "fr"},
        {token.LPAREN, "("},
        {token.INT, "5"},
        {token.LESS_THAN, "<"},
        {token.INT, "10"},
        {token.RPAREN, ")"},
        {token.LBRACE, "{"},
        {token.RETURN, "hereis"},
        {token.TRUE, "fax"},
        {token.SEMICOLON, ";"},
        {token.RBRACE, "}"},
        {token.ELSE, "otherwise"},
        {token.LBRACE, "{"},
        {token.RETURN, "hereis"},
        {token.FALSE, "cap"},
        {token.SEMICOLON, ";"},
        {token.RBRACE, "}"},
        {token.EOF, ""},
       }
    lxr := New(input)
    // loop through the tests
    for i, test := range tests {
        tok := lxr.ReadToken()
        if tok.Type != test.expectedType {
            t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
            i, test.expectedType, tok.Type)
        }
        if tok.Literal != test.expectedLiteral {
            t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.Literal)
        }
    }
}
