package main

type area struct {
	height, index int
}

func largestRectangleArea(heights []int) int {
	stack := []area{}
	max := -1
	for i, h := range heights {
		tempI := i
		for len(stack) > 0 && stack[len(stack)-1].height > h {
			top := stack[len(stack)-1]
			weight := i - top.index
			if top.height*weight > max {
				max = top.height * weight
			}
			tempI = top.index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, area{h, tempI})
	}
	for _, v := range stack {
		if v.height*(len(heights)-v.index) > max {
			max = v.height * (len(heights) - v.index)
		}
	}
	return max
}
