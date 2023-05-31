package repl

import (
	"bufio"
	"fmt"
	"io"
	"xhod/pkg/lexer"
	"xhod/pkg/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		lxr := lexer.New(line)
		for tok := lxr.ReadToken(); tok.Type != token.EOF; tok = lxr.ReadToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
