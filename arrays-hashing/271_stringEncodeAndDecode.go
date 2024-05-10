package main

import (
	"fmt"
	"strconv"
	"strings"
)

func encode(strs []string) string {
	result := strings.Builder{}
	result.Grow(len(strs))

	for _, s := range strs {
		result.Write([]byte(fmt.Sprintf("%v:%v", len(s), s)))
	}
	return result.String()
}

func decode(s string) []string {
	result := []string{}

	for i := 0; i < len(s); {
		shift := strings.Index(s[i:], ":")
		numero, _ := strconv.Atoi(s[i : i+shift])
		result = append(result, s[i+shift+1:i+shift+1+numero])
		i += shift + 1 + numero
	}
	return result
}
