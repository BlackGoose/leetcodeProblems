package main

import "strconv"

func evaluate(op string, num1, num2 int) int {
	switch {
	case op == "+":
		return num1 + num2
	case op == "*":
		return num1 * num2
	case op == "/":
		return num1 / num2
	default:
		return num1 - num2
	}
}

func evalRPN(tokens []string) int {
	num_stack := []int{}
	for _, t := range tokens {
		num, err := strconv.Atoi(t)
		if err != nil {
			num2 := num_stack[len(num_stack)-1]
			num1 := num_stack[len(num_stack)-2]
			num_stack = num_stack[:len(num_stack)-2]
			num_stack = append(num_stack, evaluate(t, num1, num2))
		} else {
			num_stack = append(num_stack, num)
		}
	}
	return num_stack[0]
}
