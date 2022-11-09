package repl

import (
	"bufio"
	"fmt"
	"github.com/weiyouwozuiku/Pangu/lexer"
	"io"

	"github.com/weiyouwozuiku/Pangu/consts"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, consts.PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		//p := parser.New(l)
	}
}
