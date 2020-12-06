package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  // Slice of strings to store all the answers of every group
  var groups []string = ReadGroups()
  var total int
  
  // For every group, count the total of different letters
  for _, g := range groups {
    total += CountLetters(g)
  }

  fmt.Println(total)
}

// Function to store the datas in input into a slice
func ReadGroups() (groups []string) {
  file, _ := os.Open("Input.txt")
  scanner := bufio.NewScanner(file)

  var s string
  for scanner.Scan() {
    // By doing this, every block of text separated by a blank row is stored in a single string
    if scanner.Text() == "" {
      groups = append(groups, s)
      s = ""
    } else {
      s += scanner.Text()
    }
  }
  groups = append(groups, s)
  
  return
}

// Function to count the number of different letters in a string
func CountLetters(s string) (count int) {
  // We can assume that all letters are from the latin alphabet and lowercase, so create an array which tells if any letter is present
  var letters [26]bool

  // In this way, if a character c is present, the value of the array which represents c is set to true
  for _, c := range s {
    var charIndex int = int(c - 'a')
    letters[charIndex] = true
  }

  // Now count the total of letters present simply by counting the trues in the array
  for _, l := range letters {
    if l == true {
      count++
    }
  }

  return
}
