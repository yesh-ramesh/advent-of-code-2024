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

	// Check safety of each report
	safeCount := 0
	for _, report := range reports {
		// Assume the report is safe
		safe := true

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
			if (nextLevel > currentLevel && !increasing) || (nextLevel < currentLevel && increasing) {
				safe = false
				break
			}

			// If the difference between levels is less than 1 or greater than 3, we're unsafe
			difference := math.Abs(float64(currentLevel) - float64(nextLevel))
			if difference < 1 || difference > 3 {
				safe = false
				break
			}
		}

		if safe {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}
