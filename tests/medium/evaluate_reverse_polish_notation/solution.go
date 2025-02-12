package evaluate_reverse_polish_notation

import (
	"fmt"
	"strconv"
)

/*
You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:

The valid operators are '+', '-', '*', and '/'.
Each operand may be an integer or another expression.
The division between two integers always truncates toward zero.
There will not be any division by zero.
The input represents a valid arithmetic expression in a reverse polish notation.
The answer and all the intermediate calculations can be represented in a 32-bit integer.
*/
func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	if len(tokens) == 1 {
		n, _ := strconv.Atoi(tokens[0])
		return n
	}
	stack := []int{}
	push := func(n int) {
		stack = append(stack, n)
	}
	pop := func() (int, error) {
		if len(stack) == 0 {
			return 0, fmt.Errorf("")
		}
		value := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return value, nil
	}
	// 先 pop 出來的是第二個運算元
	for _, token := range tokens {
		switch token {
		case "+":
			n2, _ := pop()
			n1, _ := pop()
			n := n1 + n2
			// fmt.Printf("%v%s%v=%v\n", n1, token, n2, n)
			push(n)
		case "-":
			n2, _ := pop()
			n1, _ := pop()
			n := n1 - n2
			// fmt.Printf("%v%s%v=%v\n", n1, token, n2, n)
			push(n)
		case "*":
			n2, _ := pop()
			n1, _ := pop()
			n := n1 * n2
			// fmt.Printf("%v%s%v=%v\n", n1, token, n2, n)
			push(n)
		case "/":
			n2, _ := pop()
			n1, _ := pop()
			n := n1 / n2
			// fmt.Printf("%v%s%v=%v\n", n1, token, n2, n)
			push(n)
		default:
			n, _ := strconv.Atoi(token)
			push(n)
		}
	}
	return stack[0]
}
