package main

import (
  "fmt"
  "os"
  "bufio"
)

type seat struct {
  row int
  col int
  id int
}

func main() {
  // Slice which contains all the seats which are taken
  var seats []seat = ReadSeats()

  // Find the max id
  var maxId = FindMaxId(seats)

  fmt.Println(maxId)
}

// Function to store the input into a slice for organization
func ReadSeats() (seats []seat) {
  // Open the input file and create a scanner to parse it
  file, _ := os.Open("Input.txt")
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    // Create a seat variabile and convert the code of the row in an actual seat
    var s seat = ConvertRow(scanner.Text())

    // Add the seat to the slice
    seats = append(seats, s)
  }

  return
}

func ConvertRow(row string) (s seat)  {
  // Split the string, one part is the rowFinder, the second part is the colFinder
  var rowFinder, colFinder string = row[:7], row[7:]

  // Find the row index and the col index of the seat
  s.row = LocateIndex(rowFinder, 0, 127) 
  s.col = LocateIndex(colFinder, 0, 7)

  // Find the seat ID by multiplying the row by 8 and adding the col
  s.id = (s.row * 8) + s.col

  return
}

// Function to locate the resulting index (of rows or cols) given the string of instructions
func LocateIndex (instructions string, minRange int, maxRange int) int {
  for _, instr := range instructions {
    // If the instr is "F" or "L", the maxRange becomes its center, the opposite otherwise
    if instr == 'F' || instr == 'L' {
      maxRange = (minRange + maxRange - 1) / 2
    } else {
      minRange = (minRange + maxRange + 1) / 2
    }
  }

  // minRange and maxRange are now equal, and they are the index which was searched
  return minRange
}

// Function to return the index of the max id in the taken seats
func FindMaxId (seats []seat) (max int) {
  // The IDs can't be less than 0, so the max starts by being 0
  // Then scan the entire slice, if it's found a greater element than max, it becomes the new max
  for _, s := range seats {
    if s.id > max {
      max = s.id
    }
  }

  return
}
