// repl is a real-eval-print loop. For the time being, it only
// prints out the items it has tokenized, but it does not parse
// or evaluate yet.
package repl

import (
	"bufio"
	"fmt"
	"io"
	"bee/lexer"
	"bee/token"
)

const PROMPT = ">> "

// Start is what starts the REPL.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			return
		}

		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}