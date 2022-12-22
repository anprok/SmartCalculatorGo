package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	variables = make(map[string]int)
	for {
		userInput := readString()
		if userInput == "" {
			continue
		} else if strings.HasPrefix(userInput, "/") {
			commandOperation(userInput)
		} else if strings.Contains(userInput, "=") {
			assignmentOperation(userInput)
		} else if strings.ContainsAny(userInput, "+-*/") {
			arithmeticOperation(userInput)
		} else if len(strings.Fields(userInput)) == 1 {
			printValue(userInput)
		}
	}
}

var variables map[string]int

func printValue(input string) {
	if IsLetter(input) {
		if val, ok := variables[input]; ok {
			fmt.Println(val)
		} else {
			fmt.Println("Unknown variable")
		}
	} else {
		fmt.Println("Invalid identifier")
	}
}

func arithmeticOperation(input string) {
	inpFields := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(input, ")", " )"), "(", "( "))
	var (
		num               int
		postfixExpression []string
		operators         Stack
		calculation       StackInt
	)
	if !validate(inpFields) {
		fmt.Println("Invalid expression")
		return
	}
	// String processing
	for _, field := range inpFields {
		if IsInt(field) {
			//num, _ = strconv.Atoi(strings.TrimSpace(field))
			//nums = append(nums, num)
			postfixExpression = append(postfixExpression, strings.TrimSpace(field))
		} else if IsLetter(strings.TrimSpace(field)) {
			//nums = append(nums, variables[strings.TrimSpace(field)])
			postfixExpression = append(postfixExpression, strings.TrimSpace(field))
		} else if field[0] == '+' || (field[0] == '-' && len(field)%2 == 0) {
			for {
				char, err := operators.Pop()
				// If the stack is empty or contains a left parenthesis on top, push the incoming operator on the stack.
				if err != nil || char == "(" {
					if char != "" {
						operators.Push(char)
					}
					operators.Push("+")
					// Stop processing  stack
					break
				} else {
					if char == "(" {
						operators.Push(char)
						operators.Push("+")
						break
					} else {
						postfixExpression = append(postfixExpression, char)
					}
				}
			}
		} else if field[0] == '-' {
			for {
				char, err := operators.Pop()
				// If the stack is empty or contains a left parenthesis on top, push the incoming operator on the stack.
				if err != nil || char == "(" {
					if char != "" {
						operators.Push(char)
					}
					operators.Push("-")
					// Stop processing  stack
					break
				} else {
					if char == "(" {
						operators.Push(char)
						operators.Push(string(field[0]))
						break
					} else {
						postfixExpression = append(postfixExpression, char)
					}
				}
			}

		} else if field[0] == '*' || field[0] == '/' {
			for {
				char, err := operators.Pop()
				// If the stack is empty or contains a left parenthesis on top,
				// push the incoming operator on the stack.
				// If the incoming operator has higher precedence than the top of the stack,
				// push it on the stack.
				if err != nil || char == "(" || char == "+" || char == "-" {
					if char != "" {
						operators.Push(char)
					}
					operators.Push(string(field[0]))
					// Stop processing  stack
					break
				} else {
					// If the precedence of the incoming operator is
					// equal to that of the top of the stack, pop the stack and add operators
					// to the result until you see an operator that has smaller precedence
					// or a left parenthesis on the top of the stack;
					// then add the incoming operator to the stack.
					if char == "(" || char == "+" || char == "-" {
						operators.Push(char)
						operators.Push(string(field[0]))
						break
					} else {
						postfixExpression = append(postfixExpression, char)
					}
				}
			}
		} else if field[0] == '(' {
			operators.Push(string(field[0]))
		} else if field[0] == ')' {
			for {
				char, err := operators.Pop()
				if err != nil {
					break
				}
				if char == "(" {
					break
				}
				postfixExpression = append(postfixExpression, char)
			}
		}
	}
	for {
		char, err := operators.Pop()
		if err != nil {
			break
		}
		if char == "(" || char == ")" {
			fmt.Println("Invalid expression")
			return
		}
		postfixExpression = append(postfixExpression, char)
	}
	for _, element := range postfixExpression {
		if IsInt(element) {
			num, _ = strconv.Atoi(strings.TrimSpace(element))
			calculation.Push(num)
		} else if IsLetter(element) {
			calculation.Push(variables[strings.TrimSpace(element)])
		} else {
			number1, _ := calculation.Pop()
			number2, _ := calculation.Pop()
			if element == "+" {
				calculation.Push(number1 + number2)
			} else if element == "-" {
				calculation.Push(number2 - number1)
			} else if element == "*" {
				calculation.Push(number2 * number1)
			} else if element == "/" {
				calculation.Push(number2 / number1)
			}
		}
	}
	result, _ := calculation.Pop()
	fmt.Println(result)
}

func validate(fields []string) bool {
	var solver Queue

	for _, fld := range fields {
		if fld == "(" {
			solver.Push(fld)
		}
		if fld == ")" {
			_, err := solver.Pop()
			if err != nil {
				return false
			}
		}
		if (strings.HasSuffix(fld, "*") || strings.HasSuffix(fld, "/")) && len(fld) > 1 {
			return false
		}
	}

	_, err := solver.Pop()
	if err != nil {
		return true
	} else {
		return false
	}
}

func assignmentOperation(input string) {
	subs := strings.Split(input, "=")
	if !IsLetter(strings.TrimSpace(subs[0])) {
		fmt.Println("Invalid identifier")
		return
	}
	if !IsLetter(strings.TrimSpace(subs[1])) && !IsInt(strings.TrimSpace(subs[1])) {
		fmt.Println("Invalid assignment")
		return
	}
	num, err := strconv.Atoi(strings.TrimSpace(subs[1]))
	if err == nil {
		variables[strings.TrimSpace(subs[0])] = num
	} else {
		if val, ok := variables[strings.TrimSpace(subs[1])]; ok {
			variables[strings.TrimSpace(subs[0])] = val
		} else {
			fmt.Println("Unknown variable")
		}
	}
}

func commandOperation(input string) {
	if input == "/exit" {
		fmt.Println("Bye!")
		os.Exit(0)
	} else if input == "/help" {
		fmt.Println("The program calculates the sum or substraction of numbers and variables. Supports  both unary and binary operators. Variable name can contain only Latin letters.Variable value can be integer numbero or another existing variable.")
	} else {
		fmt.Println("Unknown command")
	}
}

func readString() string {
	var userInput string
	reader := bufio.NewReader(os.Stdin)
	userInput, _ = reader.ReadString('\n')
	return strings.TrimSpace(userInput)
}

func IsInt(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func IsLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

type Stack struct {
	storage []string
}

func (s *Stack) Push(value string) {
	s.storage = append(s.storage, value)
}

func (s *Stack) Pop() (string, error) {
	last := len(s.storage) - 1
	if last <= -1 { // check the size
		return "", errors.New("Stack is empty") // and return error
	}
	value := s.storage[last]     // save the value
	s.storage = s.storage[:last] // remove the last element
	return value, nil            // return saved value and nil error
}

type StackInt struct {
	storage []int
}

func (s *StackInt) Push(value int) {
	s.storage = append(s.storage, value)
}

func (s *StackInt) Pop() (int, error) {
	last := len(s.storage) - 1
	if last <= -1 { // check the size
		return 0, errors.New("Stack is empty") // and return error
	}
	value := s.storage[last]     // save the value
	s.storage = s.storage[:last] // remove the last element
	return value, nil            // return saved value and nil error
}

type Queue struct {
	input  Stack
	output Stack
}

func (q *Queue) Push(value string) {
	q.input.Push(value)
}

func (q *Queue) Pop() (string, error) {
	outpuVal, outputErr := q.output.Pop()
	if outputErr == nil { // if output stack is not empty
		return outpuVal, nil // just return value
	}

	inputVal, inputErr := q.input.Pop()
	if inputErr != nil { // if input stack is empty
		return string(0), errors.New("Queue is empty") // return the error
	}

	// if the output stack is empty but the input is not empty
	for inputErr == nil { // while input stack not empty...
		q.output.Push(inputVal)            // rearrange input to output
		inputVal, inputErr = q.input.Pop() // and read again
	}

	return q.output.Pop() // and Pop the output
}
