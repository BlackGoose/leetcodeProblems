package main

import "slices"

func findTime(position, speed, target int) float32 {
	return float32(target-position) / float32(speed)
}

func carFleet(target int, position []int, speed []int) int {
	index := make([]int, len(position))
	stack := []int{}
	for i := range position {
		index[i] = i
	}
	slices.SortFunc(index, func(i, j int) int {
		return position[j] - position[i]
	})
	for _, i := range index {
		if len(stack) == 0 ||
			findTime(position[i], speed[i], target) > findTime(position[stack[len(stack)-1]], speed[stack[len(stack)-1]], target) {
			stack = append(stack, i)
		}
	}
	return len(stack)
}
