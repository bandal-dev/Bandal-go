package main

import (
	"fmt"
	"go/token"

	"github.com/Bandal-dev/Bandal-go/bandal/lexer"
)

//source code here

var source string = `
false
true
if
wllse
else
2134
k 1 
rl1w
`

var l []token.Token = lexer.Lex(source)

func main() {
	for _, i := range l {
		fmt.Println(i)
	}
}
