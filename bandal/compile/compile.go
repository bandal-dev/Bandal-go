package compile

import (
	"fmt"
	"bandal/token"
	"bandal/lexer"
	"bandal/parser"
	"bandal/vm"
)

func Compile(src string, run bool) "parser type" {
	var l []token.Token = lexer.Lex(src)
	var p "parser type" = parser.Parse(l)
	if run {
		// vm.Run(p) //run
		// return nil
	} else {
		// return p //build only, write in file by using pickle
	}
}