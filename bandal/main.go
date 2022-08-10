package main

import (
	"errors"
	"fmt"
	"os"
	"os/user"

	"github.com/Bandal-dev/Bandal-go/bandal/repl"
)

//source code here
var errUser error = errors.New("ERROR")

func main() {
	_, err := user.Current()
	if err != nil {
		panic(errUser)
	}
	fmt.Printf("BanDal Beta REPL")
	repl.Start(os.Stdin, os.Stdout)
}
