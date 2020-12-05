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
  
  // The var which contains the max id
  var maxId int

  for scanner.Scan() {
    // Converts the F, B, L, R into 0, 1, 0, 1 and then in decimal
    var id int = ConvertInDecimal(scanner.Text())
    
    // If the new id is greater than the max, the new max is the id
    if id > maxId {
      maxId = id
    }
  }
  
  fmt.Println(maxId)
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
