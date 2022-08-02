package vm

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"bandal/parser"
)

var errCantReadFile = errors.New("Cannot Read Current File")
var errRuntimeError = errors.New("RunTime Error")
var fileData string = ""
var errNothingToPop = errors.New("Stack has been popped without anything in it")
var errStackEmpty = errors.New("Stack is Empty")

// Struct Data Type
type Stack struct {
	stack []int
}

// main
func main() {
	ReadFile()
	fmt.Println(fileData)
	eval()
}

// eval
func eval() error {
	var stack Stack
	var PTR int = 0
	const MAX = 65535 //2^16 - 1\
	var err error = nil
	content := strings.Split(fileData, "\n")

	for _, str := range content {
		if PTR > MAX {
			return errRuntimeError
		}

		rule := regexp.MustCompile(" ")
		op := rule.Split(str, strings.Count(str, " ")+1) // 공백 갯수 + 1 ex> "push 1" -> 2개
		//operand = make([]string, 10)
		fmt.Println(op)

		switch op[0] {
		case "add":
			stack.Push(stack.Pop().(float64) + stack.Pop().(float64))
		case "sub":
			last = stack.Pop().(float64)
			first = stack.Pop().(float64)
			stack.Push(first - last)
		case "mul":
			stack.Push(stack.Pop().(float64) * stack.Pop().(float64))
		case "div":
			last = stack.Pop().(float64)
			first = stack.Pop().(float64)
			stack.Push(first / last)
		case "push":
			//
		case "lds":
			//
		case "load":
			//
		case "store":
			//

		}

	}
	return err
}

//ReadFile : 파일 처리(읽기)!
func ReadFile() error {
	file, err := os.Open("src/src.basm")
	if err != nil {
		return errCantReadFile
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		fileData += string(fileScanner.Text()) + "\n"
	}
	return nil
}

func (s Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) Push(data int) {
	s.stack = append(s.stack, data)
}

func (s *Stack) Pop() (int, error) {
	var err error = nil
	if s.IsEmpty() {
		err = errNothingToPop
	}

	poppedValue := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return poppedValue, err

}

func (s Stack) Top() (int, error) {
	var err error = nil
	if s.IsEmpty() {
		err = errStackEmpty
	}
	top := len(s.stack) - 1
	return s.stack[top], err
}