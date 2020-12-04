package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
  // Slice to store the various elements of the field
	var grid []string = ReadGrid()
  // Array to store the number of trees in the five paths
  var trees [5]int


  // For every path, compute the number of trees encounteres
  for i := range trees {
    // Vars to store the present position
    var row, col int

    // Cycle until you exceed the last row
    for {
      // Update the position, based on the path chosen
			row, col = UpdatePosition(row, col, i)
			
      // If the row is exceeded, break the loop
			if row >= len(grid) {
				break
			}
			
      // If the position is a tree, increase the counter
			if IsTree(row, col, grid) {
				trees[i]++
			}
		}
  }
  
  // Print of the product of the various trees on the paths
  fmt.Println(Product(trees))
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
func UpdatePosition(row int, col int, route int) (int, int) {
	switch route {
    // The first path is 1 right and 1 down
		case 0: return row + 1, col + 1
    // The second path is 3 right and 1 down
		case 1: return row + 1, col + 3
    // The third path is 5 right and 1 down
    case 2: return row + 1, col + 5
		// The fourth path is 7 right and 1 down
    case 3: return row + 1, col + 7
		// The fifth path is 1 right and 2 down
    default: return row + 2, col + 1
	}
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

// Function to compute the product of the trees on the different paths
func Product(trees [5]int) int64 {
  var prod int64 = 1

  for _, t := range trees {
    prod *= int64(t)
  }

  return prod
}
