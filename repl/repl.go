// repl/repl.go

package repl

import (
	"Aipom/lexer"
	"Aipom/token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

/*
 * 读取输入源码，直到读完一行代码
 * 读取的代码行传递给实例化的Lexer
 * 输出Lexer生成的词法单元，直到遇到EOF
 */
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
