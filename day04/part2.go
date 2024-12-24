package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Build grid
	gridSize := 140
	grid := make([][]rune, gridSize)
	for i := range grid {
		grid[i] = make([]rune, gridSize)
	}

	index := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		grid[index] = runes
		index++
	}

	// Check if the sliding window matches the possibilities
	count := 0
	for i := 0; i <= gridSize-3; i++ {
		for j := 0; j <= gridSize-3; j++ {
			// Populate a sliding grid
			slider := make([][]rune, 3)
			for i := range slider {
				slider[i] = make([]rune, 3)
			}

			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					slider[x][y] = grid[i+x][j+y]
				}
			}

			// Check if the sliding grid forms a X-MAS
			if isMatch(slider) {
				count++
			}
		}
	}

	fmt.Println(count)
}

/*
Check if the sliding window forms one of these configurations:

M . M
. A .
S . S

S . S
. A .
M . M

M . S
. A .
M . S

S . M
. A .
S . M
*/
func isMatch(slider [][]rune) bool {
	if slider[1][1] != 'A' {
		return false
	}

	if slider[0][0] == 'M' && slider[0][2] == 'M' && slider[2][0] == 'S' && slider[2][2] == 'S' {
		return true
	} else if slider[0][0] == 'M' && slider[0][2] == 'S' && slider[2][0] == 'M' && slider[2][2] == 'S' {
		return true
	} else if slider[0][0] == 'S' && slider[0][2] == 'S' && slider[2][0] == 'M' && slider[2][2] == 'M' {
		return true
	} else if slider[0][0] == 'S' && slider[0][2] == 'M' && slider[2][0] == 'S' && slider[2][2] == 'M' {
		return true
	}

	return false
}
