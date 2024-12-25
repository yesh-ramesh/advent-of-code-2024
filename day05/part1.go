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

	scanner := bufio.NewScanner(file)

	// Read the first chunk of input and create a map to keep track of what comes after what
	// after is a key-value map that gives me a slice of all numbers that come after a key
	after := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, "|")

		if len(split) != 2 {
			break
		}

		first, _ := strconv.Atoi(split[0])
		second, _ := strconv.Atoi(split[1])

		if _, ok := after[second]; !ok {
			after[second] = make([]int, 0)
		}

		after[second] = append(after[second], first)
	}

	// Read the second chunk of input to get our list of numbers
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

	// Use the map of ordering to determine valid, ordered lists
	var validLists [][]int
	for _, list := range listOfLists {
		if isValid(list, after) {
			validLists = append(validLists, list)
		}
	}

	// Sum up the middle value of the valid, ordered lists
	middleTotal := 0
	for _, list := range validLists {
		middle := len(list) / 2
		middleTotal += list[middle]
	}

	fmt.Println(middleTotal)
}

func isValid(list []int, after map[int][]int) bool {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if !slices.Contains(after[list[j]], list[i]) {
				return false
			}
		}
	}

	return true
}
