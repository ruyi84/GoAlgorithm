package main

//失败
import "fmt"

func isValid2(s string) bool {
	bra := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	var str string
	left := make([]string, 0)
	right := make([]string, 0)

	for i := 0; i < (len(s)); i++ {
		str = string(s[i])
		if str == "(" || str == "[" || str == "{" {
			left = append(left, str)
		} else if str == ")" || str == "]" || str == "}" {
			right = append(right, str)
		} else {
			return false
		}
	}

	if len(left) != len(right) {
		return false
	}
	for k, _ := range left {
		if left[k] != bra[right[len(right)-k-1]] {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(isValid2("()[]{}"))
}
