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

	var validLists [][]int
	for _, list := range listOfLists {
		if isValid(list, isAfter) {
			validLists = append(validLists, list)
		}
	}

	middleTotal := 0
	for _, list := range validLists {
		middle := len(list) / 2
		middleTotal += list[middle]
	}

	fmt.Println(middleTotal)
}

func isValid(list []int, isAfter map[int][]int) bool {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			afterList, _ := isAfter[list[j]]
			if !slices.Contains(afterList, list[i]) {
				return false
			}
		}
	}

	return true
}
