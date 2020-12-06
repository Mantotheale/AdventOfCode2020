package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  var groups []string = ReadGroups()
  var total int
  
  for _, g := range groups {
    total += CountLetters(g)
  }

  fmt.Println(total)
}

func ReadGroups() (groups []string) {
  file, _ := os.Open("Input.txt")
  scanner := bufio.NewScanner(file)

  var s string
  for scanner.Scan() {
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

func CountLetters(s string) (count int) {
  var letters [26]bool

  for _, c := range s {
    var charIndex int = int(c - 'a')
    letters[charIndex] = true
  }

  for _, l := range letters {
    if l == true {
      count++
    }
  }

  return
}
