package main

import (
  "fmt"
  "os"
  "bufio"
)

type group struct {
  answers string
  size int
}

func main() {
  var groups []group = ReadGroups()
  var total int
  
  for _, g := range groups {
    total += CountLetters(g)
  }

  fmt.Println(total)
}

func ReadGroups() (groups []group) {
  file, _ := os.Open("Input.txt")
  scanner := bufio.NewScanner(file)

  var g group
  for scanner.Scan() {
    if scanner.Text() == "" {
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

func CountLetters(g group) (count int) {
  var letters [26]int

  for _, c := range g.answers {
    var charIndex int = int(c - 'a')
    letters[charIndex]++
  }

  for _, l := range letters {
    if l == g.size {
      count++
    }
  }

  return
}
