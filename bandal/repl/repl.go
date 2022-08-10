package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Bandal-dev/Bandal-go/bandal/lexer"
	"github.com/Bandal-dev/Bandal-go/bandal/token"
)

const PROMT = ">>>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	fmt.Fprintf(out, "\n")

	for {

		fmt.Fprintf(out, PROMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "!killbandalide!" {
			os.Exit(3)
		}
		l := lexer.Lex(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
