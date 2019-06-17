// Write a program that takes an arithmetical expression in RPN and returns
// the number that the expression evaluates to.

package main

import (
	"fmt"
	"strconv"
	"strings"
	"./stack"
)

type Stack = stack.Stack
var InitStack = stack.InitStack

func isOperator(s string) bool {
	if s == "*" || s == "/" || s == "+" || s == "-" {
		return true
	}

	return false
}

func evaluate(right, left int, operator string) int {
	if operator == "*" {
		return left * right
	} else if operator == "/" {
		return left / right
	} else if operator == "+" {
		return left + right
	} else { // operator == "-"
		return left - right
	}
}

func EvaluateReversePolish(expression, delimiter string) int {
	operandStack:= InitStack()

	tokens := strings.Split(expression, delimiter)
	for i := 0; i < len(tokens); i++ {
		if isOperator(tokens[i]) {
			right := operandStack.Pop().(int)
			left := operandStack.Pop().(int)
			operandStack.Push(evaluate(right, left, tokens[i]))
		} else {
			data, _ := strconv.Atoi(tokens[i])
			operandStack.Push(data)
		}
	}

	return  operandStack.Pop().(int)
}

func EvaluatePolish(expression, delimiter string) int {
	stack:= InitStack()

	tokens := strings.Split(expression, delimiter)
	for i := 0; i < len(tokens); i++ {
		// Push the operator.
		if isOperator(tokens[i]) {
			stack.Push(tokens[i])
			continue
		}

		// If found two operands in a row, then evaluate.
		right, _ := strconv.Atoi(tokens[i])

		var data interface{}
		var left int
		var ok bool

		if !stack.Empty() {
			data = stack.Peek()
			left, ok = data.(int)
		}

		for ok {
			// If the top of the stack is an operand, then pop,
			// an operand, pop an operator, evaluate, and check
			// again.
			stack.Pop()
			operator := stack.Pop().(string)
			right = evaluate(right, left, operator)

			if stack.Empty() {
				ok = false
			} else {
				data = stack.Peek()
				left, ok = data.(int)
			}
		}

		// Push the result/token to the stack.
		stack.Push(right)
	}

	// Pop the result.
	return stack.Pop().(int)
}

func main() {
	// 1. (3 + 4) * 2 + 1
	// 2. (1 + 1) * -2
	// 3. -641 / 6 / 28
	// 4. 3 + 4 * (2 + 3) - 1

	for _, expr := range []string{"3,4,+,2,*,1,+", "1,1,+,-2,*",
		"-641,6,/,28,/", "3,4,2,3,+,*,+,1,-" } {
		fmt.Println(expr, "=", EvaluateReversePolish(expr, ","))
	}

	fmt.Println()

	for _, expr := range []string{"+,*,+,3,4,2,1", "*,+,1,1,-2",
		"/,/,-641,6,28", "-,+,3,*,4,+,2,3,1" } {
		fmt.Println(expr, "=", EvaluatePolish(expr, ","))
	}
}
