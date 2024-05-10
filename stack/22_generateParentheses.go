package main

import "strings"

type paranthesisStack struct {
	stack []string
	open  int
	close int
}

func generateParenthesis(n int) []string {

	result := []string{}
	ps := paranthesisStack{stack: []string{}, open: 0, close: 0}

	var recurs func(ps paranthesisStack)

	recurs = func(ps paranthesisStack) {

		if ps.close == ps.open && ps.close == n {
			result = append(result, strings.Join(ps.stack, ""))
		}

		if ps.open < n {
			ps.stack = append(ps.stack, "(")
			ps.open++
			recurs(ps)
			ps.stack = ps.stack[:len(ps.stack)-1]
			ps.open--
		}

		if ps.close < ps.open {
			ps.stack = append(ps.stack, ")")
			ps.close++
			recurs(ps)
		}
	}

	recurs(ps)

	return result
}
