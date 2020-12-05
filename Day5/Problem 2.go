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
  // Array in which the i-th element is true if the i-th row is full, false otherwise
  var rows [128]bool = FindFullRows(seats)
  
  // Find the only row not full in which the row greater than and less than that are full
  var finalRow int = FindFinalRow(rows)
  // Find the only col of the row in which the seat is free
  var finalCol int = FindFinalCol(seats, finalRow)

  fmt.Println((finalRow * 8) + finalCol)
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

// Function to find which rows are all full
func FindFullRows(seats []seat) (rows [128]bool) {
  // Scan every row, for every row, if all the seats are found, the row is full
  for i := 0; i < 128; i++ {
    var isFull = true

    for j := 0; j < 8; j++ {
      if !IsFound(seats, i, j) {
        isFull = false
        break
      }
    }

    rows[i] = isFull 
  }

  return
}

// Function to tell wheter or not a certain seat is taken
func IsFound(seats []seat, row int, col int) bool {
  for _, s := range seats {
    if s.row == row && s.col == col {
      return true
    }
  }

  return false
}

// Function to find the row of the main character's seat
func FindFinalRow(rows [128]bool) int {
  // The row is the right one if it's not full and the previous and the next rows are full
  for i := 1; i < len(rows) - 1; i++ {
    if rows[i] == false {
      if rows[i - 1] == true && rows[i + 1] == true {
        return i
      }
    }
  }

  return -1
}

// Function to find the only seat free of the given row
func FindFinalCol(seats []seat, row int) int {
  for i := 0; i < 8; i++ {
    var isFree bool = true 

    for _, s := range seats {
      if s.row == row && s.col == i {
        isFree = false
      }
    }

    if isFree {
      return i
    }
  }

  return -1
}
