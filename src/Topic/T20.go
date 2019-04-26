package main

//未通过
import "fmt"

func isValid(s string) bool {
	bra := map[string]string{")": "(", "]": "[", "}": "{"}

	sli := make([]string, 0)
	index := -1
	var str string

	for k := 0; k < (len(s)); k++ {
		str = string(s[k])
		if str == "(" || str == "{" || str == "[" {
			index = len(sli)
			sli = append(sli, str)
		} else if index >= 0 && bra[str] == sli[index] {
			sli = sli[:index]
			index -= 1
		} else {
			return false
		}

	}
	if len(sli) == 0 {
		return true
	} else {
		return false
	}

}

func main() {

	fmt.Println(isValid("()"))
}
