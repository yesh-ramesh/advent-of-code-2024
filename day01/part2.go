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
	file, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse input and save to slices
	var list1 []int
	var list2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)

		num1, _ := strconv.Atoi(split[0])
		num2, _ := strconv.Atoi(split[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	// Populate frequency where key = <number in list1> and value = <number of times its in list2>
	freqMap := make(map[int]int)
	totalSimilarityScore := 0
	for _, v1 := range list1 {
		for _, v2 := range list2 {
			if v2 == v1 {
				freqMap[v1]++
			}
		}
	}

	// Calculate similarity score
	for key, val := range freqMap {
		similarityScore := key * val
		totalSimilarityScore += similarityScore
	}

	fmt.Println(totalSimilarityScore)
}
