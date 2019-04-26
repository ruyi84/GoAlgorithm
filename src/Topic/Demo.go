package main

import "fmt"

func main() {
	f1()
}

func f1() {
	sli := []int{1, 2, 3, 4, 5}
	str := 2
	for k2, v2 := range sli {
		if str == v2 {
			sli = sli[k2+1:]
			break
		}
	}
	fmt.Println(sli)
}

func sli() {
	sli := make([]string, 3)
	sli[0] = "0"
	sli[1] = "1"
	sli[2] = "2"
	fmt.Printf("%v\n", sli)
	sli = append(sli, "3")
	fmt.Printf("%v\n", sli)
	sli = sli[1:]
	fmt.Printf("%v\n", sli)
	sli = sli[1:]
	sli = append(sli, "3")
	fmt.Printf("%v\n", sli)
}
