package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"strconv"
	"reflect"
	"regexp"
	//"bandal/parser"
)

var errCantReadFile = errors.New("Cannot Read Current File")
var errRuntimeError = errors.New("RunTime Error")
var fileData string = ""
var errNothingToPop = errors.New("Stack has been popped without anything in it")
var errStackEmpty = errors.New("Stack is Empty")
var DEBUG bool = true

// Struct Data Type
type Stack struct {
	stack []interface{}
}

// main
func main() {
	ReadFile()
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

		op := strings.Fields(string(str))
		if len(op) == 0 {
			continue
		}
		//operand = make([]string, 10)

		switch op[0] {
		case "add":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first).Kind() {
				case reflect.Slice:
					// stack.Push(append(reflect.ValueOf(first), reflect.ValueOf(last)))
				default:
					switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(first.(int8) + last.(int8))
					case reflect.TypeOf(int16(1)): stack.Push(first.(int16) + last.(int16))
					case reflect.TypeOf(int32(1)): stack.Push(first.(int32) + last.(int32))
					case reflect.TypeOf(int64(1)): stack.Push(first.(int64) + last.(int64))
					case reflect.TypeOf(uint8(1)): stack.Push(first.(uint8) + last.(uint8))
					case reflect.TypeOf(uint16(1)): stack.Push(first.(uint16) + last.(uint16))
					case reflect.TypeOf(uint32(1)): stack.Push(first.(uint32) + last.(uint32))
					case reflect.TypeOf(uint64(1)): stack.Push(first.(uint64) + last.(uint64))
					case reflect.TypeOf(float32(1)): stack.Push(first.(float32) + last.(float32))
					case reflect.TypeOf(float64(1)): stack.Push(first.(float64) + last.(float64))
					case reflect.TypeOf(string("")):  stack.Push(first.(string) + last.(string))
					}
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) + last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) + last.(float64))
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float64) + last.(float64))
					}
				}
			}
		case "sub":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(first.(int8) - last.(int8))
					case reflect.TypeOf(int16(1)): stack.Push(first.(int16) - last.(int16))
					case reflect.TypeOf(int32(1)): stack.Push(first.(int32) - last.(int32))
					case reflect.TypeOf(int64(1)): stack.Push(first.(int64) - last.(int64))
					case reflect.TypeOf(uint8(1)): stack.Push(first.(uint8) - last.(uint8))
					case reflect.TypeOf(uint16(1)): stack.Push(first.(uint16) - last.(uint16))
					case reflect.TypeOf(uint32(1)): stack.Push(first.(uint32) - last.(uint32))
					case reflect.TypeOf(uint64(1)): stack.Push(first.(uint64) - last.(uint64))
					case reflect.TypeOf(float32(1)): stack.Push(first.(float32) - last.(float32))
					case reflect.TypeOf(float64(1)): stack.Push(first.(float64) - last.(float64))
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) - last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) - last.(float64))
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float64) - last.(float64))
					}
				}
			}
		case "mul":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(first.(int8) * last.(int8))
					case reflect.TypeOf(int16(1)): stack.Push(first.(int16) * last.(int16))
					case reflect.TypeOf(int32(1)): stack.Push(first.(int32) * last.(int32))
					case reflect.TypeOf(int64(1)): stack.Push(first.(int64) * last.(int64))
					case reflect.TypeOf(uint8(1)): stack.Push(first.(uint8) * last.(uint8))
					case reflect.TypeOf(uint16(1)): stack.Push(first.(uint16) * last.(uint16))
					case reflect.TypeOf(uint32(1)): stack.Push(first.(uint32) * last.(uint32))
					case reflect.TypeOf(uint64(1)): stack.Push(first.(uint64) * last.(uint64))
					case reflect.TypeOf(float32(1)): stack.Push(first.(float32) * last.(float32))
					case reflect.TypeOf(float64(1)): stack.Push(first.(float64) * last.(float64))
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) * last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) * last.(float64))
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float64) * last.(float64))
					}
				}
			}
		case "div":
			last, _ := stack.Pop()
			first, _ := stack.Pop()
			if reflect.TypeOf(first) == reflect.TypeOf(last) {
				switch reflect.TypeOf(first) {
					case reflect.TypeOf(int8(1)): stack.Push(first.(int8) / last.(int8))
					case reflect.TypeOf(int16(1)): stack.Push(first.(int16) / last.(int16))
					case reflect.TypeOf(int32(1)): stack.Push(first.(int32) / last.(int32))
					case reflect.TypeOf(int64(1)): stack.Push(first.(int64) / last.(int64))
					case reflect.TypeOf(uint8(1)): stack.Push(first.(uint8) / last.(uint8))
					case reflect.TypeOf(uint16(1)): stack.Push(first.(uint16) / last.(uint16))
					case reflect.TypeOf(uint32(1)): stack.Push(first.(uint32) / last.(uint32))
					case reflect.TypeOf(uint64(1)): stack.Push(first.(uint64) / last.(uint64))
					case reflect.TypeOf(float32(1)): stack.Push(first.(float32) / last.(float32))
					case reflect.TypeOf(float64(1)): stack.Push(first.(float64) / last.(float64))
				}
			} else {
				switch reflect.TypeOf(first) {
				case reflect.TypeOf(int8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					}
				case reflect.TypeOf(int16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					}
				case reflect.TypeOf(int32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					}
				case reflect.TypeOf(int64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					}
				case reflect.TypeOf(uint8(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					}
				case reflect.TypeOf(uint16(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					}
				case reflect.TypeOf(uint32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					}
				case reflect.TypeOf(uint64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float32) / last.(float32))
					} else if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					}
				case reflect.TypeOf(float32(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float64(1)) {
						stack.Push(first.(float64) / last.(float64))
					}
				case reflect.TypeOf(float64(1)):
					if reflect.TypeOf(last) == reflect.TypeOf(float32(1)) {
						stack.Push(first.(float64) / last.(float64))
					}
				}
			}
		case "push":
			rawValue := strings.Join(op[1:len(op)], " ")
			flag := false
			for _, s := range rawValue {
				if string(s) == `"` { //string
					stack.Push(regexp.MustCompile(`"`).ReplaceAllString(rawValue, ``))
					flag = true
					break
				} else if string(s) == `[` { //array //error occur
					newReplacer := strings.NewReplacer("[", "", "]", "")
					replaced := newReplacer.Replace(rawValue)
					arr := strings.Split(replaced, ", ")
					Value := make([]interface{}, 0)
					for _, elem := range arr {
						for _, elemStr := range elem {
							if string(elemStr) == `"` { //string
								Value = append(Value, regexp.MustCompile(`"`).ReplaceAllString(elem, ``))
								break
							} else if string(elemStr) == `.` { //float
								tmp, _ := strconv.ParseFloat(elem, 64)
								Value = append(Value, tmp)
								break
							} else { //int
								tmp, _ := strconv.ParseInt(elem, 10, 32)
								Value = append(Value, tmp)
								break
							}
						}
					}
					stack.Push(Value)
					flag = true
					break
				} else if string(s) == `.` { //float
					tmp, _ := strconv.ParseFloat(rawValue, 64)
					stack.Push(tmp)
					flag = true
					break
				}
			}
			if flag == false {
				tmp, _ := strconv.ParseInt(rawValue, 10, 32)
				stack.Push(tmp)
			}
		case "lds":
			//
		case "load":
			//
		case "store":
			//pop -> type assertion -> push

		}

	}
	if DEBUG {
		fmt.Println(stack.stack)
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

func (s *Stack) Push(data interface{}) {
	s.stack = append(s.stack, data)
}

func (s *Stack) Pop() (interface{}, error) {
	var err error = nil
	if s.IsEmpty() {
		err = errNothingToPop
	}

	poppedValue := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return poppedValue, err

}

func (s Stack) Top() (interface{}, error) {
	var err error = nil
	if s.IsEmpty() {
		err = errStackEmpty
	}
	top := len(s.stack) - 1
	return s.stack[top], err
}
