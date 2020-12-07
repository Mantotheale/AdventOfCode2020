package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

type Node struct {
  color string
  content []string
}

var colorsVisited []string

func main() {
  var nodes []Node = ReadNodes()
  fmt.Println(CountBags(nodes, "shiny gold"))
}

func ReadNodes() (nodes []Node) {
  file, _ := os.Open("Input.txt")
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    var n Node = ParseNode(scanner.Text())
    nodes = append(nodes, n)
  }

  return
}

func ParseNode(s string) (n Node) {
  // Extract the color
  n.color = s[:FindSecondSpace(s)]
  s = s[FindSecondSpace(s) + 14:]
  
  var containedStrings []string = strings.Split(s, ", ")

  // Return with an empty container slice if no other bags in it
  if s[:2] == "no" {
    return
  }

  for _, c := range containedStrings {
    c = c[2:]
    c = c[:FindSecondSpace(c)]

    n.content = append(n.content, c)
  }

  return
}

func FindSecondSpace (s string) int {
  var spaces int

  for i := 0; i < len(s); i++ {
    if s[i] == ' ' {
      spaces++
      if spaces == 2 {
        return i
      }
    }
  }

  return 0
}

func CountBags(nodes []Node, color string) (count int) {
  for _, n := range nodes {
    if IsContained(n, color) && !IsVisited(n) {
      colorsVisited = append(colorsVisited, n.color)
      count += CountBags(nodes, n.color) + 1
    }
  }

  return
}

func IsContained(n Node, color string) bool {
  for _, c := range n.content {
    if c == color {
      return true
    }
  }

  return false
}

func IsVisited(n Node) bool {
  for _, c := range colorsVisited {
    if c == n.color {
      return true
    }
  }

  return false
}
