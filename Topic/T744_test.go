package main

import (
	"testing"
)

func nextGreatestLetter(letters []byte, target byte) byte {

	left, right := 0, len(letters)-1

	if target >= letters[right] || target < letters[left] {
		return letters[0]
	}

	for left < right {
		mid := left + (right-left)/2

		if target >= letters[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if letters[right] <= target {
		return letters[right+1]
	}

	return letters[right]
}

func nextGreatestLetter1(letters []byte, target byte) byte {

	left, right := 0, len(letters)-1

	if target >= letters[right] || target < letters[left] {
		return letters[0]
	}

	result := 0
	for left <= right {
		mid := left + (right-left)/2

		if target >= letters[mid] {
			result = right
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	//if letters[right] <= target {
	//	return letters[right+1]
	//}

	return letters[result]
}

func TestT744(t *testing.T) {
	t.Log(nextGreatestLetter1([]byte{'c', 'f', 'j'}, 'a'))
	t.Log(nextGreatestLetter1([]byte{'c', 'f', 'j'}, 'd'))
	t.Log(nextGreatestLetter1([]byte{'c', 'f', 'j'}, 'c'))
	t.Log(nextGreatestLetter1([]byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'}, 'e'))
}
