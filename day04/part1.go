package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var totalCount = 0

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

	// Search every cell and every direction using a DFS word-search algo
	word := []rune("XMAS")
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			wordSearch(i, j, grid, word, 0, -1, 0)
			wordSearch(i, j, grid, word, 0, -1, -1)
			wordSearch(i, j, grid, word, 0, 0, -1)
			wordSearch(i, j, grid, word, 0, 1, -1)
			wordSearch(i, j, grid, word, 0, 1, 0)
			wordSearch(i, j, grid, word, 0, 1, 1)
			wordSearch(i, j, grid, word, 0, 0, 1)
			wordSearch(i, j, grid, word, 0, -1, 1)
		}
	}

	fmt.Println(totalCount)
}

func wordSearch(row int, col int, grid [][]rune, word []rune, index int, dirX int, dirY int) bool {
	// Ignore out of bounds
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid) {
		return false
	}

	// We've found the target word
	if index == len(word)-1 && grid[row][col] == word[index] {
		totalCount++
		return true
	}

	// Cell doesn't match target word
	if grid[row][col] != word[index] {
		return false
	}

	// Check adjacent cells
	newRow := row + dirX
	newCol := col + dirY

	return wordSearch(newRow, newCol, grid, word, index+1, dirX, dirY)
}
