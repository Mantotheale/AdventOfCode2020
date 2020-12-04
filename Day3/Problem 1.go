package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
  // Slice to store the various elements of the field
	var grid []string = ReadGrid()
  // Vars to store the present position and the total of trees
	var row, col, trees int
	
  // Repeat the cycle until you exceed the last row
	for {
    // Update the position int the grid
		row, col = UpdatePosition(row, col)
		
    // If the last row is exceeded, exit from the loop
		if row >= len(grid) {
			break
		}
		
    // If the present position is a tree, increase the counter
		if IsTree(row, col, grid) {
			trees++
		}
	}
	
	fmt.Println(trees)
}

// Function to store the input grid
func ReadGrid() (grid []string) {
  // Open the input file and create a scanner to parse it
  file, _ := os.Open("Input.txt")
  scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
    // Aggiungo la nuova riga alla slice
		grid = append(grid, scanner.Text())
	}

  return
}

// Function to compute the next positon
func UpdatePosition(row int, col int) (int, int) {
  // 1 down, 3 to the right
	return row + 1, col + 3
}

// Function to tell if a position is a tree
func IsTree (row int, col int, grid []string) bool {
  // Save the length of the given row (the same pattern continues after exceeding the row length)
	var rowLength int = len(grid[0])
  
	if grid[row][col % rowLength] == '#' {
		return true
	}
	
	return false
}
