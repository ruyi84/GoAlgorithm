package main

// 直接双重for循环暴力解决问题
func findTheDistanceValue1(arr1 []int, arr2 []int, d int) int {
	num := len(arr1)
	for _, v := range arr1 {
		for _, v2 := range arr2 {
			diff := v - v2
			if (diff > 0 && diff > d) || (diff<0 && diff < 0 - d){
				continue
			}
			num--
			break
		}
	}

	return num
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {

}
