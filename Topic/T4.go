package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	num1 := 0

	for i := 0; i < len(nums2); i++ {
		for num1 < len(nums1) {
			if nums2[i] > nums1[num1] {
				if num1+1 > len(nums1) {
					nums1 = append(nums1, nums2[i])
					num1++
					continue
				}
				nums1 = append(nums1[:num1+1], nums2[i], nums1[num1+1])
				num1++
				continue
			}

			num1++
		}
	}

	if len(nums1)%2 == 0 {
		return float64((nums1[len(nums1)%2/2] + nums1[len(nums1)%22-1]) / 2)
	} else {
		return float64(nums1[len(nums1)/2])
	}
}
