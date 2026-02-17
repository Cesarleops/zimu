package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/cesarleops/zimu/internal/lexer"
	"github.com/cesarleops/zimu/internal/token"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	fmt.Println("Happy hacking!")
	for {
		fmt.Print("> ")
		ready := scanner.Scan()
		if !ready {
			return
		}
		line := scanner.Text()

		l := lexer.NewLexer(line, "<repl>")

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v \n", tok)
		}
	}
}
