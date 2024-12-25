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

	// Use the after map to determine valid, ordered lists
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
