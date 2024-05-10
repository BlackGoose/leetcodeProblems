package main

type tempStack struct {
	index int
	value int
}

func dailyTemperatures(temperatures []int) []int {
	stack := []tempStack{}
	result := make([]int, len((temperatures)))

	for i, t := range temperatures {
		for len(stack) > 0 && stack[len(stack)-1].value < t {
			result[stack[len(stack)-1].index] = i - stack[len(stack)-1].index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, tempStack{i, t})
	}
	return result
}
