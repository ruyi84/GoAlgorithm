package main

func canThreePartsEqualSum(A []int) bool {

	sum := 0

	for _, v := range A {
		sum += v
	}

	if sum%3 != 0 {
		return false
	}

	num1 := 0
	num2 := 0

	for _, v := range A {
		num1 += v
		if num1 == sum/3 {
			num2++
			num1 = 0
		}
	}

	if num2 >= 3 {
		return true
	} else {
		return false
	}
}
