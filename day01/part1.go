package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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

	// Sort slices
	slices.Sort(list1)
	slices.Sort(list2)

	// Pair up the numbers and sum their differences
	totalDiff := 0
	for i, v1 := range list1 {
		v2 := list2[i]

		diff := math.Abs(float64(v1) - float64(v2))
		totalDiff += int(diff)
	}

	fmt.Println(totalDiff)
}
