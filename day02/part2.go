package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Parse and save input
	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report := scanner.Text()
		levelStrings := strings.Fields(report)

		var levels []int
		for _, levelStr := range levelStrings {
			level, _ := strconv.Atoi(levelStr)
			levels = append(levels, level)
		}

		reports = append(reports, levels)
	}

	safeCount := 0
	for _, report := range reports {
		if isSafe(report) {
			safeCount++
		} else {
			// Check every possible report where one level is removed
			safeWithTolerance := false
			for i := 0; i < len(report); i++ {
				newReport := append(append([]int{}, report[:i]...), report[i+1:]...)
				if isSafe(newReport) {
					safeWithTolerance = true
					break
				}
			}

			if safeWithTolerance {
				safeCount++
			}
		}
	}

	fmt.Println(safeCount)
}

func isSafe(report []int) bool {
	// Check if all values should be increasing or decreasing by comparing the first two levels
	increasing := false
	if report[1] > report[0] {
		increasing = true
	}

	// Loop through each level in the report
	for i := 0; i < len(report)-1; i++ {
		// Get the current and next level for comparison
		currentLevel := report[i]
		nextLevel := report[i+1]

		// If the next level is greater than the current level, but the report should be decreasing, we're unsafe
		// If the next level is less than the current level, but the report should be increasing, we're unsafe
		// If the difference between levels is less than 1 or greater than 3, we're unsafe
		difference := math.Abs(float64(currentLevel) - float64(nextLevel))
		if ((nextLevel > currentLevel && !increasing) || (nextLevel < currentLevel && increasing)) || (difference < 1 || difference > 3) {
			return false
		}
	}

	return true
}
