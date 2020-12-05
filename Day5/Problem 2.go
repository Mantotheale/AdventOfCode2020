package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  // Open the input file and create a scanner to parse it
  file, _ := os.Open("Input.txt")
  scanner := bufio.NewScanner(file)
  
  // Slice to store all the seats IDs
  var seats []int

  for scanner.Scan() {
    // Converts the F, B, L, R into 0, 1, 0, 1 and then in decimal
    var id int = ConvertInDecimal(scanner.Text())

    // Add the new Id to the slice
    seats = append(seats, id)
  }
  
  // Sort the ids in increasing order
  seats = Sort(seats)

  fmt.Println(FindId(seats))
}

// Function to convert the string of F, B, L, R in its actual decimal representation of seat id
func ConvertInDecimal(clear string) (dec int) {
  // Every Bs and Rs are 1s in binary, so when one of this chars is encountered, you add the i-th power of 2, where i is the position of the char (from right)
  for i := 0; i < len(clear); i++ {
    var char rune = rune(clear[len(clear) - 1 - i])
                         
    if char == 'B' || char == 'R' {
      dec += PowTwo(i) 
    }
  }
  return
}

// Function to return the n-th power of 2
func PowTwo(exp int) int {
  var res int = 1

  for i := 0; i < exp; i++ {
    res *= 2
  }

  return res
}

// Sorts the seats in increasing order
func Sort(seats []int) []int {
  for i := 0; i < len(seats); i++ {
    var indMin = i

    for j := i + 1; j < len(seats); j++ {
      if seats[j] < seats[indMin] {
        indMin = j
      }
    }

    seats[i], seats[indMin] = seats[indMin], seats[i]
  }

  return seats
}

// Function to find the id of the free seat
func FindId(seats []int) int {
  for i := 0; i < len(seats) - 1; i++ {
    if seats[i] % 2 == seats[i+1] % 2 {
      return seats[i] + 1
    }
  }
  return -1
}
