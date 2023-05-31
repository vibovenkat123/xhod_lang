package lexer
import "xhod/pkg/token"
type Lexer struct {
    input string
    pos int // current char
    readPos int // next char
    ch byte // curr char
}

func New(input string) *Lexer {
    lxr := &Lexer{input: input}
    lxr.readChar()
    return lxr
}

func (lxr *Lexer) ReadToken() token.Token {
    var tok token.Token
    lxr.destroyWhitespace()
    switch lxr.ch {
    case '=':
        tok = newToken(token.ASSIGN, lxr.ch)
    case ';':
        tok = newToken(token.SEMICOLON, lxr.ch)
    case '(':
        tok = newToken(token.LPAREN, lxr.ch)
    case ')':
        tok = newToken(token.RPAREN, lxr.ch)
    case ',':
        tok = newToken(token.COMMA, lxr.ch)
    case '+':
        tok = newToken(token.PLUS, lxr.ch)
    case '{':
        tok = newToken(token.LBRACE, lxr.ch)
    case '}':
        tok = newToken(token.RBRACE, lxr.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    default:
        if isLetter(lxr.ch) {
            tok.Literal = lxr.readIdent()
            tok.Type = token.GetIdent(tok.Literal)
            return tok
        } else if isDigit(lxr.ch) {
            tok.Type = token.INT
            tok.Literal = lxr.readNum()
            return tok
        } else {
            tok = newToken(token.ILLEGAL, lxr.ch)
        }
    }
    lxr.readChar()
    return tok
}

func (lxr *Lexer) readIdent() string {
    pos := lxr.pos
    for isLetter(lxr.ch) {
        lxr.readChar()
    }
    // the chars between the start and end of the identifier
    return lxr.input[pos:lxr.pos]
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}


// read next char and move forward
func (lxr *Lexer) readChar() {
    if lxr.readPos >= len(lxr.input) {
        lxr.ch = 0
    } else {
        lxr.ch = lxr.input[lxr.readPos]
    }
    lxr.pos = lxr.readPos
    lxr.readPos ++
}

func (lxr *Lexer) destroyWhitespace() {
    for lxr.ch == ' ' || lxr.ch == '\t' || lxr.ch == '\n' || lxr.ch == '\r' {
        lxr.readChar()
    }
}

func (lxr *Lexer) readNum() string {
    pos := lxr.pos
    for isDigit(lxr.ch) {
        lxr.readChar()
    }
    return lxr.input[pos:lxr.pos]
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}
