package main

import (
  "fmt"
  "os"
  "bufio"
)

// Struct which represents all the answers of a group, and its size. It's made only for convenience in writing clear code
type group struct {
  answers string
  size int
}

func main() {
  // Store all the inputs into a slice of groups
  var groups []group = ReadGroups()
  var total int
  
  // For every group, add the number of letters present in every answer
  for _, g := range groups {
    total += CountLetters(g)
  }

  fmt.Println(total)
}

// Function to store all inputs in a slice of strings
func ReadGroups() (groups []group) {
  file, _ := os.Open("Input.txt")
  scanner := bufio.NewScanner(file)

  var g group
  for scanner.Scan() {
    if scanner.Text() == "" {
      // This way, every block of text separated by a blank row is store in a group struct
      groups = append(groups, g)
      g.answers = ""
      g.size = 0
    } else {
      g.answers += scanner.Text()
      g.size++
    }
  }
  groups = append(groups, g)
  
  return
}

// Function to count the letters which are presents in every answer of the group
func CountLetters(g group) (count int) {
  // We can assume that all the letters are latin and lowercase.
  // So we create an array in which is stored the number of occurencies of every letter
  var letters [26]int

  // For every character c, the value at its index is increased by one
  for _, c := range g.answers {
    var charIndex int = int(c - 'a')
    letters[charIndex]++
  }

  // So in the end, the only chars which are present n times in the answers (where n is the size of the group), are the only ones which are answered by everyone
  for _, l := range letters {
    if l == g.size {
      count++
    }
  }

  return
}
