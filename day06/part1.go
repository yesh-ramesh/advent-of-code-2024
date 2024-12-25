package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Location struct {
	row, col int
}

func main() {
	file, err := os.Open("day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Build grid
	gridSize := 130
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

	// Save the starting position of the guard and positions of obstacles
	obstacles := make(map[Location]bool)
	visited := make(map[Location]bool)
	guard := Location{row: 0, col: 0}
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] == '#' {
				obstacles[Location{i, j}] = true
			}

			if grid[i][j] == '^' {
				guard = Location{i, j}
			}
		}
	}

	// Move the guard keeping track of the positions visited
	// 0 - up, 1 - right, 2 - down, 3 - left
	direction := 0
	for guard.row > 0 && guard.row < gridSize-1 && guard.col > 0 && guard.col < gridSize-1 {
		switch direction {
		case 0:
			if !obstacles[Location{guard.row - 1, guard.col}] {
				guard.row--
			} else {
				direction = 1
			}
		case 1:
			if !obstacles[Location{guard.row, guard.col + 1}] {
				guard.col++
			} else {
				direction = 2
			}
		case 2:
			if !obstacles[Location{guard.row + 1, guard.col}] {
				guard.row++
			} else {
				direction = 3
			}
		case 3:
			if !obstacles[Location{guard.row, guard.col - 1}] {
				guard.col--
			} else {
				direction = 0
			}
		}

		visited[guard] = true
	}

	fmt.Println(len(visited))
}
