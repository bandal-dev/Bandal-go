package main

import (
	"fmt"
	"bandal/parser"
	"bandal/compile"
)

//To Do -> argparse
var source string = 
`fn main()
	pass
`

func main() {
	var ret "parser type" = compile.Compile(source, true)
	if ret != nil {
		//pickle
	}
}
