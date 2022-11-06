package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/weiyouwozuiku/Pangu/consts"
	"github.com/weiyouwozuiku/Pangu/lexer"
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
