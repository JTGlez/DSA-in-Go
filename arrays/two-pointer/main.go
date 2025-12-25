package main

import "fmt"

func main() {

	result := palindromeFinder("racecar")
	fmt.Printf("Result :%v", result)

	sorted := []int{1, 2, 4, 6, 8, 9, 14, 15}
	resultTwo := twoSumSorted(sorted, 13)
	fmt.Printf("Result :%v", resultTwo)

	arr1 := []int{1, 3, 6, 9}
	arr2 := []int{2, 4, 9, 10}

	merged := twoArrayNewSorted(arr1, arr2)

	fmt.Printf("Merged :%v", merged)

}

/*
Example 1: Given a string s, return true if it is a palindrome, false otherwise.

A string is a palindrome if it reads the same forward as backward. That means, after reversing it,
it is still the same string. For example: "abcdcba", or "racecar".
*/
func palindromeFinder(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

/*
Example 2: Given a sorted array of unique integers and a target integer, return true if there exists
a pair of numbers that sum to target, false otherwise. This problem is similar to Two Sum.
(In Two Sum, the input is not sorted).

For example, given nums = [1, 2, 4, 6, 8, 9, 14, 15] and target = 13, return true because 4 + 9 = 13.
*/

func twoSumSorted(s []int, target int) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left]+s[right] > target {
			right--
			continue
		}

		if s[left]+s[right] < target {
			left++
			continue
		}

		if s[left]+s[right] == target {
			return true
		}

	}

	return false
}

/*
Example 3:
Given two sorted integer arrays arr1 and arr2, return a new array that combines
both of them and is also sorted.
*/

func twoArrayNewSorted(n, m []int) []int {
	i, j := 0, 0
	r := make([]int, 0, len(n)+len(m))

	for i < len(n) && j < len(m) {
		if n[i] < m[j] {
			r = append(r, n[i])
			i++
			continue
		}

		if n[i] > m[j] {
			r = append(r, m[j])
			j++
			continue
		}

		if n[i] == m[j] {
			r = append(r, n[i], m[j])
			i++
			j++
			continue
		}
	}

	if i < len(n) {
		r = append(r, n[i:]...)
	}

	if j < len(m) {
		r = append(r, m[j:]...)
	}

	return r
}
