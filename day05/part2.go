package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	// Read the first chunk of input and create a map to keep track of what comes after what
	// after is a key-value map that gives me a slice of all numbers that come after a key
	// For example, after[1] = { 2, 4, 5 } means 1 must come after 2, 4, and 5
	isAfter := make(map[int][]int)
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

	// Read the second chunk of input to get our lists of numbers
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

	// Use the after map to determine invalid, unordered lists
	var invalidLists [][]int
	for _, list := range listOfLists {
		if isValid(list, isAfter) {
			invalidLists = append(invalidLists, list)
		}
	}

	// Put the invalid lists in the correct order and sum their middle values
	middleTotal := 0
	for _, list := range invalidLists {
		ordered := order(list, isAfter)
		middle := len(ordered) / 2
		middleTotal += ordered[middle]
	}
	fmt.Println(middleTotal)
}
