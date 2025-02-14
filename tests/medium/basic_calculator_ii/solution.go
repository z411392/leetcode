package basic_calculator_ii

import (
	"strconv"
)

type Stack []interface{}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *Stack) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

func (s *Stack) Top() interface{} {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// 檢查運算子優先級
func getPrecedence(op byte) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

// 將中綴表達式轉換為後綴表達式
func infixToRPN(exp string) []string {
	var output []string
	var operators Stack

	i := 0
	for i < len(exp) {
		// 跳過空格
		if exp[i] == ' ' {
			i++
			continue
		}

		// 處理數字
		if isDigit(exp[i]) {
			num := ""
			for i < len(exp) && isDigit(exp[i]) {
				num += string(exp[i])
				i++
			}
			output = append(output, num)
			// fmt.Printf("num=%v, outputs=%v\n", num, output)
			continue
		}

		if !isOperator(exp[i]) {
			i++
			continue
		}

		// 處理運算子
		for {
			if operators.IsEmpty() {
				break
			}
			top := operators.Top().(byte)
			if getPrecedence(top) < getPrecedence(exp[i]) {
				break
			}
			operator := string(operators.Pop().(byte))
			output = append(output, operator)
			// fmt.Printf("operator=%v, outputs=%v\n", operator, output)
		}
		operators.Push(exp[i])
		// fmt.Printf("operator=%c\n", exp[i])
		i++
	}

	// 處理剩餘的運算子
	for !operators.IsEmpty() {
		operator := string(operators.Pop().(byte))
		output = append(output, operator)
		// fmt.Printf("operator=%v, outputs=%v\n", operator, output)
	}

	return output
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isOperator(c byte) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

// 計算 RPN 表達式的結果
func evaluateRPN(tokens []string) int {
	var stack Stack

	for _, token := range tokens {
		if !isOperator(token[0]) {
			num, _ := strconv.Atoi(token)
			stack.Push(num)
			continue
		}

		// 處理運算子
		b := stack.Pop().(int)
		a := stack.Pop().(int)

		switch token {
		case "+":
			stack.Push(a + b)
		case "-":
			stack.Push(a - b)
		case "*":
			stack.Push(a * b)
		case "/":
			stack.Push(a / b)
		}
	}

	return stack.Pop().(int)
}

/*
Given a string s which represents an expression, evaluate this expression and return its value.

The integer division should truncate toward zero.

You may assume that the given expression is always valid. All intermediate results will be in the range of [-231, 231 - 1].

Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().
*/
func calculate(s string) int {
	rpn := infixToRPN(s)
	return evaluateRPN(rpn)
}
