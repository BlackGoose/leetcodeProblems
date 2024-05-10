package main

func isValid(s string) bool {
	parenthesesDict := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}
	for _, w := range s {
		if w == '(' || w == '{' || w == '[' {
			stack = append(stack, w)
		} else {
			if len(stack) > 0 && w == parenthesesDict[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
