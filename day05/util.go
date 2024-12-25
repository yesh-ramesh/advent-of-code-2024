package main

import "slices"

// isValid returns true if a list is in the correct order.
// Order is determined via the after map, whose key is a number and whose value is a list of numbers that come after the key.
// For example, after[1] = { 2, 4, 5 } means 1 must come after 2, 4, and 5
func isValid(list []int, after map[int][]int) bool {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if !slices.Contains(after[list[j]], list[i]) {
				return true
			}
		}
	}

	return false
}

// order returns an ordered version of a list given the after map, whose key is a number and whose value is a list of numbers that come after the key.
// For example, after[1] = { 2, 4, 5 } means 1 must come after 2, 4, and 5
func order(unordered []int, after map[int][]int) []int {
	for i := 0; i < len(unordered); i++ {
		for j := i + 1; j < len(unordered); j++ {
			if !slices.Contains(after[unordered[j]], unordered[i]) {
				temp := unordered[i]
				unordered[i] = unordered[j]
				unordered[j] = temp
			}
		}
	}

	return unordered
}
