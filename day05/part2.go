package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	isAfter := make(map[int][]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, "|")

		if len(split) != 2 {
			break
		}

		first, _ := strconv.Atoi(split[0])
		second, _ := strconv.Atoi(split[1])

		if _, ok := isAfter[second]; !ok {
			isAfter[second] = make([]int, 0)
		}

		isAfter[second] = append(isAfter[second], first)
	}

	var listOfLists [][]int
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		var list []int

		for _, v := range split {
			i, _ := strconv.Atoi(v)
			list = append(list, i)
		}
		listOfLists = append(listOfLists, list)
	}

	var invalidLists [][]int
	for _, list := range listOfLists {
		if isValid(list, isAfter) {
			invalidLists = append(invalidLists, list)
		}
	}

	middleTotal := 0
	for _, list := range invalidLists {
		ordered := order(list, isAfter)
		middle := len(ordered) / 2
		middleTotal += ordered[middle]
	}
	fmt.Println(middleTotal)
}

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
