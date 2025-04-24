package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][]int{
	// row, column
	{-1, -1}, // top left
	{-1, 0},  // top
	{-1, 1},  // top right
	{0, -1},  // center left
	{0, 1},   // center right
	{1, -1},  // bottom left
	{1, 0},   // bottom
	{1, 1},   // bottom right
}

func search(word string, grid []string) (count int) {
	for y, row := range grid {
		for x := range row {
		Search:
			for _, direction := range directions {
				posX, posY := x, y
				// Iterate over the word and move grid position in the current direction
				for i := 0; i < len(word); i, posX, posY = i+1, posX+direction[0], posY+direction[1] {
					// Bail on grid boundaries or character mismatch
					if posX < 0 ||
						posY < 0 ||
						posX >= len(row) ||
						posY >= len(grid) ||
						grid[posY][posX] != word[i] {
						continue Search
					}
				}
				count++
			}
		}
	}
	return
}

func main() {
	// Load the file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	var data []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	// Find all instances of "XMAS"
	word := "XMAS"
	count := search(word, data)
	fmt.Printf("%v instances of %v", count, word)
}
