package token

// the tokens
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// each token type

var identKeywords = map[string]TokenType{
	"fun": FUNCTION,
	"let": LET,
	"fax": TRUE,
	"cap": FALSE,
	"fr": IF,
	"otherwise": ELSE,
	"hereis": RETURN,
}

func GetIdent(ident string) TokenType {
	if tok, ok := identKeywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	BANG	TokenType = "!"
	TRUE TokenType = "TRUE"
	FALSE TokenType = "FALSE"
	EOF     TokenType = "EOF"
	// Identifiers + literals
	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"
	// Operators
	ASSIGN TokenType = "="
	PLUS   TokenType = "+"
	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"
	// Keywords
	FUNCTION TokenType = "FUNCTION"
	IF 	 TokenType = "IF"
	ELSE	 TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
	LET      TokenType = "LET"
	SLASH   TokenType = "/"
	ASTERISK TokenType = "*"
	MINUS TokenType = "-"
	LESS_THAN TokenType = "<"
	GREATER_THAN TokenType = ">"
)
