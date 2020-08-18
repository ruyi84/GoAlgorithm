package main

import (
	"fmt"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	str := ""

	for k := range strs[0] {
		keyWord := strs[0][k]

		for i := 1; i < len(strs); i++ {
			if len(strs[i]) <= k {
				return str
			}
			if strs[i][k] == keyWord {
				continue
			}
			return str
		}
		str += string(keyWord)
	}
	return str
}

func Test_T014(t *testing.T) {
	fmt.Println(longestCommonPrefix([]string{}))
}
